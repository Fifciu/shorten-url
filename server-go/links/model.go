package links

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/omeid/pgerror"
)

const (
	JWT_TTL = 60 * 60 * 24 // 1 day
	JWT_KEY = "some-random-zaq1@WSX"
)

const PER_PAGE = 12

type Link struct {
	ID          uint      `json:"id"`
	UserId      uint      `json:"user_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	OriginalUrl string    `json:"original_url" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at" validate:"required"`
	Alias       string    `json:"alias" validate:""` // TODO: add better validation
}

func (l *Link) MarshalJSON() ([]byte, error) {
	type Alias Link
	time := fmt.Sprintf("%02d-%02d-%d %02d:%02d", l.UpdatedAt.Day(), l.UpdatedAt.Month(), l.UpdatedAt.Year(), l.UpdatedAt.Hour(), l.UpdatedAt.Minute())
	return json.Marshal(&struct {
		*Alias
		UpdatedAt string `json:"updated_at"`
	}{
		Alias:     (*Alias)(l),
		UpdatedAt: time,
	})
}

type PaginatedLinks struct {
	Count      uint    `json:"count"`
	PerPage    uint    `json:"per_page"`
	Page       uint    `json:"page"`
	PageAmount uint    `json:"page_amount"`
	Links      []*Link `json:"links"`
}

type LinksModel interface {
	CreateLink(c CreateLinkDto) (*Link, error)
	BuildAlias(finalLen int) (string, error)
	AliasExists(alias string) (bool, error)
	FindLinkByAlias(alias string) (*Link, error)
	IsLinkOwner(userId uint, linkId uint) (bool, error)
	DeleteLink(linkId uint) error
	GetLinksOfUser(userId uint, sort string, direction string, page uint) (*PaginatedLinks, error)
	UserUsedLinkName(linkName string, userId uint) (bool, error)
	UpdateLink(linkId, userId uint, body UpdateLinkDto) (*Link, error)
}

type PostgresLinkModel struct {
	Db *sql.DB // todo: put here db handler
}

// TODO: Write tests

func (p *PostgresLinkModel) CreateLink(c CreateLinkDto) (*Link, error) {
	var lastInsertedId uint
	err := p.Db.QueryRow("INSERT INTO links (user_id, name, original_url, updated_at, alias) VALUES ($1, $2, $3, $4, $5) RETURNING id", c.UserId, c.Name, c.OriginalUrl, c.UpdatedAt, c.Alias).Scan(&lastInsertedId)
	if err != nil {
		if e := pgerror.UniqueViolation(err); e != nil {
			return nil, errors.New(http.StatusText(http.StatusConflict)) // email exists
		}
		return nil, err
	}

	return &Link{
		ID:          lastInsertedId,
		UserId:      c.UserId,
		Name:        c.Name,
		OriginalUrl: c.OriginalUrl,
		UpdatedAt:   c.UpdatedAt,
		Alias:       c.Alias,
	}, nil
}

func (p *PostgresLinkModel) BuildAlias(finalLen int) (string, error) {
	// 65A - 90Z, 97a-122z,
	characters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	charLen := len(characters)
	final := make([]byte, finalLen)
	for i := 0; i < finalLen; i++ {
		rand := int(math.Round(rand.Float64() * float64(charLen-1)))
		final[i] = characters[rand]
	}
	result := string(final)
	exists, err := p.AliasExists(result)
	if err != nil {
		return "", err
	}
	if exists {
		return p.BuildAlias(finalLen)
	}
	return result, err
}

func (p *PostgresLinkModel) AliasExists(alias string) (bool, error) {
	var exists bool
	err := p.Db.QueryRow("SELECT EXISTS(SELECT 1 FROM links WHERE alias = $1) as exists", alias).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (p *PostgresLinkModel) UserUsedLinkName(linkName string, userId uint) (bool, error) {
	var exists bool
	err := p.Db.QueryRow("SELECT EXISTS(SELECT 1 FROM links WHERE name = $1 AND user_id = $2) as exists", linkName, userId).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (p *PostgresLinkModel) FindLinkByAlias(alias string) (*Link, error) {
	var link Link
	err := p.Db.QueryRow("SELECT id, user_id, name, original_url, updated_at, alias FROM links WHERE alias = $1", alias).Scan(&link.ID, &link.UserId, &link.Name, &link.OriginalUrl, &link.UpdatedAt, &link.Alias)
	if err != nil {
		if e := pgerror.NoDataFound(err); e != nil {
			return nil, errors.New(http.StatusText(http.StatusNotFound))
		}
		return nil, err
	}
	return &link, nil
}

func (p *PostgresLinkModel) IsLinkOwner(userId uint, linkId uint) (bool, error) {
	var exists bool // TODO: Can I safely remove exists? Or .Scan assign err?
	err := p.Db.QueryRow("SELECT EXISTS(SELECT 1 FROM links WHERE user_id = $1 AND id = $2)", userId, linkId).Scan(&exists)
	if err != nil {
		if e := pgerror.NoDataFound(err); e != nil {
			return false, nil
		}
		return false, err
	}
	// TODO: Simplify
	return exists, nil
}

func (p *PostgresLinkModel) DeleteLink(linkId uint) error {
	res, err := p.Db.Exec("DELETE FROM links WHERE id = $1", linkId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New(http.StatusText(http.StatusNotFound))
	}
	return nil
}

func (p *PostgresLinkModel) GetLinksOfUser(userId uint, sort string, sortDirection string, page uint) (*PaginatedLinks, error) {
	if sort != "updated_at" && sort != "name" {
		return nil, errors.New(http.StatusText(http.StatusBadRequest))
	}

	var count int
	err := p.Db.QueryRow("SELECT COUNT(*) FROM links WHERE user_id = $1", userId).Scan(&count)
	if err != nil {
		return nil, err
	}

	links := make([]*Link, 0) // TODO: Can I extract rows amount from query?

	pageAmount := uint(math.Ceil(float64(count) / float64(PER_PAGE)))
	if pageAmount == 0 {
		return &PaginatedLinks{
			Count:      0,
			PerPage:    PER_PAGE,
			Page:       uint(1),
			PageAmount: uint(0),
			Links:      links,
		}, nil
	}
	if page > pageAmount {
		return nil, errors.New(http.StatusText(http.StatusBadRequest))
	}
	offset := (page - 1) * PER_PAGE
	rows, err := p.Db.Query(fmt.Sprintf("SELECT id, name, original_url, updated_at, alias FROM links WHERE user_id = $1 ORDER BY %s %s LIMIT %d OFFSET %d", sort, sortDirection, PER_PAGE, offset), userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		link := &Link{}
		if err := rows.Scan(&link.ID, &link.Name, &link.OriginalUrl, &link.UpdatedAt, &link.Alias); err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &PaginatedLinks{
		Count:      uint(count),
		Page:       page,
		PerPage:    PER_PAGE,
		PageAmount: pageAmount,
		Links:      links,
	}, nil
}

func (p *PostgresLinkModel) UpdateLink(linkId, userId uint, body UpdateLinkDto) (*Link, error) {
	// Updating
	argCounter := 1
	q := `UPDATE links SET `
	qParts := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)

	// TODO: what if there are no fields?

	if body.Name != nil {
		qParts = append(qParts, fmt.Sprintf(`name = $%d`, argCounter))
		args = append(args, body.Name)
		argCounter++
	}

	if body.OriginalUrl != nil {
		qParts = append(qParts, fmt.Sprintf(`original_url = $%d`, argCounter))
		args = append(args, body.OriginalUrl)
		argCounter++
	}

	if body.Alias != nil {
		qParts = append(qParts, fmt.Sprintf(`alias = $%d`, argCounter))
		args = append(args, body.Alias)
		argCounter++
	}

	qParts = append(qParts, fmt.Sprintf(`updated_at = $%d`, argCounter))
	args = append(args, time.Now())
	argCounter++

	q += strings.Join(qParts, ",") + fmt.Sprintf(` WHERE id = $%d `, argCounter)
	argCounter++
	q += fmt.Sprintf(`AND user_id = $%d`, argCounter)
	args = append(args, linkId)
	args = append(args, userId)

	res, err := p.Db.Exec(q, args...)
	if err != nil {
		return nil, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if count < 1 {
		return nil, errors.New("You are not an owner of specified link")
	}

	// Fetching link
	var link Link
	err = p.Db.QueryRow("SELECT id, user_id, name, original_url, updated_at, alias FROM links WHERE id = $1", linkId).Scan(&link.ID, &link.UserId, &link.Name, &link.OriginalUrl, &link.UpdatedAt, &link.Alias)
	if err != nil {
		if e := pgerror.NoDataFound(err); e != nil {
			return nil, errors.New(http.StatusText(http.StatusNotFound))
		}
		return nil, err
	}
	return &link, nil
}

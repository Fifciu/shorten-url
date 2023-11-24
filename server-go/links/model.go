package links

import (
	"database/sql"
	"errors"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/omeid/pgerror"
)

const (
	JWT_TTL = 60 * 60 * 24 // 1 day
	JWT_KEY = "some-random-zaq1@WSX"
)

type Link struct {
	ID          uint      `json:"id"`
	UserId      uint      `json:"user_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	OriginalUrl string    `json:"original_url" validate:"required"`
	CreatedAt   time.Time `json:"created_at" validate:"required"`
	Alias       string    `json:"alias" validate:""` // TODO: add better validation
}

type LinksModel interface {
	CreateLink(c CreateLinkDto) (*Link, error)
	BuildAlias(finalLen int) (string, error)
	AliasExists(alias string) (bool, error)
	FindLinkByAlias(alias string) (*Link, error)
	IsLinkOwner(userId uint, linkId uint) (bool, error)
	DeleteLink(linkId uint) error
	GetLinksOfUser(userId uint) ([]*Link, error)
}

type PostgresLinkModel struct {
	Db *sql.DB // todo: put here db handler
}

// TODO: Write tests

func (p *PostgresLinkModel) CreateLink(c CreateLinkDto) (*Link, error) {
	var lastInsertedId uint
	err := p.Db.QueryRow("INSERT INTO links (user_id, name, original_url, created_at, alias) VALUES ($1, $2, $3, $4, $5) RETURNING id", c.UserId, c.Name, c.OriginalUrl, c.CreatedAt, c.Alias).Scan(&lastInsertedId)
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
		CreatedAt:   c.CreatedAt,
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

func (p *PostgresLinkModel) FindLinkByAlias(alias string) (*Link, error) {
	var link Link
	err := p.Db.QueryRow("SELECT id, user_id, name, original_url, created_at, alias FROM links WHERE alias = $1", alias).Scan(&link.ID, &link.UserId, &link.Name, &link.OriginalUrl, &link.CreatedAt, &link.Alias)
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

func (p *PostgresLinkModel) GetLinksOfUser(userId uint) ([]*Link, error) {
	rows, err := p.Db.Query("SELECT id, name, original_url, created_at, alias FROM links WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	links := make([]*Link, 0) // TODO: Can I extract rows amount from query?
	for rows.Next() {
		link := &Link{}
		if err := rows.Scan(&link.ID, &link.Name, &link.OriginalUrl, &link.CreatedAt, &link.Alias); err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return links, nil
}

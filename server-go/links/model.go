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
	CreatedAt   time.Time `json:"name" validate:"required"`
	Alias       string    `json:"alias" validate:""` // TODO: add better validation
}

type LinksModel interface {
	CreateLink(c CreateLinkDto) (*Link, error)
	BuildAlias(finalLen int) (string, error)
	AliasExists(alias string) (bool, error)
}

type PostgresLinkModel struct {
	db *sql.DB // todo: put here db handler
}

// TODO: Write tests

func (p *PostgresLinkModel) CreateLink(c CreateLinkDto) (*Link, error) {
	var lastInsertedId uint
	err := p.db.QueryRow("INSERT INTO links (user_id, name, original_url, created_at, alias) VALUES ($1, $2, $3, $4, $5) RETURNING id", c.UserId, c.Name, c.OriginalUrl, c.CreatedAt, c.Alias).Scan(&lastInsertedId)
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
	err := p.db.QueryRow("SELECT EXISTS(SELECT 1 FROM links WHERE alias = $1) as exists", alias).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

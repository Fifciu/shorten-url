package redirects

import (
	"database/sql"

	"github.com/Fifciu/shorten-url/server-go/links"
	"github.com/go-chi/chi/v5"
)

type RedirectsModule struct {
	db *sql.DB
}

func (r *RedirectsModule) GetNamespace() string {
	return "r"
}

func (r *RedirectsModule) GetRouter() *chi.Mux {
	redirectsController := NewRedirectsController(&links.PostgresLinkModel{
		Db: r.db,
	})
	redirectsRouter := chi.NewRouter()
	for path, handler := range redirectsController.Endpoints {
		redirectsRouter.Get(path, handler)
	}
	return redirectsRouter
}

func NewRedirectsModule(db *sql.DB) *RedirectsModule {
	return &RedirectsModule{db}
}

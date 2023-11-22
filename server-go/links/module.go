package links

import (
	"database/sql"
	"log"

	"github.com/go-chi/chi/v5"
)

type LinksModule struct {
	db *sql.DB
}

func (a *LinksModule) GetNamespace() string {
	return "links"
}

func (a *LinksModule) GetRouter() *chi.Mux {
	linksController := NewLinksController(&PostgresLinkModel{
		Db: a.db,
	})
	linkRouter := chi.NewRouter()
	for _, endpoint := range linksController {
		switch endpoint.Method {
		case "POST":
			linkRouter.With(endpoint.Middlewares...).Post(endpoint.Path, endpoint.Handler)
		case "GET":
			linkRouter.With(endpoint.Middlewares...).Get(endpoint.Path, endpoint.Handler)
		case "DELETE":
			linkRouter.With(endpoint.Middlewares...).Delete(endpoint.Path, endpoint.Handler)
		case "PUT":
			linkRouter.With(endpoint.Middlewares...).Put(endpoint.Path, endpoint.Handler)
		case "PATCH":
			linkRouter.With(endpoint.Middlewares...).Patch(endpoint.Path, endpoint.Handler)
		default:
			log.Fatalf("%q HTTP method used in Links controller. It's not supported!", endpoint.Method)
		}
	}
	return linkRouter
}

func NewLinksModule(db *sql.DB) *LinksModule {
	return &LinksModule{db}
}

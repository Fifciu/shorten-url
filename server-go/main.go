package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	authentication "github.com/Fifciu/shorten-url/server-go/authentication"
	links "github.com/Fifciu/shorten-url/server-go/links"
	redirects "github.com/Fifciu/shorten-url/server-go/redirects"
	users "github.com/Fifciu/shorten-url/server-go/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

const PORT = 2222

type Module interface {
	GetNamespace() string
	GetRouter() *chi.Mux
}

var db *sql.DB

func init() {
	var err error
	connStr := "postgres://user:zaq1@WSX@localhost:5432/shortenurl?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to the DB")
}

func main() {
	mainRouter := chi.NewRouter()
	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))
	modules := []Module{
		authentication.NewAuthModule(db),
		users.NewUsersModule(db),
		links.NewLinksModule(db),
		redirects.NewRedirectsModule(db),
	}

	for _, module := range modules {
		namespace := fmt.Sprintf("/%s", module.GetNamespace())
		log.Printf("Registered module under %q namespace", namespace)
		mainRouter.Mount(namespace, module.GetRouter())
	}

	port := fmt.Sprintf(":%d", PORT)
	err := http.ListenAndServe(port, mainRouter)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	links "github.com/Fifciu/shorten-url/server-go/links"
	users "github.com/Fifciu/shorten-url/server-go/users"
	"github.com/go-chi/chi/v5"
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
	modules := []Module{
		users.NewAuthModule(db),
		links.NewLinksModule(db),
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

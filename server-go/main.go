package main

import (
	"fmt"
	"log"
	"net/http"

	AuthModule "github.com/Fifciu/shorten-url/server-go/authentication"
	"github.com/go-chi/chi/v5"
)

const PORT = 2222

type Module interface {
	GetNamespace() string
	GetRouter() *chi.Mux
}

func main() {
	mainRouter := chi.NewRouter()
	modules := []Module{
		AuthModule.NewAuthModule(),
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

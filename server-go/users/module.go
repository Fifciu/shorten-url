package users

import (
	"database/sql"
	"log"

	"github.com/go-chi/chi/v5"
)

type UsersModule struct {
	db *sql.DB
}

func (a *UsersModule) GetNamespace() string {
	return "users"
}

func (a *UsersModule) GetRouter() *chi.Mux {
	usersController := NewUsersController(&PostgresUserModel{
		Db: a.db,
	})
	userRouter := chi.NewRouter()
	for _, endpoint := range usersController {
		switch endpoint.Method {
		case "POST":
			userRouter.With(endpoint.Middlewares...).Post(endpoint.Path, endpoint.Handler)
		case "GET":
			userRouter.With(endpoint.Middlewares...).Get(endpoint.Path, endpoint.Handler)
		case "DELETE":
			userRouter.With(endpoint.Middlewares...).Delete(endpoint.Path, endpoint.Handler)
		case "PUT":
			userRouter.With(endpoint.Middlewares...).Put(endpoint.Path, endpoint.Handler)
		case "PATCH":
			userRouter.With(endpoint.Middlewares...).Patch(endpoint.Path, endpoint.Handler)
		default:
			log.Fatalf("%q HTTP method used in Links controller. It's not supported!", endpoint.Method)
		}
	}
	return userRouter
}

func NewUsersModule(db *sql.DB) *UsersModule {
	return &UsersModule{db}
}

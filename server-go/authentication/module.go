package authentication

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

type AuthModule struct {
	db *sql.DB
}

func (a *AuthModule) GetNamespace() string {
	return "authentication"
}

func (a *AuthModule) GetRouter() *chi.Mux {
	authController := NewAuthController(&PostgresUserModel{
		db: a.db,
	})
	authRouter := chi.NewRouter()
	for path, handler := range authController.Endpoints {
		authRouter.HandleFunc(path, handler)
	}
	return authRouter
}

func NewAuthModule(db *sql.DB) *AuthModule {
	return &AuthModule{db}
}

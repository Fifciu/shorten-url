package users

import (
	"github.com/go-chi/chi/v5"
)

type AuthModule struct{}

func (a *AuthModule) GetNamespace() string {
	return "authentication"
}

func (a *AuthModule) GetRouter() *chi.Mux {
	authController := NewAuthController(&MysqlUserModel{})
	authRouter := chi.NewRouter()
	for path, handler := range authController.Endpoints {
		authRouter.HandleFunc(path, handler)
	}
	return authRouter
}

func NewAuthModule() *AuthModule {
	return &AuthModule{}
}

package links

import (
	"net/http"

	"github.com/Fifciu/shorten-url/server-go/middlewares"
	"github.com/Fifciu/shorten-url/server-go/utils"
	"github.com/go-playground/validator/v10"
)

type LinksHandler struct {
	Path        string
	Handler     http.HandlerFunc
	Middlewares []func(http.Handler) http.Handler
	Method      string
}

func AddNewLink(validate *validator.Validate, model LinksModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.JsonErrorResponse(w, 200, "xd")
	})
}

func NewLinksController(model LinksModel) []*LinksHandler {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return []*LinksHandler{
		{
			Path:        "/",
			Method:      "POST",
			Middlewares: []func(http.Handler) http.Handler{middlewares.Authenticated},
			Handler:     AddNewLink(validate, model),
		},
	}
}

package redirects

import (
	"fmt"
	"net/http"

	"github.com/Fifciu/shorten-url/server-go/links"
	utils "github.com/Fifciu/shorten-url/server-go/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type RedirectsController struct {
	Endpoints map[string]http.HandlerFunc
}

const (
	frontendNotFoundUrl = "/page-not-found"
)

func Redirect(validate *validator.Validate, model links.LinksModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alias := chi.URLParam(r, "alias")

		link, err := model.FindLinkByAlias(alias)
		if err != nil {
			if err.Error() == http.StatusText(http.StatusInternalServerError) {
				log.Error(fmt.Sprintf("controller/redirects/Redirect/fetching link: %s", err.Error()))
				utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
			} else {
				log.Error(fmt.Sprintf("controller/redirects/Redirect/fetching link: %s %q", "Couldn't find alias", alias))
				http.Redirect(w, r, frontendNotFoundUrl, http.StatusMovedPermanently)
			}
			return
		}
		http.Redirect(w, r, link.OriginalUrl, http.StatusMovedPermanently)
	}
}

func NewRedirectsController(model links.LinksModel) *RedirectsController {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return &RedirectsController{
		Endpoints: map[string]http.HandlerFunc{
			"/{alias}": Redirect(validate, model),
		},
	}
}

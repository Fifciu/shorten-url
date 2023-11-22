package links

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Fifciu/shorten-url/server-go/middlewares"
	"github.com/Fifciu/shorten-url/server-go/users"
	"github.com/Fifciu/shorten-url/server-go/utils"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type LinksHandler struct {
	Path        string
	Handler     http.HandlerFunc
	Middlewares []func(http.Handler) http.Handler
	Method      string
}

func AddNewLink(validate *validator.Validate, model LinksModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var body CreateLinkDto
		err := d.Decode(&body)
		if err != nil {
			log.Error(fmt.Sprintf("controller/links/AddNewLink/creating link: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Couldn't parse the payload")
			return
		}

		err = validate.Struct(body)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Some fields are missing or have incorrect value")
			return
		}

		values := r.Context().Value("claims").(*users.Claims)
		body.UserId = values.UserClaims.ID
		body.CreatedAt = time.Now()
		// TODO: Add min length
		if body.Alias == "" {
			body.Alias, err = model.BuildAlias(10)
			if err != nil {
				log.Error(fmt.Sprintf("controller/links/AddNewLink/failed checking building alias: %s", err.Error()))
				utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}
		} else {
			exists, err := model.AliasExists(body.Alias)
			if err != nil {
				log.Error(fmt.Sprintf("controller/links/AddNewLink/failed checking if alias exists: %s", err.Error()))
				utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}
			if exists {
				log.Error("controller/links/AddNewLink/alias exists!")
				utils.JsonErrorResponse(w, http.StatusConflict, http.StatusText(http.StatusConflict))
				return
			}
		}
		link, err := model.CreateLink(body)
		if err != nil {
			log.Error(fmt.Sprintf("controller/links/AddNewLink/creating link: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		utils.JsonResponse(w, 201, link)
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

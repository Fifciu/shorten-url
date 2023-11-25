package links

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Fifciu/shorten-url/server-go/authentication"
	"github.com/Fifciu/shorten-url/server-go/middlewares"
	"github.com/Fifciu/shorten-url/server-go/utils"
	"github.com/go-chi/chi/v5"
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

		values := r.Context().Value("claims").(*authentication.Claims)
		body.UserId = values.UserClaims.ID
		userUsedLinkName, err := model.UserUsedLinkName(body.Name, body.UserId)
		if err != nil {
			log.Error(fmt.Sprintf("controller/links/AddNewLink/failed checking name for user: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		if userUsedLinkName {
			log.Error(fmt.Sprintf("controller/links/AddNewLink/user already used this name: %s", body.Name))
			utils.JsonErrorResponse(w, http.StatusConflict, http.StatusText(http.StatusConflict)) // TODO: Better error message
			return
		}
		body.UpdatedAt = time.Now()
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

func DeleteMyLink(validate *validator.Validate, model LinksModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		linkIdString := chi.URLParam(r, "linkId")
		intLinkId, err := strconv.Atoi(linkIdString)
		if err != nil {
			log.Error(fmt.Sprintf("controller/links/DeleteMyLink/reading link id: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		linkId := uint(intLinkId)
		values := r.Context().Value("claims").(*authentication.Claims)
		userId := values.UserClaims.ID

		isOwner, err := model.IsLinkOwner(userId, linkId)
		if err != nil {
			log.Error(fmt.Sprintf("controller/links/DeleteMyLink/checking is link owner: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		if !isOwner {
			log.Error(fmt.Sprintf("controller/links/DeleteMyLink/checking ownerity. UserId %d is not an owner of Link with Id %d", userId, linkId))
			utils.JsonErrorResponse(w, http.StatusForbidden, http.StatusText(http.StatusForbidden))
			return
		}

		err = model.DeleteLink(linkId)

		if err != nil {
			if err.Error() == http.StatusText(http.StatusNotFound) {
				utils.JsonErrorResponse(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			} else {
				log.Error(fmt.Sprintf("controller/links/DeleteMyLink/deleting link: %s", err.Error()))
				utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func GetMyLinks(validate *validator.Validate, model LinksModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := r.Context().Value("claims").(*authentication.Claims)
		userId := values.UserClaims.ID

		links, err := model.GetLinksOfUser(userId)
		if err != nil {
			log.Error(fmt.Sprintf("controller/links/GetMyLinks/fetching user's links: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		utils.JsonResponse(w, 200, links)
	})
}

func PatchMyLink(validate *validator.Validate, model LinksModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		linkIdString := chi.URLParam(r, "linkId")
		intLinkId, err := strconv.Atoi(linkIdString)
		if err != nil {
			log.Error(fmt.Sprintf("controller/links/PatchMyLink/reading link id: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		linkId := uint(intLinkId)
		values := r.Context().Value("claims").(*authentication.Claims)
		userId := values.UserClaims.ID

		isOwner, err := model.IsLinkOwner(userId, linkId)
		if err != nil {
			log.Error(fmt.Sprintf("controller/links/PatchMyLink/checking is link owner: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		if !isOwner {
			log.Error(fmt.Sprintf("controller/links/PatchMyLink/checking ownerity. UserId %d is not an owner of Link with Id %d", userId, linkId))
			utils.JsonErrorResponse(w, http.StatusForbidden, http.StatusText(http.StatusForbidden))
			return
		}

		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var body UpdateLinkDto
		err = d.Decode(&body)
		if err != nil {
			log.Error(fmt.Sprintf("controller/links/PatchMyLink/reading payload: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Couldn't parse the payload")
			return
		}

		err = validate.Struct(body)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Some fields are missing or have incorrect value")
			return
		}

		link, err := model.UpdateLink(linkId, userId, body)
		if err != nil {
			log.Error(fmt.Sprintf("controller/links/PatchMyLink/updating links: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		utils.JsonResponse(w, 200, link)
	})
}

// Actualy it's about MY Links
func NewLinksController(model LinksModel) []*LinksHandler {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return []*LinksHandler{
		{
			Path:        "/",
			Method:      "GET",
			Middlewares: []func(http.Handler) http.Handler{middlewares.Authenticated},
			Handler:     GetMyLinks(validate, model),
		},
		{
			Path:        "/",
			Method:      "POST",
			Middlewares: []func(http.Handler) http.Handler{middlewares.Authenticated},
			Handler:     AddNewLink(validate, model),
		},
		{
			Path:        "/{linkId}", // TODO: Validate it's number
			Method:      "DELETE",
			Middlewares: []func(http.Handler) http.Handler{middlewares.Authenticated},
			Handler:     DeleteMyLink(validate, model),
		},
		{
			Path:        "/{linkId}", // TODO: Validate it's number
			Method:      "PATCH",
			Middlewares: []func(http.Handler) http.Handler{middlewares.Authenticated},
			Handler:     PatchMyLink(validate, model),
		},
	}
}

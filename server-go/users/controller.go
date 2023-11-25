package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Fifciu/shorten-url/server-go/authentication"
	"github.com/Fifciu/shorten-url/server-go/middlewares"
	"github.com/Fifciu/shorten-url/server-go/utils"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UsersHandler struct {
	Path        string
	Handler     http.HandlerFunc
	Middlewares []func(http.Handler) http.Handler
	Method      string
}

func PatchPassword(validate *validator.Validate, model UsersModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := r.Context().Value("claims").(*authentication.Claims)
		userId := values.UserClaims.ID

		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var body UpdateUserPasswordDto
		err := d.Decode(&body)
		if err != nil {
			log.Error(fmt.Sprintf("controller/users/PatchPassword/reading payload: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Couldn't parse the payload")
			return
		}

		err = validate.Struct(body)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Some fields are missing or have incorrect value")
			return
		}

		currentPasswordHash, err := model.GetUserPasswordHash(userId)
		if err != nil {
			log.Error(fmt.Sprintf("controller/users/PatchPassword/fetching current password hash: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(currentPasswordHash), []byte(body.CurrentPassword))
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		authModel := &authentication.PostgresUserModel{} // inject it
		newHashedPassword, err := authModel.HashPassword(body.NewPassword)
		if err != nil {
			log.Error(fmt.Sprintf("controller/users/PatchPassword/hashing password: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		err = model.UpdatePassword(userId, newHashedPassword)
		if err != nil {
			log.Error(fmt.Sprintf("controller/users/PatchPassword/updating password: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		w.WriteHeader(200)
	})
}

func NewUsersController(model UsersModel) []*UsersHandler {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return []*UsersHandler{
		{
			Path:        "/password",
			Method:      "PATCH",
			Middlewares: []func(http.Handler) http.Handler{middlewares.Authenticated},
			Handler:     PatchPassword(validate, model),
		},
	}
}

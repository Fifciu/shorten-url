package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Fifciu/shorten-url/server-go/utils"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	Endpoints map[string]http.HandlerFunc
}

func Register(validate *validator.Validate, model UserModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var body UserCreateDto
		err := d.Decode(&body)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Couldn't parse the payload")
			return
		}

		err = validate.Struct(body)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Some fields are missing or have incorrect value")
			return
		}

		// todo: check d.More()
		hashedPassword, err := model.HashPassword(body.Password)
		if err != nil {
			log.Error(fmt.Sprintf("controller/authentication/Register/hashing password: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		user, err := model.CreateUser(body.Fullname, body.Email, hashedPassword)
		if err != nil {
			log.Error(fmt.Sprintf("controller/authentication/Register/creating user: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())

			return
		}

		token, expirationTime, err := model.GenerateJwtToken(user)
		if err != nil {
			log.Error(fmt.Sprintf("controller/authentication/Register/generating token: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		utils.CookieResponse(w, token, expirationTime)
	}
}

func Login(validate *validator.Validate, model UserModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var body UserLoginDto
		err := d.Decode(&body)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Couldn't parse the payload")
			return
		}

		err = validate.Struct(body)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Some fields are missing or have incorrect value")
			return
		}

		user, err := model.GetUserByEmail(body.Email)
		if err != nil {
			log.Error(fmt.Sprintf("controller/authentication/Login/generating token: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		}

		token, expirationTime, err := model.GenerateJwtToken(user)
		if err != nil {
			log.Error(fmt.Sprintf("controller/authentication/Login/generating token: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		utils.CookieResponse(w, token, expirationTime)
	})
}

// func RefreshToken()                                   {}
// func Logout()                                         {}
// func Me()                                             {}

func NewAuthController(model UserModel) *AuthController {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return &AuthController{
		Endpoints: map[string]http.HandlerFunc{
			"/register": Register(validate, model),
			"/login":    Login(validate, model),
		},
	}
}

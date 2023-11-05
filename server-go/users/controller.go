package users

import (
	"encoding/json"
	"net/http"

	"github.com/Fifciu/shorten-url/server-go/utils"
	"github.com/go-playground/validator/v10"
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
			utils.JsonErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = validate.Struct(body)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Some fields are missing or have incorrect value")
		}

		// todo: check d.More()
		hashedPassword, err := model.HashPassword(body.Password)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		user, err := model.CreateUser(body.Name, body.Email, hashedPassword)
		user.Name = "asdsad"
	}
}

// func Login()                                          {}
// func RefreshToken()                                   {}
// func Logout()                                         {}
// func Me()                                             {}

func NewAuthController(model UserModel) *AuthController {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return &AuthController{
		Endpoints: map[string]http.HandlerFunc{
			"/register": Register(validate, model),
		},
	}
}

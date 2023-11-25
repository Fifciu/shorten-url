package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	utils "github.com/Fifciu/shorten-url/server-go/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

		body.Password = hashedPassword
		body.UpdatedAt = time.Now()
		user, err := model.CreateUser(body)
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
			return
		}

		token, expirationTime, err := model.GenerateJwtToken(user)
		if err != nil {
			log.Error(fmt.Sprintf("controller/authentication/Login/generating token: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		utils.CookieResponse(w, token, expirationTime)
	})
}

func RefreshToken(validate *validator.Validate, model UserModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieKey := "session-token" // put it in one place
		c, err := r.Cookie(cookieKey)
		if err != nil {
			if err == http.ErrNoCookie {
				utils.JsonErrorResponse(w, http.StatusUnauthorized, "Authorization cookie not sent.")
			} else {
				utils.JsonErrorResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			}
			return
		}
		tokenFromCookie := c.Value
		claims := &Claims{}
		jwtKey := JWT_KEY
		tkn, err := jwt.ParseWithClaims(tokenFromCookie, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				log.Error("controller/authentication/RefreshToken/refreshing token: Invalid signature. Unauthorized")
				utils.JsonErrorResponse(w, http.StatusUnauthorized, "Invalid signature. Unauthorized")
			} else {
				log.Error(fmt.Sprintf("controller/authentication/RefreshToken/refreshing token: %s", err.Error()))
				utils.JsonErrorResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			}
			return
		}

		if !tkn.Valid {
			utils.JsonErrorResponse(w, http.StatusUnauthorized, "Invalid token. Unauthorized")
		}

		if (*claims.ExpiresAt).Time.Sub(time.Now()) > 30*time.Second {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "It is too early to refresh the token")
		}
		expirationTime := time.Now().Add(time.Duration(JWT_TTL) * time.Second)
		claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(jwtKey))
		if err != nil {
			log.Error(fmt.Sprintf("controller/authentication/RefreshToken/refreshing token: %s", err.Error()))
			utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		utils.CookieResponse(w, tokenString, expirationTime)
	})
}

func Logout(validate *validator.Validate, model UserModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expirationTime := time.Date(1970, 1, 1, 0, 0, 0, 0, &time.Location{})
		utils.CookieResponse(w, "", expirationTime)
	})
}

// func Me()                                             {}

func NewAuthController(model UserModel) *AuthController {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return &AuthController{
		Endpoints: map[string]http.HandlerFunc{
			"/register": Register(validate, model),
			"/login":    Login(validate, model),
			"/refresh":  RefreshToken(validate, model),
			"/logout":   Logout(validate, model),
		},
	}
}

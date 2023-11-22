package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Fifciu/shorten-url/server-go/users"
	"github.com/Fifciu/shorten-url/server-go/utils"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieKey := "session-token" // TODO: put it in one place
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
		claims := &users.Claims{}
		jwtKey := users.JWT_KEY
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

		log.Print("Authenticated via middleware")
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

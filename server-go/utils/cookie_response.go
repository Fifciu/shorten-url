package utils

import (
	"net/http"
	"time"
)

const (
	COOKIE_NAME_SESSION = "session-token"
)

func CookieResponse(w http.ResponseWriter, token string, expirationTime time.Time) {
	w.Header().Set("Content-Type", "application/json") // ???
	http.SetCookie(w, &http.Cookie{
		Name:    COOKIE_NAME_SESSION,
		Value:   token,
		Expires: expirationTime,
		Path:    "/",
	})
}

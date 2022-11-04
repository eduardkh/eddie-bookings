package middleware

import (
	"net/http"

	"github.com/eduardkh/eddie-bookings/pkg/config"
	"github.com/justinas/nosurf"
)

var app = &config.App

// WriteToConsole simple middleware example
//
//	func WriteToConsole(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			fmt.Println("hit the page: ", app.InProduction)
//			next.ServeHTTP(w, r)
//		})
//	}

// CSRFToken adds csrf protection to all POST requests
func CSRFToken(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionState loads and saves the session on any request
func SessionState(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}

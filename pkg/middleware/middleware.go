package middleware

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole simple middleware example
// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("hit the page: ")
// 		next.ServeHTTP(w, r)
// 	})
// }

func CSRFToken(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

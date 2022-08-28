package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Hit the page")
// 		next.ServeHTTP(w, r)
// 	})
// }

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// uses cookies to make sure the tokens generated are valid
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",   // entire site
		Secure:   false, // we are not in production for now (set to True if in production)
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)


//adds CSRF protection to all posted request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
//SessionLoad loads and saves the session on reques
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
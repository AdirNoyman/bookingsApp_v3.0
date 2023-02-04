package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {

	// create a 'nosurf' object
	csrfHandler := nosurf.New(next)

	// Create a cookie so the csrfToken will available for all the pages in our web app
	csrfHandler.SetBaseCookie(http.Cookie{
		// Cookie will be valid only for http
		HttpOnly: true,
		// Set cookie for the entire site
		Path: "/",
		// Set the cookie not to use https but only http trafic
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	// return cookie(token)
	return csrfHandler
}

// SessionLoad loads and save the session data on every request
// by calling LoadAndSave provides middleware which automatically loads and saves session data for the current request, and communicates the session token to and from the client in a cookie
func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)
}

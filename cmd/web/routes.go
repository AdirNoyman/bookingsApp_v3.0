package main

import (
	"github.com/AdirNoyman/hotelBookings/pkg/config"
	"github.com/AdirNoyman/hotelBookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	// Mux is my app router
	mux := chi.NewRouter()

	// This is ready made third party middleware
	mux.Use(middleware.Recoverer)
	// This is a custom middleware we created (csrfToken)
	mux.Use(NoSurf)
	// Load session data => this will get us access to the session data
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

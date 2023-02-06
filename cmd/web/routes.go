package main

import (
	"github.com/AdirNoyman/hotelBookings/internal/config"
	"github.com/AdirNoyman/hotelBookings/internal/handlers"
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
	mux.Get("/president", handlers.Repo.President)
	mux.Get("/royal-sweet", handlers.Repo.RoyalSweet)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

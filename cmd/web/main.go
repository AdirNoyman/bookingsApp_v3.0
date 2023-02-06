package main

import (
	"fmt"
	"github.com/AdirNoyman/hotelBookings/internal/config"
	"github.com/AdirNoyman/hotelBookings/internal/handlers"
	"github.com/AdirNoyman/hotelBookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

// Declare a variable that will my apps configuration
var app config.AppConfig

// Declare a session variable that will hold my session data
var session *scs.SessionManager

func main() {

	// Set my environment configuration to DEV
	app.InProduction = false
	// Create a new session object that will the user data we need
	session = scs.New()
	// Set how long we need the session cookie to last
	session.Lifetime = 24 * time.Hour
	// Set the session cookie to persist even after I close the browser window
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	// Save session data to my app config
	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache 😩")
	}
	app.TemplateCache = templateCache

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	// Give render access to our app config variable
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting the application on port %s 😎🤟", portNumber))

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}

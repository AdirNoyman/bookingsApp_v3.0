package handlers

import (
	"github.com/AdirNoyman/hotelBookings/pkg/config"
	"github.com/AdirNoyman/hotelBookings/pkg/models"
	"github.com/AdirNoyman/hotelBookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
// (rep *Repository) - give this function access to what is in the repository (appConfig). Making it a method of the Repository struct
func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// Grab the clients IP address and save it to the session data
	remoteIP := r.RemoteAddr
	rep.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

// About is the about page handler
// (rep *Repository) - give this function access to what is in the repository (appConfig). Making it a method of the Repository struct
func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again 🤓"

	remoteIP := rep.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// Send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

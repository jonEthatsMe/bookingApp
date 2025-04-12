package handlers

import (
	"BookingApp/pkg/config"
	"BookingApp/pkg/models"
	"BookingApp/pkg/render"
	"net/http"
)

// repo used by the handlers
var Repo *Repository

// repo structure
type Repository struct {
	App *config.AppConfig
}

// create a new repo and returns dereference repo struct
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// creates new handler and sets repo to repo variable
func NewHandlers(r *Repository) {
	Repo = r
}

// creates home page handler or path
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) { 
	remoteIP  := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// this is the about page handler or path
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello there!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	//send the template out
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

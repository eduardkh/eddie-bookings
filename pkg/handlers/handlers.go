package handlers

import (
	"net/http"

	"github.com/eduardkh/eddie-bookings/pkg/config"
	"github.com/eduardkh/eddie-bookings/pkg/models"
	"github.com/eduardkh/eddie-bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	data := make(map[string]string)
	data["data"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{
		StringMap: data,
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{})
}

package main

import (
	"log"
	"net/http"

	"github.com/eduardkh/eddie-bookings/pkg/config"
	"github.com/eduardkh/eddie-bookings/pkg/handlers"
	"github.com/eduardkh/eddie-bookings/pkg/render"
)

const port string = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Server running on port '%s'\n", port)
	http.ListenAndServe(port, nil)
}

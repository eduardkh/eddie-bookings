package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/eduardkh/eddie-bookings/pkg/config"
	"github.com/eduardkh/eddie-bookings/pkg/handlers"
	middle "github.com/eduardkh/eddie-bookings/pkg/middleware"
	"github.com/eduardkh/eddie-bookings/pkg/render"
	"github.com/eduardkh/eddie-bookings/pkg/routes"
)

const port string = ":8080"

var app = &config.App

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middle.CSRFToken)
	// router.Use(middle.WriteToConsole)
	routes.Routes(router)

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.InProduction = true

	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)

	render.NewTemplates(app)

	log.Printf("Server running on port '%s'\n", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}

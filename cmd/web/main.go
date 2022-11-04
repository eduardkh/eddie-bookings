package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
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

var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middle.CSRFToken)
	router.Use(middle.SessionState)
	// router.Use(middle.WriteToConsole)
	routes.Routes(router)

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)

	render.NewTemplates(app)

	log.Printf("Server running on port '%s'\n", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}

package routes

import (
	"github.com/eduardkh/eddie-bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func Routes(router *chi.Mux) {
	router.HandleFunc("/", handlers.Repo.Home)
	router.HandleFunc("/about", handlers.Repo.About)
}

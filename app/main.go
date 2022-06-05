package main

import (
	"net/http"

	"github.com/gba-3/tweat/handler"
	"github.com/gba-3/tweat/registry"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	ah := registry.NewRegistry()
	r.Route("/tweats", func(r chi.Router) {
		r.Get("/", handler.JsonHandler(ah.Th.GetAll).ServeHTTP)
	})
	http.ListenAndServe(":3000", r)
}

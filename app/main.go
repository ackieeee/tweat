package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sugartr3e/tweat/handler"
	mw "github.com/sugartr3e/tweat/middleware"
	"github.com/sugartr3e/tweat/registry"
)

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	ah := registry.NewRegistry()
	r.Route("/tweats", func(r chi.Router) {
		r.Use(mw.WithToken)
		r.Get("/", handler.JsonHandler(ah.Th.GetAll).ServeHTTP)
		r.Route("/like", func(r chi.Router) {
			r.Post("/add", handler.JsonHandler(ah.Th.AddLike).ServeHTTP)
			r.Post("/delete", handler.JsonHandler(ah.Th.DeleteLike).ServeHTTP)
			r.Post("/toggle", handler.JsonHandler(ah.Th.ToggleLike).ServeHTTP)
		})
	})
	r.Route("/signup", func(r chi.Router) {
		r.Post("/", handler.JsonHandler(ah.Uh.Signup).ServeHTTP)
	})
	r.Route("/login", func(r chi.Router) {
		r.Post("/", handler.JsonHandler(ah.Uh.Login).ServeHTTP)
	})
	http.ListenAndServe(":3030", r)
}

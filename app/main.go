package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gba-3/gologger"
	"github.com/gba-3/tweat/handler"
	mw "github.com/gba-3/tweat/middleware"
	"github.com/gba-3/tweat/registry"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

func main() {
	setLogger()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	ah := registry.NewRegistry()
	r.Route("/tweats", func(r chi.Router) {
		r.Use(mw.WithToken)
		r.Get("/", handler.JsonHandler(ah.Th.GetAll).ServeHTTP)
		r.Route("/like", func(r chi.Router) {
			r.Post("/add", handler.JsonHandler(ah.Th.AddLike).ServeHTTP)
			r.Post("/delete", handler.JsonHandler(ah.Th.DeleteLike).ServeHTTP)
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

func setLogger() {
	logLevel := os.Getenv("LEVEL")
	if logLevel == "" {
		logLevel = "INFO"
	}
	gologger.SetLogger(logLevel)
}

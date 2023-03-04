package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sugartr3e/tweat/handler"
	mw "github.com/sugartr3e/tweat/middleware"
	"github.com/sugartr3e/tweat/registry"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

func Server(ctx context.Context) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	ah := registry.NewRegistry()
	// ルーティング
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

	// サーバー起動
	srv := &http.Server{
		Handler: otelhttp.NewHandler(r, "",
			otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
				return r.Method + ": " + r.URL.Path
			}),
			otelhttp.WithMessageEvents(
				otelhttp.ReadEvents,
				otelhttp.WriteEvents,
			)),
		Addr: "0.0.0.0:3030",
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Server closed with error: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutdown\n", <-quit)

	// タイムアウト処理
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}

func main() {
	ctx := context.Background()
	if err := Server(ctx); err != nil {
		log.Fatal(err.Error())
	}
}

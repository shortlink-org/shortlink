package main

import (
	"fmt"
	"github.com/batazor/shortlink/pkg/api"
	additionalMiddleware "github.com/batazor/shortlink/pkg/api/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	// Run HTTP-server
	PORT := "7070"

	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	r := chi.NewRouter()

	// CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		//Debug:            true,
	})

	r.Use(cors.Handler)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.Recoverer)

	// Additional middleware
	r.Use(additionalMiddleware.Logger(logger))

	r.NotFound(NotFoundHandler)

	r.Mount("/", api.Routes())
	sugar.Info(fmt.Sprintf("Run on port %s", PORT))

	srv := http.Server{Addr: ":" + PORT, Handler: r}

	// start HTTP-server
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

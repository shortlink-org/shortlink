package httpchi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/batazor/shortlink/internal/logger"
	additionalMiddleware "github.com/batazor/shortlink/pkg/api/http-chi/middleware"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// Run HTTP-server
func (api *API) Run(ctx context.Context, config api_type.Config, log logger.Logger) error {
	api.ctx = ctx

	log.Info("Run HTTP-CHI API")

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
	r.Use(additionalMiddleware.Logger(log))

	r.NotFound(NotFoundHandler)

	r.Mount("/api", api.Routes())

	log.Info(fmt.Sprintf("Run on port %d", config.Port))
	srv := http.Server{Addr: fmt.Sprintf(":%d", config.Port), Handler: chi.ServerBaseContext(ctx, r)}

	// start HTTP-server
	err := srv.ListenAndServe()
	return err
}

// NotFoundHandler - default handler for don't existing routers
func NotFoundHandler(w http.ResponseWriter, r *http.Request) { // nolint unused
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

package singleflight_middleware

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/singleflight"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

type Option func(*singleFlight)

func WithKeyFn(keyFn func(r *http.Request) string) Option {
	return func(s *singleFlight) {
		s.keyFn = keyFn
	}
}

type singleFlight struct {
	log   logger.Logger
	group singleflight.Group
	keyFn func(r *http.Request) string
}

// SingleFlight is a middleware that prevents duplicate requests from hitting your server.
func SingleFlight(log logger.Logger, options ...Option) func(next http.Handler) http.Handler {
	// Default keyFn is path + query
	keyFn := func(r *http.Request) string {
		return fmt.Sprintf("%s?%s", r.URL.Path, r.URL.RawQuery)
	}

	sf := &singleFlight{
		log:   log,
		keyFn: keyFn,
	}

	for _, option := range options {
		option(sf)
	}

	return sf.middleware
}

func (s *singleFlight) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			key := s.keyFn(r)

			response, err, _ := s.group.Do(key, func() (any, error) {
				next.ServeHTTP(w, r)

				//nolint:nilnil // nil, nil is valid return
				return nil, nil
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			_, err = fmt.Fprint(w, response)
			if err != nil {
				s.log.Error("failed to write response", field.Fields{
					"error": err,
				})
			}
		} else {
			next.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}

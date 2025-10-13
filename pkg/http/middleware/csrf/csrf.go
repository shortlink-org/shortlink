package csrf

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/spf13/viper"
)

// Middleware creates a CSRF protection middleware using Go's built-in CrossOriginProtection
func Middleware(log logger.Logger) func(http.Handler) http.Handler {
	// Initialize CrossOriginProtection
	antiCSRF := http.NewCrossOriginProtection()

	// Configure trusted origins from environment variables
	configureTrustedOrigins(antiCSRF, log)

	// Return a middleware function that wraps the handler with CSRF protection
	return func(next http.Handler) http.Handler {
		return antiCSRF.Handler(next)
	}
}

// configureTrustedOrigins sets up trusted origins from environment variables
func configureTrustedOrigins(antiCSRF *http.CrossOriginProtection, log logger.Logger) {
	// Set default environment variable names
	viper.SetDefault("CSRF_TRUSTED_ORIGINS_ENV", "CSRF_TRUSTED_ORIGINS")
	viper.SetDefault("CSRF_TRUSTED_ORIGINS", "")

	// Get trusted origins from environment variable
	envVarName := viper.GetString("CSRF_TRUSTED_ORIGINS_ENV")
	trustedOrigins := os.Getenv(envVarName)

	// If not found in the direct env var, try viper config
	if trustedOrigins == "" {
		trustedOrigins = viper.GetString("CSRF_TRUSTED_ORIGINS")
	}

	if trustedOrigins != "" {
		origins := strings.Split(trustedOrigins, ",")
		for _, origin := range origins {
			trimmedOrigin := strings.TrimSpace(origin)
			if trimmedOrigin != "" {
				if err := antiCSRF.AddTrustedOrigin(trimmedOrigin); err != nil {
					log.Error("CSRF trusted origin configuration error", slog.String("origin", trimmedOrigin), slog.Any("error", err))
				} else {
					log.Info("CSRF trusted origin added", slog.String("origin", trimmedOrigin))
				}
			}
		}
	} else {
		log.Info("No CSRF trusted origins configured. All cross-origin requests will be protected.")
	}
}

// Config represents CSRF middleware configuration
type Config struct {
	TrustedOrigins []string
}

// New creates a new CSRF middleware with custom configuration
func New(config Config) func(http.Handler) http.Handler {
	antiCSRF := http.NewCrossOriginProtection()

	// Add trusted origins from config
	for _, origin := range config.TrustedOrigins {
		if err := antiCSRF.AddTrustedOrigin(origin); err != nil {
			log.Printf("Failed to add trusted origin %s: %v", origin, err)
		} else {
			log.Printf("Added trusted origin: %s", origin)
		}
	}

	return func(next http.Handler) http.Handler {
		return antiCSRF.Handler(next)
	}
}

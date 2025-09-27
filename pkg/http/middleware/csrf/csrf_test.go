package csrf

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		envVar         string
		envValue       string
		origin         string
		method         string
		expectedStatus int
		description    string
	}{
		{
			name:           "same_origin_request_allowed",
			envVar:         "CSRF_TRUSTED_ORIGINS",
			envValue:       "",
			origin:         "", // Same origin (no Origin header)
			method:         "POST",
			expectedStatus: http.StatusOK,
			description:    "Same-origin requests should be allowed",
		},
		{
			name:           "trusted_origin_allowed",
			envVar:         "CSRF_TRUSTED_ORIGINS",
			envValue:       "https://shortlink.best,https://www.shortlink.best",
			origin:         "https://shortlink.best",
			method:         "POST",
			expectedStatus: http.StatusOK,
			description:    "Requests from trusted origins should be allowed",
		},
		{
			name:           "untrusted_origin_blocked",
			envVar:         "CSRF_TRUSTED_ORIGINS",
			envValue:       "https://shortlink.best",
			origin:         "https://malicious.com",
			method:         "POST",
			expectedStatus: http.StatusForbidden,
			description:    "Requests from untrusted origins should be blocked",
		},
		{
			name:           "get_request_always_allowed",
			envVar:         "CSRF_TRUSTED_ORIGINS",
			envValue:       "",
			origin:         "https://malicious.com",
			method:         "GET",
			expectedStatus: http.StatusOK,
			description:    "GET requests should be allowed regardless of origin",
		},
		{
			name:           "localhost_development",
			envVar:         "CSRF_TRUSTED_ORIGINS",
			envValue:       "http://localhost:3000,http://127.0.0.1:3000",
			origin:         "http://localhost:3000",
			method:         "POST",
			expectedStatus: http.StatusOK,
			description:    "Localhost origins for development should be allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up environment
			os.Unsetenv("CSRF_TRUSTED_ORIGINS")
			viper.Reset()

			// Set up environment variable if provided
			if tt.envValue != "" {
				os.Setenv(tt.envVar, tt.envValue)
			}

			// Create a test handler
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("OK"))
			})

			// Apply CSRF middleware
			middleware := Middleware()
			protectedHandler := middleware(handler)

			// Create test request
			req := httptest.NewRequest(tt.method, "/test", nil)
			if tt.origin != "" {
				req.Header.Set("Origin", tt.origin)
			}

			// Record response
			rr := httptest.NewRecorder()
			protectedHandler.ServeHTTP(rr, req)

			// Assert response
			assert.Equal(t, tt.expectedStatus, rr.Code, tt.description)

			// Clean up
			os.Unsetenv(tt.envVar)
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name           string
		config         Config
		origin         string
		method         string
		expectedStatus int
		description    string
	}{
		{
			name: "custom_config_trusted_origin",
			config: Config{
				TrustedOrigins: []string{
					"https://shortlink.best",
					"https://api.shortlink.best",
				},
			},
			origin:         "https://shortlink.best",
			method:         "POST",
			expectedStatus: http.StatusOK,
			description:    "Custom config should allow configured trusted origins",
		},
		{
			name: "custom_config_untrusted_origin",
			config: Config{
				TrustedOrigins: []string{
					"https://shortlink.best",
				},
			},
			origin:         "https://malicious.com",
			method:         "POST",
			expectedStatus: http.StatusForbidden,
			description:    "Custom config should block untrusted origins",
		},
		{
			name: "empty_config",
			config: Config{
				TrustedOrigins: []string{},
			},
			origin:         "https://example.com",
			method:         "POST",
			expectedStatus: http.StatusOK, // Go's CrossOriginProtection allows when no origins are configured but no additional origins are added
			description:    "Empty config with no trusted origins configured",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test handler
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("OK"))
			})

			// Apply CSRF middleware with custom config
			middleware := New(tt.config)
			protectedHandler := middleware(handler)

			// Create test request
			req := httptest.NewRequest(tt.method, "/test", nil)
			req.Header.Set("Origin", tt.origin)

			// Record response
			rr := httptest.NewRecorder()
			protectedHandler.ServeHTTP(rr, req)

			// Assert response
			assert.Equal(t, tt.expectedStatus, rr.Code, tt.description)
		})
	}
}

func TestConfigureTrustedOrigins(t *testing.T) {
	tests := []struct {
		name     string
		envVar   string
		envValue string
		expected []string
	}{
		{
			name:     "single_origin",
			envVar:   "CSRF_TRUSTED_ORIGINS",
			envValue: "https://shortlink.best",
			expected: []string{"https://shortlink.best"},
		},
		{
			name:     "multiple_origins",
			envVar:   "CSRF_TRUSTED_ORIGINS",
			envValue: "https://shortlink.best,https://www.shortlink.best,http://localhost:3000",
			expected: []string{"https://shortlink.best", "https://www.shortlink.best", "http://localhost:3000"},
		},
		{
			name:     "origins_with_spaces",
			envVar:   "CSRF_TRUSTED_ORIGINS",
			envValue: " https://shortlink.best , https://www.shortlink.best , http://localhost:3000 ",
			expected: []string{"https://shortlink.best", "https://www.shortlink.best", "http://localhost:3000"},
		},
		{
			name:     "empty_value",
			envVar:   "CSRF_TRUSTED_ORIGINS",
			envValue: "",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up
			os.Unsetenv("CSRF_TRUSTED_ORIGINS")
			viper.Reset()

			// Set environment variable
			if tt.envValue != "" {
				os.Setenv(tt.envVar, tt.envValue)
			}

			// Create a mock CrossOriginProtection to test configuration
			antiCSRF := http.NewCrossOriginProtection()

			// Call the function under test
			configureTrustedOrigins(antiCSRF)

			// Verify that origins were configured (we can't directly inspect them,
			// but we can test the behavior with a handler)
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
			protectedHandler := antiCSRF.Handler(handler)

			// Test each expected origin
			for _, expectedOrigin := range tt.expected {
				if expectedOrigin != "" {
					req := httptest.NewRequest("POST", "/test", nil)
					req.Header.Set("Origin", expectedOrigin)
					rr := httptest.NewRecorder()
					protectedHandler.ServeHTTP(rr, req)
					
					// If the origin was properly configured, it should be allowed
					assert.Equal(t, http.StatusOK, rr.Code, 
						"Expected origin %s to be allowed", expectedOrigin)
				}
			}

			// Clean up
			os.Unsetenv(tt.envVar)
		})
	}
}

func TestCustomEnvironmentVariable(t *testing.T) {
	// Test using a custom environment variable name
	customEnvVar := "MY_TRUSTED_ORIGINS"
	customValue := "https://shortlink.best,https://api.shortlink.best"

	// Clean up
	os.Unsetenv("CSRF_TRUSTED_ORIGINS")
	os.Unsetenv(customEnvVar)
	viper.Reset()

	// Set custom environment variable name in viper
	viper.Set("CSRF_TRUSTED_ORIGINS_ENV", customEnvVar)
	os.Setenv(customEnvVar, customValue)

	// Create test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Apply middleware
	middleware := Middleware()
	protectedHandler := middleware(handler)

	// Test with one of the configured origins
	req := httptest.NewRequest("POST", "/test", nil)
	req.Header.Set("Origin", "https://shortlink.best")
	rr := httptest.NewRecorder()
	protectedHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, 
		"Should allow origin from custom environment variable")

	// Clean up
	os.Unsetenv(customEnvVar)
	viper.Reset()
}

func TestViperConfiguration(t *testing.T) {
	// Test configuration via viper instead of environment variable
	viper.Reset()
	viper.Set("CSRF_TRUSTED_ORIGINS", "https://shortlink.best,https://api.shortlink.best")

	// Create test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Apply middleware
	middleware := Middleware()
	protectedHandler := middleware(handler)

	// Test with configured origin
	req := httptest.NewRequest("POST", "/test", nil)
	req.Header.Set("Origin", "https://shortlink.best")
	rr := httptest.NewRecorder()
	protectedHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, 
		"Should allow origin configured via viper")

	viper.Reset()
}

// Benchmark tests
func BenchmarkMiddleware(b *testing.B) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	middleware := Middleware()
	protectedHandler := middleware(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	b.ResetTimer()
	for b.Loop() {
		protectedHandler.ServeHTTP(rr, req)
	}
}

func BenchmarkMiddlewareWithOrigin(b *testing.B) {
	os.Setenv("CSRF_TRUSTED_ORIGINS", "https://shortlink.best")
	defer os.Unsetenv("CSRF_TRUSTED_ORIGINS")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	middleware := Middleware()
	protectedHandler := middleware(handler)

	req := httptest.NewRequest("POST", "/test", nil)
	req.Header.Set("Origin", "https://shortlink.best")
	rr := httptest.NewRecorder()

	b.ResetTimer()
	for b.Loop() {
		protectedHandler.ServeHTTP(rr, req)
	}
}
# CSRF Protection Middleware

This package provides CSRF (Cross-Site Request Forgery) protection using Go 1.25.1's built-in `http.CrossOriginProtection` functionality.

## Features

- Built-in CSRF protection using Go's standard library
- Environment variable configuration for trusted origins
- Support for multiple trusted origins via comma-separated values
- Integration with existing middleware stack
- Proper CORS header configuration

## Configuration

### Environment Variables

The middleware can be configured using the following environment variables:

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `CSRF_TRUSTED_ORIGINS` | Comma-separated list of trusted origins | "" | `https://example.com,https://www.example.com` |
| `CSRF_TRUSTED_ORIGINS_ENV` | Name of environment variable containing trusted origins | `CSRF_TRUSTED_ORIGINS` | `TRUSTED_ORIGINS` |

### Example Configuration

```bash
# Set trusted origins directly
export CSRF_TRUSTED_ORIGINS="https://shortlink.best,https://www.shortlink.best,https://api.shortlink.best"

# Development with localhost
export CSRF_TRUSTED_ORIGINS="http://localhost:3000,http://127.0.0.1:3000,https://localhost:3000,https://127.0.0.1:3000"

# Or use a custom environment variable name
export CSRF_TRUSTED_ORIGINS_ENV="MY_TRUSTED_ORIGINS"
export MY_TRUSTED_ORIGINS="https://shortlink.best,https://www.shortlink.best"
```

## Usage

### Basic Usage

```go
import csrf_middleware "github.com/shortlink-org/shortlink/pkg/http/middleware/csrf"

// Add to your router
r.Use(csrf_middleware.Middleware())
```

### Custom Configuration

```go
import csrf_middleware "github.com/shortlink-org/shortlink/pkg/http/middleware/csrf"

// Create middleware with custom config
csrfMiddleware := csrf_middleware.New(csrf_middleware.Config{
    TrustedOrigins: []string{
        "https://shortlink.best",
        "https://www.shortlink.best",
        "http://localhost:3000",
        "http://127.0.0.1:3000",
    },
})

r.Use(csrfMiddleware)
```

## How It Works

1. The middleware wraps your HTTP handlers with Go's `CrossOriginProtection`
2. For each request, it validates the origin against configured trusted origins
3. Cross-origin requests from untrusted origins are blocked unless they meet CSRF safety criteria
4. Same-origin requests and requests from trusted origins are allowed through

For more details on Go's built-in CSRF protection, see the [official Go documentation](https://pkg.go.dev/net/http#NewCrossOriginProtection).

## CORS Headers

When using this middleware, ensure your CORS configuration includes:

```go
AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"}
ExposedHeaders: []string{"X-CSRF-Token"}
```

## Security Notes

- Always configure trusted origins in production environments
- Use HTTPS origins only in production
- The middleware should be placed early in the middleware stack
- Ensure CORS is properly configured to work with CSRF protection

## Integration with Existing Code

This middleware is already integrated into the BFF service. To use it in other services:

1. Import the package
2. Add the middleware to your router before authentication middleware
3. Configure trusted origins via environment variables
4. Update CORS headers if needed
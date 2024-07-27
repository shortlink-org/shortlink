package auth_middleware

import (
	"context"
	"fmt"
	"net/http"

	ory "github.com/ory/client-go"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/pkg/auth/session"
)

const (
	// ContextCookieKey is the key used to store the cookie in the context.
	contextCookieKey = "cookie"
)

type auth struct {
	ory *ory.APIClient
}

// Auth returns a new Auth Middleware handler.
func Auth() func(next http.Handler) http.Handler {
	viper.SetDefault("AUTH_URI", "http://127.0.0.1:4433")

	c := ory.NewConfiguration()
	c.Servers = ory.ServerConfigurations{{URL: viper.GetString("AUTH_URI")}}

	return auth{
		ory: ory.NewAPIClient(c),
	}.middleware
}

func (a auth) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Header.Get("Cookie")

		// check if the cookie is valid
		sess, _, err := a.ory.FrontendAPI.ToSession(r.Context()).Cookie(cookies).Execute() //nolint:bodyclose // false positive
		if (err != nil && sess == nil) || (err == nil && !*sess.Active) {
			// this will redirect the user to the managed Ory Login UI
			// NOTE:
			// 	- we use 302 instead of 303 because proxy servers might not understand the 303 status code
			// details -> https://stackoverflow.com/questions/2839585/what-is-correct-http-status-code-when-redirecting-to-a-login-page
			http.Redirect(w, r, fmt.Sprintf("%s/self-service/login/browser", viper.GetString("AUTH_URI")), http.StatusFound)
			return
		}

		ctx := withCookie(r.Context(), cookies)
		ctx = session.WithSession(ctx, sess)

		// set the new context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func withCookie(ctx context.Context, cookie string) context.Context {
	return context.WithValue(ctx, contextCookieKey, cookie) //nolint:staticcheck // TODO: fix
}

func GetCookie(ctx context.Context) string {
	return ctx.Value(contextCookieKey).(string) //nolint:forcetypeassert // simple type assertion
}

package middleware

import (
	"context"
	"fmt"
	"net/http"

	ory "github.com/ory/client-go"
	"github.com/spf13/viper"
)

const (
	// contextCookieKey is the key used to store the cookie in the context.
	contextCookieKey = "cookie"
	// contextSessionKey is the key used to store the session in the context.
	contextSessionKey = "session"
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
		session, _, err := a.ory.FrontendApi.ToSession(r.Context()).Cookie(cookies).Execute()
		if (err != nil && session == nil) || (err == nil && !*session.Active) {
			// this will redirect the user to the managed Ory Login UI
			http.Redirect(w, r, fmt.Sprintf("%s/self-service/login/browser", viper.GetString("AUTH_URI")), http.StatusSeeOther)
			return
		}

		ctx := withCookie(r.Context(), cookies)
		ctx = withSession(ctx, session)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func withCookie(ctx context.Context, cookie string) context.Context {
	return context.WithValue(ctx, contextCookieKey, cookie)
}

func GetCookie(ctx context.Context) string {
	return ctx.Value(contextCookieKey).(string)
}

func withSession(ctx context.Context, session *ory.Session) context.Context {
	return context.WithValue(ctx, contextSessionKey, session)
}

func GetSession(ctx context.Context) *ory.Session {
	return ctx.Value(contextSessionKey).(*ory.Session)
}

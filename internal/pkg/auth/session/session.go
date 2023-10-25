package session

import (
	"context"

	ory "github.com/ory/client-go"
)

type Session string

const (
	// ContextSessionKey is the key used to store the session in the context.
	contextSessionKey = Session("session")
)

func WithSession(ctx context.Context, session *ory.Session) context.Context {
	return context.WithValue(ctx, contextSessionKey, session)
}

func GetSession(ctx context.Context) *ory.Session {
	sess := ctx.Value(contextSessionKey)
	if sess == nil {
		return nil
	}

	return sess.(*ory.Session)
}

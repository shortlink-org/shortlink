package session

import (
	"context"

	ory "github.com/ory/client-go"
)

type Session string

const (
	// ContextSessionKey is the key used to store the session in the context.
	contextSessionKey = Session("session")

	// ContextUserIDKey is the key used to store the user id in the context.
	ContextUserIDKey = Session("user-id")
)

// String returns the string representation of the session.
func (s Session) String() string {
	return string(s)
}

func WithSession(ctx context.Context, session *ory.Session) context.Context {
	return context.WithValue(ctx, contextSessionKey, session)
}

func GetSession(ctx context.Context) (*ory.Session, error) {
	sess := ctx.Value(contextSessionKey)
	if sess == nil {
		return nil, ErrSessionNotFound
	}

	if session, ok := sess.(*ory.Session); ok {
		return session, nil
	}

	return nil, ErrSessionNotFound
}

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ContextUserIDKey, userID)
}

func GetUserID(ctx context.Context) (string, error) {
	userID := ctx.Value(ContextUserIDKey)
	if userID == nil {
		return "", ErrUserIDNotFound
	}

	if uid, ok := userID.(string); ok {
		return uid, nil
	}

	return "", ErrUserIDNotFound
}

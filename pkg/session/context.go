package session

import (
	"calendorario/pkg/database"
	"context"
)

type sessionContext struct {
	Database          *database.Queries
	AuthenticatedUser *database.User
	AuthenticationErr error
}

type sessionContextKey struct{}

func FromContext(ctx context.Context) sessionContext {
	return ctx.Value(sessionContextKey{}).(sessionContext)
}

func New(
	database *database.Queries,
	authenticatedUser *database.User,
	authenticationErr error,
) sessionContext {
	return sessionContext{database, authenticatedUser, authenticationErr}
}

func NewContext(
	ctx context.Context,
	database *database.Queries,
	authenticatedUser *database.User,
	authenticationErr error,
) context.Context {
	return context.WithValue(
		ctx,
		sessionContextKey{},
		sessionContext{database, authenticatedUser, authenticationErr},
	)
}

func (s *sessionContext) User() (*database.User, error) {
	return s.AuthenticatedUser, s.AuthenticationErr
}

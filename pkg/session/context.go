package session

import (
	"calendorario/pkg/database"
	"context"
)

type sessionContext struct {
	Database          *database.Queries
	AuthenticatedUser *database.User
	AuthenticationErr error
	TermID            int
}

type sessionContextKey struct{}

func FromContext(ctx context.Context) sessionContext {
	return ctx.Value(sessionContextKey{}).(sessionContext)
}

func New(
	database *database.Queries,
	authenticatedUser *database.User,
	authenticationErr error,
	termID int,
) sessionContext {
	return sessionContext{
		database,
		authenticatedUser,
		authenticationErr,
		termID,
	}
}

func NewContext(
	ctx context.Context,
	database *database.Queries,
	authenticatedUser *database.User,
	authenticationErr error,
	termID int,
) context.Context {
	return context.WithValue(
		ctx,
		sessionContextKey{},
		sessionContext{
			database,
			authenticatedUser,
			authenticationErr,
			termID,
		},
	)
}

func (s *sessionContext) User() (*database.User, error) {
	return s.AuthenticatedUser, s.AuthenticationErr
}

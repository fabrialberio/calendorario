package session

import (
	"calendorario/pkg/database"
	"context"
	"time"
)

type sessionContext struct {
	Database          *database.Queries
	AuthenticatedUser *database.User
	AuthenticationErr error
	SelectedTermID    int
	SelectedDate      *time.Time
}

type sessionContextKey struct{}

func FromContext(ctx context.Context) sessionContext {
	return ctx.Value(sessionContextKey{}).(sessionContext)
}

func NewContext(
	ctx context.Context,
	database *database.Queries,
	authenticatedUser *database.User,
	authenticationErr error,
	selectedTermID int,
	selectedDate *time.Time,
) context.Context {
	return context.WithValue(
		ctx,
		sessionContextKey{},
		sessionContext{
			database,
			authenticatedUser,
			authenticationErr,
			selectedTermID,
			selectedDate,
		},
	)
}

func (s *sessionContext) User() (*database.User, error) {
	return s.AuthenticatedUser, s.AuthenticationErr
}

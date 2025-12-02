package session

import (
	"calendorario/pkg/database"
	"context"
	"time"
)

type sessionContext struct {
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
	authenticatedUser *database.User,
	authenticationErr error,
	selectedTermID int,
	selectedDate *time.Time,
) context.Context {
	return context.WithValue(
		ctx,
		sessionContextKey{},
		sessionContext{
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

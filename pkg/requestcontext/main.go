package requestcontext

import (
	"calendorario/pkg/database"
	"context"
)

type requestContext struct {
	Database          *database.Queries
	AuthenticatedUser *database.User
	AuthenticationErr error
}

type requestContextKey struct{}

func FromContext(ctx context.Context) requestContext {
	return ctx.Value(requestContextKey{}).(requestContext)
}

func New(
	database *database.Queries,
	authenticatedUser *database.User,
	authenticationErr error,
) requestContext {
	return requestContext{database, authenticatedUser, authenticationErr}
}

func NewContext(
	ctx context.Context,
	database *database.Queries,
	authenticatedUser *database.User,
	authenticationErr error,
) context.Context {
	return context.WithValue(
		ctx,
		requestContextKey{},
		requestContext{database, authenticatedUser, authenticationErr},
	)
}

func (rc *requestContext) User() (*database.User, error) {
	return rc.AuthenticatedUser, rc.AuthenticationErr
}

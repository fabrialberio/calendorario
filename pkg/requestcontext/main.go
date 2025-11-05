package requestcontext

import (
	"calendorario/pkg/database"
	"context"
	"net/http"
)

type requestContext struct {
	database          *database.Queries
	authenticatedUser *database.User
	authenticationErr error
}

type requestContextKey struct{}

func FromRequest(r *http.Request) requestContext {
	return r.Context().Value(requestContextKey{}).(requestContext)
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

func (rc *requestContext) Database() *database.Queries {
	return rc.database
}

func (rc *requestContext) AuthenticatedUser() (*database.User, error) {
	return rc.authenticatedUser, rc.authenticationErr
}

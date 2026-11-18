package newterm

import (
	"calendorario/pkg/database"
	"calendorario/routes"
	"calendorario/routes/newterm/step1"
	"net/http"
)

const (
	keyID        = routes.KeyID
	keyName      = "name"
	keyStartDate = "start_date"
	keyEndDate   = "end_date"
)

const (
	routeStep1 = routes.RouteNewTerm + "/step1"
	routeStep2 = routes.RouteNewTerm + "/step2"
	routeStep3 = routes.RouteNewTerm + "/step3"
)

func NewHandler(database *database.Queries) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(routeStep1, &step1.Handler{Database: database})

	return mux
}

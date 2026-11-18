package step1

import (
	"calendorario/pkg/database"
	"net/http"
)

type Handler struct {
	Database *database.Queries
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

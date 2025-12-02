package week

import (
	"calendorario/pkg/session"
	"net/http"
	"time"
)

const keyDate = "date"

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse(time.DateOnly, r.FormValue(keyDate))
	if err != nil {
		date = time.Now()
	}

	s := session.FromContext(r.Context())
	term, _ := s.Database.GetTerm(r.Context(), int64(s.TermID))

	innerView(date, time.Now(), term).Render(r.Context(), w)
}

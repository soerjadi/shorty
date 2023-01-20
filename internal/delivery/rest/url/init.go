package url

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soerjadi/short/internal/delivery/rest"
	"github.com/soerjadi/short/internal/usecase/url"
)

func NewHandler(usecase url.Usecase) rest.API {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	// url := r.PathPrefix("/")

	// url := r.PathPrefix("/generate").Subrouter()

	r.HandleFunc("/generate", rest.HandlerFunc(h.Generate).Serve).Methods(http.MethodPost)
	r.HandleFunc("/{shortURL}/", h.Redirect).Methods(http.MethodGet)

}

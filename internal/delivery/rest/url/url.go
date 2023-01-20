package url

import (
	"errors"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/soerjadi/short/internal/model"
)

func (h *Handler) Generate(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	param := r.URL.Query()

	var longUrl string
	var err error

	longUrlParam := param.Get("url")

	if longUrlParam == "" {
		return nil, errors.New("url empty")
	}

	longUrl = longUrlParam

	req := model.URL{
		LongURL: longUrl,
	}

	res, err := h.usecase.InsertURL(r.Context(), req)

	if err != nil {
		return nil, err
	}

	return res, err

}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	res, err := h.usecase.GetShortURL(r.Context(), shortURL)

	if err != nil {
		p := path.Dir("../../public/index.html")

		http.ServeFile(w, r, p)
	}

	http.Redirect(w, r, res.LongURL, http.StatusMovedPermanently)
}

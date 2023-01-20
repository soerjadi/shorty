package url

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/soerjadi/short/internal/model"
)

type Repository interface {
	GetLongURL(ctx context.Context, longURL string) (model.URL, error)
	GetShortURL(ctx context.Context, shortURL string) (model.URL, error)
	InsertShortURL(ctx context.Context, url model.URL) (model.URL, error)
	DeleteShortURL(ctx context.Context, shortURL string) error
	GetListedURL(ctx context.Context, req model.URLRequest) ([]model.URL, error)
}

type urlRepository struct {
	query prepareQuery
	DB    *sqlx.DB
}

type buildQueryParam struct {
	query string
	req   model.URLRequest
}

type buildQueryResult struct {
	query string
	args  []interface{}
}

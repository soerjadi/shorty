package url

import (
	"context"

	"github.com/soerjadi/short/internal/config"
	"github.com/soerjadi/short/internal/model"
	"github.com/soerjadi/short/internal/repository/url"
)

type Usecase interface {
	GetShortURL(ctx context.Context, shortURL string) (model.URL, error)
	GetLongURL(ctx context.Context, longURL string) (model.URL, error)
	GetListedURL(ctx context.Context, req model.URLRequest) ([]model.URL, error)
	DeleteURL(ctx context.Context, shortURL string) error
	InsertURL(ctx context.Context, req model.URL) (model.URL, error)
}

type urlUsecase struct {
	repository url.Repository
	config     *config.Config
}

package url

import (
	"github.com/soerjadi/short/internal/config"
	"github.com/soerjadi/short/internal/repository/url"
)

func GetUsecase(repo url.Repository, cfg *config.Config) Usecase {
	return &urlUsecase{
		repository: repo,
		config:     cfg,
	}
}

package url

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/soerjadi/short/internal/model"
	"github.com/soerjadi/short/internal/pkg/util"
)

func (u *urlUsecase) GetShortURL(ctx context.Context, shortURL string) (model.URL, error) {

	res, err := u.repository.GetShortURL(ctx, shortURL)
	if err != nil {
		return model.URL{}, err
	}

	return res, nil

}

func (u *urlUsecase) GetLongURL(ctx context.Context, longURL string) (model.URL, error) {

	res, err := u.repository.GetLongURL(ctx, longURL)
	if err != nil {
		return model.URL{}, err
	}

	return res, nil

}

func (u *urlUsecase) GetListedURL(ctx context.Context, req model.URLRequest) ([]model.URL, error) {

	res, err := u.repository.GetListedURL(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (u *urlUsecase) DeleteURL(ctx context.Context, shortURL string) error {

	err := u.repository.DeleteShortURL(ctx, shortURL)
	if err != nil {
		return err
	}

	return nil

}

func (u *urlUsecase) InsertURL(ctx context.Context, req model.URL) (model.URL, error) {

	url, err := u.GetLongURL(ctx, req.LongURL)
	if err != nil && err != sql.ErrNoRows {
		return model.URL{}, err
	}

	/**
	if longURL is already exists then return the existing shortURL
	*/
	var urlModel model.URL
	if url != urlModel {
		url.FullShortURL = fmt.Sprintf("%s/%s", u.config.Server.BaseURL, url.ShortURL)
		return url, nil
	}

	generateUrl, err := util.GenerateShortURL(ctx, req.LongURL)
	if err != nil {
		return model.URL{}, err
	}

	req.ID = generateUrl.ID
	req.ShortURL = generateUrl.ShortURL
	req.Domain = generateUrl.Domain
	req.DomainExt = generateUrl.DomainExt
	req.ClickCount = 0
	req.FullShortURL = fmt.Sprintf("%s/%s", u.config.Server.BaseURL, generateUrl.ShortURL)

	res, err := u.repository.InsertShortURL(ctx, req)
	if err != nil {
		return model.URL{}, err
	}

	res.FullShortURL = req.FullShortURL

	return res, nil

}

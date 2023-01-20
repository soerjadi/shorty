package url

import (
	"context"
	"fmt"

	"github.com/soerjadi/short/internal/model"
)

func (r *urlRepository) GetLongURL(ctx context.Context, longURL string) (model.URL, error) {
	var res model.URL

	err := r.query.getLongURL.GetContext(ctx, &res, longURL)
	if err != nil {
		fmt.Errorf("repository.url.GetLongURL failed to get context. err = %v, longURL = %v", err, longURL)

		return model.URL{}, err
	}

	return res, nil
}

func (r *urlRepository) GetShortURL(ctx context.Context, shortURL string) (model.URL, error) {
	var res model.URL

	err := r.query.getShortURL.GetContext(ctx, &res, shortURL)
	if err != nil {
		fmt.Errorf("repository.url.GetShortURL failed to get context. err = %v, shortURL = %v", err, shortURL)

		return model.URL{}, err
	}

	return res, nil
}

func (r *urlRepository) GetListedURL(ctx context.Context, req model.URLRequest) ([]model.URL, error) {
	var dbRes []model.URL

	param := buildQueryParam{
		query: getListedURL,
		req:   req,
	}

	q := buildQueries(ctx, param)

	err := r.DB.SelectContext(ctx, &dbRes, q.query, q.args...)
	if err != nil {
		fmt.Errorf("repository.url.GetListedURL fail select context. err = %v, args = %v, query = %v", err, q.args, q.query)

		return []model.URL{}, err
	}

	return dbRes, nil
}

func (r *urlRepository) InsertShortURL(ctx context.Context, url model.URL) (model.URL, error) {
	var (
		err error
		res model.URL
	)

	if err = r.query.insertQuery.GetContext(
		ctx,
		&res,
		url.ID,
		url.LongURL,
		url.ShortURL,
		url.Domain,
		url.DomainExt,
		url.IsSSL,
		url.IsAliased,
	); err != nil {
		fmt.Errorf("repository.url.InsertShortURL failed insert short url. err = %v, request = %v", err, url)

		return model.URL{}, err
	}

	return res, nil
}

func (r *urlRepository) DeleteShortURL(ctx context.Context, shortURL string) error {
	var err error

	if _, err = r.query.deleteURL.ExecContext(ctx, shortURL); err != nil {
		fmt.Errorf("repository.url.DeleteShortURL failed delete url. err = %v, shortURL = %v", err, shortURL)
		return err
	}

	return nil
}

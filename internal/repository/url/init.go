package url

import (
	"github.com/jmoiron/sqlx"
)

func prepareQueries(db *sqlx.DB) (prepareQuery, error) {
	var (
		err error
		q   prepareQuery
	)

	q.insertQuery, err = db.Preparex(insertQuery)
	if err != nil {
		return q, err
	}

	q.getLongURL, err = db.Preparex(getLongURL)
	if err != nil {
		return q, err
	}

	q.getShortURL, err = db.Preparex(getShortURL)
	if err != nil {
		return q, err
	}

	q.getListedURL, err = db.Preparex(getListedURL)
	if err != nil {
		return q, err
	}

	q.deleteURL, err = db.Preparex(deleteURL)
	if err != nil {
		return q, err
	}

	return q, err

}

func GetRepository(db *sqlx.DB) (Repository, error) {
	query, err := prepareQueries(db)
	if err != nil {
		return nil, err
	}

	return &urlRepository{
		DB:    db,
		query: query,
	}, nil
}

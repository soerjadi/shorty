package url

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	insertQuery  *sqlx.Stmt
	getLongURL   *sqlx.Stmt
	getShortURL  *sqlx.Stmt
	getListedURL *sqlx.Stmt
	deleteURL    *sqlx.Stmt
}

const (
	insertQuery = `
	INSERT INTO url(
		id,
		long_url,
		short_url,
		domain,
		domain_ext,
		is_ssl,
		is_aliased,
		created_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		NOW()
	) RETURNING id, long_url, short_url, domain, domain_ext, is_ssl, is_aliased, click_count, created_at
	`

	getLongURL = `
	SELECT * FROM url WHERE long_url = $1 LIMIT 1;
	`

	getShortURL = `
	SELECT * FROM url WHERE short_url = $1 LIMIT 1;
	`

	getListedURL = `
	SELECT * FROM url;
	`

	deleteURL = `
	DELETE FROM url WHERE short_url = $1 OR long_url = $1;
	`
)

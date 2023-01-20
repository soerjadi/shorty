package url

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

type prepareQueryMock struct {
	insertQueryMock *sqlmock.ExpectedPrepare
	getLongURL      *sqlmock.ExpectedPrepare
	getShortURL     *sqlmock.ExpectedPrepare
	getListedURL    *sqlmock.ExpectedPrepare
	deleteURL       *sqlmock.ExpectedPrepare
}

func expectPrepareMock(mock sqlmock.Sqlmock) prepareQueryMock {
	prepareQueryMock := prepareQueryMock{}

	prepareQueryMock.insertQueryMock = mock.ExpectPrepare(`
	INSERT INTO url\(
		id,
		long_url,
		short_url,
		domain,
		domain_ext,
		is_ssl,
		is_aliased,
		created_at
	\) VALUES \(
		(.+),
		(.+),
		(.+),
		(.+),
		(.+),
		(.+),
		(.+),
		NOW\(\)
	\) RETURNING id, long_url, short_url, domain, domain_ext, is_ssl, is_aliased, click_count, created_at
	`)

	prepareQueryMock.getLongURL = mock.ExpectPrepare(`
	SELECT \* FROM url WHERE long_url = (.+) LIMIT 1\;
	`)

	prepareQueryMock.getShortURL = mock.ExpectPrepare(`
	SELECT \* FROM url WHERE short_url = (.+) LIMIT 1\;
	`)

	prepareQueryMock.getListedURL = mock.ExpectPrepare(`
	SELECT \* FROM url\;
	`)

	prepareQueryMock.deleteURL = mock.ExpectPrepare(`
	DELETE FROM url WHERE short_url = (.+) OR long_url = (.+)\;
	`)

	return prepareQueryMock
}

func TestGetRepository(t *testing.T) {
	tests := []struct {
		name     string
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     func(db *sqlx.DB) Repository
		wantErr  bool
	}{
		{
			name: "success",
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				expectPrepareMock(mock)
				expectPrepareMock(mock)
				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: func(db *sqlx.DB) Repository {
				q, _ := prepareQueries(db)

				return &urlRepository{
					query: q,
					DB:    db,
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()

			got, err := GetRepository(db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want := tt.want(db)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetRepository() = %v, want %v", got, want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}

}

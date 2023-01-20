package url

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/soerjadi/short/internal/model"
)

func TestGeListedURL(t *testing.T) {
	type args struct {
		ctx context.Context
		req model.URLRequest
	}
	query := "long-url"
	isSSL := false
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     []model.URL
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: model.URLRequest{},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getListedURL.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"long_url",
					"short_url",
					"domain",
					"domain_ext",
					"is_ssl",
					"is_aliased",
					"created_at",
				}).
					AddRow(1, "long-url", "short-url", "domain", "com", false, false, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)).
					AddRow(2, "long-url2", "short-url2", "domain2", "com", false, false, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: []model.URL{
				{
					ID:        1,
					LongURL:   "long-url",
					ShortURL:  "short-url",
					Domain:    "domain",
					DomainExt: "com",
					IsSSL:     false,
					IsAliased: false,
					CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:        2,
					LongURL:   "long-url2",
					ShortURL:  "short-url2",
					Domain:    "domain2",
					DomainExt: "com",
					IsSSL:     false,
					IsAliased: false,
					CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name: "success with query applied",
			args: args{
				ctx: context.TODO(),
				req: model.URLRequest{
					Query: &query,
					SSL:   &isSSL,
				},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getListedURL.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"long_url",
					"short_url",
					"domain",
					"domain_ext",
					"is_ssl",
					"is_aliased",
					"created_at",
				}).
					AddRow(1, "long-url", "short-url", "domain", "com", false, false, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: []model.URL{
				{
					ID:        1,
					LongURL:   "long-url",
					ShortURL:  "short-url",
					Domain:    "domain",
					DomainExt: "com",
					IsSSL:     false,
					IsAliased: false,
					CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				req: model.URLRequest{},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getListedURL.ExpectQuery().WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want:    []model.URL{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := urlRepository{
				query: q,
				DB:    db,
			}
			got, err := re.GetListedURL(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetListedURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetListedURL() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}

func TestGetLongURL(t *testing.T) {
	type args struct {
		ctx     context.Context
		longUrl string
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     model.URL
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx:     context.TODO(),
				longUrl: "long-url",
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getLongURL.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"long_url",
					"short_url",
					"domain",
					"domain_ext",
					"is_ssl",
					"is_aliased",
					"created_at",
				}).
					AddRow(1, "long-url", "short-url", "domain", "com", false, false, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: model.URL{
				ID:        1,
				LongURL:   "long-url",
				ShortURL:  "short-url",
				Domain:    "domain",
				DomainExt: "com",
				IsSSL:     false,
				IsAliased: false,
				CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx:     context.TODO(),
				longUrl: "long-url",
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getLongURL.ExpectQuery().WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want:    model.URL{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := urlRepository{
				query: q,
				DB:    db,
			}
			got, err := re.GetLongURL(tt.args.ctx, tt.args.longUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetLongURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetLongURL() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}

func TestGetShortURL(t *testing.T) {
	type args struct {
		ctx      context.Context
		shortURL string
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     model.URL
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx:      context.TODO(),
				shortURL: "short-url",
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getShortURL.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"long_url",
					"short_url",
					"domain",
					"domain_ext",
					"is_ssl",
					"is_aliased",
					"created_at",
				}).
					AddRow(1, "long-url", "short-url", "domain", "com", false, false, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: model.URL{
				ID:        1,
				LongURL:   "long-url",
				ShortURL:  "short-url",
				Domain:    "domain",
				DomainExt: "com",
				IsSSL:     false,
				IsAliased: false,
				CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx:      context.TODO(),
				shortURL: "short-url",
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getShortURL.ExpectQuery().WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want:    model.URL{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := urlRepository{
				query: q,
				DB:    db,
			}
			got, err := re.GetShortURL(tt.args.ctx, tt.args.shortURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetShortURL() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}

func TestInsertShortURL(t *testing.T) {
	type args struct {
		ctx context.Context
		req model.URL
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     model.URL
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: model.URL{
					ID:        1,
					LongURL:   "long-url",
					ShortURL:  "short-url",
					Domain:    "domain",
					DomainExt: "com",
					IsSSL:     false,
					IsAliased: false,
				},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.insertQueryMock.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"long_url",
					"short_url",
					"domain",
					"domain_ext",
					"is_ssl",
					"is_aliased",
					"created_at",
				}).
					AddRow(1, "long-url", "short-url", "domain", "com", false, false, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: model.URL{
				ID:        1,
				LongURL:   "long-url",
				ShortURL:  "short-url",
				Domain:    "domain",
				DomainExt: "com",
				IsSSL:     false,
				IsAliased: false,
				CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				req: model.URL{
					ID:        1,
					LongURL:   "long-url",
					ShortURL:  "short-url",
					Domain:    "domain",
					DomainExt: "com",
					IsSSL:     false,
				},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.insertQueryMock.ExpectQuery().WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := urlRepository{
				query: q,
			}
			got, err := re.InsertShortURL(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.InsertShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.InsertShortURL() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}

func TestDeleteShortURL(t *testing.T) {
	type args struct {
		ctx      context.Context
		shortURL string
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx:      context.TODO(),
				shortURL: "string",
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.deleteURL.ExpectExec().WithArgs(
					"string",
				).WillReturnResult(sqlmock.NewResult(1, 1))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx:      context.TODO(),
				shortURL: "string",
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.deleteURL.ExpectExec().WithArgs(
					"string",
				).WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := urlRepository{
				query: q,
			}
			err := re.DeleteShortURL(tt.args.ctx, tt.args.shortURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.DeleteShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}

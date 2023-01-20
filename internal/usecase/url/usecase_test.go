package url

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/soerjadi/short/internal/mocks"
	"github.com/soerjadi/short/internal/model"
	"github.com/soerjadi/short/internal/repository/url"
)

func TestGetShortURL(t *testing.T) {
	type fields struct {
		repository url.Repository
	}

	type args struct {
		ctx      context.Context
		shortURL string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.URL
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:      context.TODO(),
				shortURL: "short-url",
			},
			fields: fields{
				repository: func() url.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockUrlRepository(ctrl)

					mock.EXPECT().
						GetShortURL(gomock.Any(), gomock.Any()).
						Return(model.URL{
							ID:        123,
							ShortURL:  "short-url",
							LongURL:   "long-url",
							Domain:    "domain",
							DomainExt: "domainext",
						}, nil)

					return mock
				}(),
			},
			want: model.URL{
				ID:        123,
				ShortURL:  "short-url",
				LongURL:   "long-url",
				Domain:    "domain",
				DomainExt: "domainext",
			},
		},
		{
			name: "failed",
			args: args{
				ctx:      context.TODO(),
				shortURL: "short-url",
			},
			fields: fields{
				repository: func() url.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockUrlRepository(ctrl)

					mock.EXPECT().
						GetShortURL(gomock.Any(), gomock.Any()).
						Return(model.URL{}, errors.New("mock failed"))

					return mock
				}(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &urlUsecase{
				repository: tt.fields.repository,
			}

			got, err := usecase.GetShortURL(tt.args.ctx, tt.args.shortURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetShortURL err = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.GetShortURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

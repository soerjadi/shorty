package util

import (
	"context"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	type args struct {
		ctx     context.Context
		longUrl string
	}

	tests := []struct {
		name    string
		args    args
		init    func() GenerateResponse
		want    GenerateResponse
		wantErr bool
	}{
		{
			name: "success generating 7 characters",
			args: args{
				ctx:     context.TODO(),
				longUrl: "https://soerja.com/hello-world/",
			},
			want: GenerateResponse{
				Domain:    "soerja.com",
				DomainExt: "com",
			},
			wantErr: false,
		},
		{
			name: "failed parse url",
			args: args{
				ctx:     context.TODO(),
				longUrl: "ini long url failed",
			},
			want:    GenerateResponse{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateShortURL(context.TODO(), tt.args.longUrl)

			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got.ShortURL) != 7 && !tt.wantErr {
				t.Errorf("GenerateShortURL() = %v", got)
				return
			}

			if got.Domain != tt.want.Domain || got.DomainExt != tt.want.DomainExt {
				t.Errorf("GenerateShortURL() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

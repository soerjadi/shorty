package model

import "time"

type URL struct {
	ID           int64     `db:"id" json:"id"`
	LongURL      string    `db:"long_url" json:"long_url"`
	ShortURL     string    `db:"short_url" json:"-"`
	Domain       string    `db:"domain" json:"-"`
	DomainExt    string    `db:"domain_ext" json:"-"`
	IsSSL        bool      `db:"is_ssl" json:"-"`
	IsAliased    bool      `db:"is_aliased" json:"-"`
	ClickCount   int       `db:"click_count" json:"click_count"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	FullShortURL string    `json:"url"`
}

type URLRequest struct {
	Query  *string `json:"query,omitempty"`
	SSL    *bool   `json:"is_ssl,omitempty"`
	Offset int32   `json:"offset"`
	Limit  int32   `json:"limit"`
}

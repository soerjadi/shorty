package util

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/soerjadi/short/internal/pkg/snowflake"
)

func GenerateShortURL(ctx context.Context, longURL string) (GenerateResponse, error) {
	url, err := url.ParseRequestURI(longURL)
	if err != nil {
		fmt.Errorf("pkg.Util.GenerateShortURL error parsing url. err = %v, url = %v", err, longURL)

		return GenerateResponse{}, err
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Errorf("pkg.Util.GenerateShortURL failed create new node. err = %v", err)

		return GenerateResponse{}, err
	}

	nodeSN := node.Generate()
	domain := strings.Split(url.Hostname(), ".")

	return GenerateResponse{
		ID:        nodeSN.Int64(),
		ShortURL:  Base62Conversion(uint64(nodeSN.Int64())),
		Domain:    url.Hostname(),
		DomainExt: domain[1],
	}, err
}

const chars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Base62Conversion(nb uint64) string {
	base := chars
	var buf bytes.Buffer
	l := uint64(len(base))

	ll := l * 6

	if nb/ll != 0 {
		encode(nb/ll, &buf, base)
	}
	buf.WriteByte(base[nb%l])
	return buf.String()
}

func encode(nb uint64, buf *bytes.Buffer, base string) {
	l := uint64(len(base))
	ll := l * 6

	if nb/ll != 0 {
		encode(nb/ll, buf, base)
	}
	buf.WriteByte(base[nb%l])
}

func GetENV() string {
	env := os.Getenv("ENV")
	if env == "" {
		return "DEVELOPMENT"
	}

	return env
}

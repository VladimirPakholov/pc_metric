package db

import (
	"net"
	"net/url"
	"os"
)

func BuildPostgreDSN() string {
	u := &url.URL{
		Scheme: "postgres",
		User: url.UserPassword(
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
		),
		Host: net.JoinHostPort(
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		),
		Path: os.Getenv("DB_NAME"),
	}
	q := u.Query()
	q.Set("sslmode", os.Getenv("DB_SSLMODE"))
	u.RawQuery = q.Encode()

	return u.String()
}

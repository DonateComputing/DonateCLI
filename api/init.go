package api

import (
	"net/http"
	"os"
)

var client = &http.Client{}

var domain = getDomain()

func getDomain() string {
	d := os.Getenv("DOMAIN")
	if d == "" {
		return "http://localhost:8080"
	}
	return d
}

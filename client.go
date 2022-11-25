package atcodergo

import (
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	*http.Client
	token string
}

func NewClient() (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		&http.Client{Jar: jar},
		"",
	}, nil
}

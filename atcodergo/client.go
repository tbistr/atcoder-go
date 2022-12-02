package atcodergo

import (
	"net/http"
	"net/http/cookiejar"
)

// Client for atcoder.
// available as http.Client
type Client struct {
	*http.Client
	token string
}

// NewClient create new Client.
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

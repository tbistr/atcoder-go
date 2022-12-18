package atcodergo

import (
	"encoding/json"
	"io"
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

func (c *Client) writeCookie(w io.Writer) error {
	b, err := json.Marshal(c.Jar.Cookies(&BASE_URL.URL))
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func (c *Client) readCookie(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	cookies := []*http.Cookie{}
	if err := json.Unmarshal(b, &cookies); err != nil {
		return err
	}
	c.Jar.SetCookies(&BASE_URL.URL, cookies)
	return nil
}

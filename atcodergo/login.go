package atcodergo

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Login to atcoder.
// save csrf_token to client.
func (c *Client) Login(username, password string) error {

	// get csrf_token
	tokenResp, err := c.Get(BASE_URL.login().String())
	if err != nil {
		return err
	}
	defer readAllClose(tokenResp.Body)
	doc, err := goquery.NewDocumentFromReader(tokenResp.Body)
	if err != nil {
		return err
	}
	token, exist := doc.Find("input[name=csrf_token]").Attr("value")
	if !exist {
		return errors.New("csrf_token not found")
	}

	// login request
	values := url.Values{}
	values.Set("username", username)
	values.Set("password", password)
	values.Set("csrf_token", token)
	tryResp, err := c.PostForm(BASE_URL.login().String(), values)
	if err != nil {
		return err
	}
	defer readAllClose(tokenResp.Body)

	b, err := io.ReadAll(tryResp.Body)
	if err != nil {
		return err
	}
	if strings.Contains(string(b), "ユーザ名またはパスワードが正しくありません。") {
		return errors.New("faild to login")
	} else {
		c.token = token
		return nil
	}
}

// Logout from atcoder.
// Removes session file if exists.
func (c *Client) Logout() error {
	values := url.Values{}
	values.Set("csrf_token", c.token)
	r, err := c.PostForm(BASE_URL.logout().String(), values)
	if err != nil {
		return err
	}
	defer readAllClose(r.Body)
	c.token = ""

	if c.sessionFile == "" {
		return nil
	} else {
		f, err := os.Create(c.sessionFile)
		if err != nil {
			return err
		}
		defer f.Close()
		return c.writeCookie(f)
	}
}

// LoginWithNewSession try to login to atcoder.
// File contents are overwritten.
func (c *Client) LoginWithNewSession(username, password, file string) error {
	if err := c.Login(username, password); err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	if err := c.writeCookie(f); err != nil {
		return err
	}

	c.sessionFile = file
	return f.Close()
}

// LoginWithSession try to login to atcoder.
// File contents are only read.
func (c *Client) LoginWithSession(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := c.readCookie(f); err != nil {
		return err
	}

	if !c.checkLoggedin() {
		return errors.New("faild to login")
	}

	c.sessionFile = file
	return nil
}

// checkLoggedin checks if logged in or not
// by GET "https://atcoder.jp/contests/practice/tasks".
// (404 if not authed.)
func (c *Client) checkLoggedin() bool {
	r, err := c.Get(BASE_URL.tasks("practice").String())
	if err != nil {
		return false
	}
	defer readAllClose(r.Body)
	return r.StatusCode != http.StatusNotFound
}

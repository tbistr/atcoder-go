package atcodergo

import (
	"errors"
	"fmt"
	"io"
	"net/url"
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
	defer tokenResp.Body.Close()
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
	defer tryResp.Body.Close()

	b, err := io.ReadAll(tryResp.Body)
	if err != nil {
		return err
	}
	if strings.Contains(string(b), "ユーザ名またはパスワードが正しくありません。") {
		return fmt.Errorf("faild to login")
	} else {
		c.token = token
		return nil
	}
}

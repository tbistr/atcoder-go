package atcodergo

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) Login(username, password string) error {

	// csrf_tokenの取得
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

	// ログインリクエスト
	values := url.Values{}
	values.Set("username", username)
	values.Set("password", password)
	values.Set("csrf_token", token)
	tryResp, err := c.PostForm(BASE_URL.login().String(), values)
	if err != nil {
		return err
	}
	defer tryResp.Body.Close()
	if tryResp.StatusCode != http.StatusOK {
		return fmt.Errorf("faild to login: %s", tryResp.Status)
	}
	return nil
}
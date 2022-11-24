package atcodergo

import "net/url"

var BASE_URL *base

type base struct{ url.URL }

func init() {
	u, err := url.Parse("https://atcoder.jp")
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Set("lang", "ja")
	u.RawQuery = q.Encode()
	BASE_URL = &base{*u}
}

// func (b base) base() *url.URL {
// 	return &b.URL
// }

func (b base) login() *url.URL {
	return b.URL.JoinPath("login")
}

func (b base) contests() *url.URL {
	return b.URL.JoinPath("contests")
}

func (b base) contestsArchive() *url.URL {
	return b.contests().JoinPath("archive")
}

func (b base) contestTop(id string) *url.URL {
	return b.contests().JoinPath(id)
}

func (b base) contestTasks(id string) *url.URL {
	return b.contestTop(id).JoinPath("tasks")
}

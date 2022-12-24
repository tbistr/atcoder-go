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

// login = ~/login
func (b base) login() *url.URL {
	return b.URL.JoinPath("login")
}

// logout = ~/logout
func (b base) logout() *url.URL {
	return b.URL.JoinPath("logout")
}

// contests = ~/contests
func (b base) contests() *url.URL {
	return b.URL.JoinPath("contests")
}

// contestsArchive = ~/contests/archive
func (b base) contestsArchive() *url.URL {
	return b.contests().JoinPath("archive")
}

// contest = ~/contests/{id}
func (b base) contest(id string) *url.URL {
	return b.contests().JoinPath(id)
}

// tasks = ~/contests/{id}
func (b base) tasks(contestID string) *url.URL {
	return b.contest(contestID).JoinPath("tasks")
}

// task = ~/contests/{contestID}/{taskID}
func (b base) task(contestID, taskID string) *url.URL {
	return b.tasks(contestID).JoinPath(taskID)
}

// submit = ~/contests/{contestID}/submit
func (b base) submit(contestID string) *url.URL {
	return b.contest(contestID).JoinPath("submit")
}

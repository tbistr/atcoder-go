package atcodergo

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/tbistr/atcoder-go/atcodergo/model"
	"github.com/tbistr/atcoder-go/atcodergo/parse"
	"github.com/tbistr/pig"
)

type Submission = model.Submission

// Submit answer program for the task.
// Returns submission result (first of submission list).
func (c *Client) Submit(contestID, taskID string, languageID string, program io.Reader) (*Submission, error) {
	if !c.loggedin {
		return nil, newNeedAuthError("Submit()")
	}

	u := BASE_URL.submit(contestID)
	v := url.Values{}
	v.Set("data.TaskScreenName", taskID)
	v.Set("data.LanguageId", languageID)
	b, err := io.ReadAll(program)
	if err != nil {
		return nil, err
	}
	v.Set("sourceCode", string(b))
	v.Set("csrf_token", c.token)
	resp, err := c.PostForm(u.String(), v)
	if err != nil {
		return nil, err
	}
	defer readAllClose(resp.Body)

	if compareURL(*BASE_URL.submissions(contestID), *resp.Request.URL) {
		// TODO: resp.BodyのHTMLをチェックしてエラー原因を出す
		return nil, errors.New("failed to submit program")
	}

	doc, err := pig.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	submissions := parse.Submissions(doc)
	if len(submissions) == 0 {
		return nil, errors.New("failed to get submission")
	}

	return submissions[0], nil
}

// SubmissionsPager is pager for submissions.
// Atcoder serves submissions list with pagination.
type SubmissionsPager struct {
	client *Client
	url    *url.URL
	page   int
}

// NewSubmissionsPager creates new SubmissionsPager.
func (c *Client) NewSubmissionsPager(contestID string) *SubmissionsPager {
	return &SubmissionsPager{
		client: c,
		url:    BASE_URL.submissions(contestID),
		page:   0,
	}
}

// parseSubmissions parses submission page.
// https://atcoder.jp/contests/practice/submissions/me
func (pager *SubmissionsPager) Next() (submissions []*Submission, ok bool) {
	pager.page++

	q := url.Values{}
	q.Set("page", strconv.Itoa(pager.page))
	pager.url.RawQuery = q.Encode()

	resp, err := pager.client.Get(pager.url.String())
	if err != nil {
		return nil, false
	}
	defer readAllClose(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, false
	}

	doc, err := pig.Parse(resp.Body)
	if err != nil {
		return nil, false
	}

	return parse.Submissions(doc), len(submissions) != 0
}

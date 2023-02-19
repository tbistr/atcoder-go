package atcodergo

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

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

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	submissions := parseSubmissions(doc)
	if len(submissions) == 0 {
		return nil, errors.New("failed to get submission")
	}

	return submissions[0], nil
}

type Submission struct {
	ID           string
	Time         time.Time
	TaskName     string
	UserID       string
	LanguageName string
	Score        string
	CodeSize     string
	Status       string
	ExecTime     string
	Memory       string
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

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, false
	}

	return parseSubmissions(doc), len(submissions) != 0
}

func parseSubmissions(doc *goquery.Document) []*Submission {
	submissions := []*Submission{}
	doc.Find("table > tbody > tr").Each(func(i int, tr *goquery.Selection) {
		td := tr.Find("td")

		inWJ := len(td.Nodes) == 8
		if len(td.Nodes) != 10 && !inWJ {
			return
		}

		// parse each td
		// doesn't catch any errors
		// (will be empty)
		submission := &Submission{}
		td.Each(func(i int, s *goquery.Selection) {
			switch {
			case i == 0: // Submission Time
				var err error
				submission.Time, err = time.Parse("2006-01-02 15:04:05-0700", s.Text())
				if err != nil {
					submission.Time, _ = time.Parse("2006-01-02 15:04:05", s.Text())
				}
			case i == 1: // Task
				// assumes href="/contests/practice/tasks/practice_1"
				if href, exists := s.Children().First().Attr("href"); exists {
					ss := strings.Split(href, "/")
					if len(ss) == 5 {
						// contestID = ss[2]
						submission.TaskName = ss[4]
					}
				}
			case i == 2: // User
				// assumes href="/users/tbistr"
				if href, exists := s.Children().First().Attr("href"); exists {
					ss := strings.Split(href, "/")
					if len(ss) == 3 {
						submission.UserID = ss[2]
					}
				}
			case i == 3:
				submission.LanguageName = s.Text()
			case i == 4:
				submission.Score = s.Text()
			case i == 5:
				submission.CodeSize = s.Text()
			case i == 6:
				submission.Status = s.Text()

			// If status == WJ, ExecTime and Memory is empty.
			case !inWJ && i == 7:
				submission.ExecTime = s.Text()
			case !inWJ && i == 8:
				submission.Memory = s.Text()

			case !inWJ && i == 9: // Details
			case inWJ && i == 7:
				// assumes href="/contests/practice/submissions/00000000"
				if href, exists := s.Children().First().Attr("href"); exists {
					ss := strings.Split(href, "/")
					if len(ss) == 5 {
						submission.ID = ss[4]
					}
				}
			}
		})
		submissions = append(submissions, submission)
	})
	return submissions
}

package atcodergo

import (
	"errors"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

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

// Submit answer program for the task.
func (c *Client) Submit(contestID, taskID string, languageID string, program io.Reader) error {
	if !c.loggedin {
		return newNeedAuthError("Submit()")
	}

	u := BASE_URL.submit(contestID)
	v := url.Values{}
	v.Set("data.TaskScreenName", taskID)
	v.Set("data.LanguageId", languageID)
	b, err := io.ReadAll(program)
	if err != nil {
		return err
	}
	v.Set("sourceCode", string(b))
	v.Set("csrf_token", c.token)
	resp, err := c.PostForm(u.String(), v)
	if err != nil {
		return err
	}
	defer readAllClose(resp.Body)

	if compareURL(*BASE_URL.submissions(contestID), *resp.Request.URL) {
		// TODO: resp.BodyのHTMLをチェックしてエラー原因を出す
		return errors.New("failed to submit program")
	}
	// TODO: tableの一番上を読んで自分の提出情報を返す

	return nil
}

func (c *Client) Submission() error {
	resp, _ := c.Get("https://atcoder.jp/contests/practice/submissions/me")

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
	return nil
	}
	parseSubmissions(doc)

	return nil
}

func parseSubmissions(doc *goquery.Document) ([]*Submission, error) {
	submissions := []*Submission{}
	var err error

	doc.Find("table > tbody >tr").Each(func(i int, tr *goquery.Selection) {
		td := tr.Find("td")
		if len(td.Nodes) != 10 || err != nil {
			err = errors.New("failed to parse submissions")
			return
		}

		// parse each td
		// don't catch any errors
		// (will be empty)
		submission := &Submission{}
		td.Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0: // Submission Time
				var err error
				submission.Time, err = time.Parse("2006-01-02 15:04:05-0700", s.Text())
				if err != nil {
					submission.Time, _ = time.Parse("2006-01-02 15:04:05", s.Text())
				}
			case 1: // Task
				// assumes href="/contests/practice/tasks/practice_1"
				if href, exists := s.Children().First().Attr("href"); exists {
					ss := strings.Split(href, "/")
					if len(ss) == 5 {
						// contestID = ss[2]
						submission.TaskName = ss[4]
					}
				}
			case 2: // User
				// assumes href="/users/tbistr"
				if href, exists := s.Children().First().Attr("href"); exists {
					ss := strings.Split(href, "/")
					if len(ss) == 3 {
						submission.UserID = ss[2]
					}
				}
			case 3:
				submission.LanguageName = s.Text()
			case 4:
				submission.Score = s.Text()
			case 5:
				submission.CodeSize = s.Text()
			case 6:
				submission.Status = s.Text()
			case 7:
				submission.ExecTime = s.Text()
			case 8:
				submission.Memory = s.Text()
			case 9: // Details
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
	return submissions, nil
}

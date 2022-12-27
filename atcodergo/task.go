package atcodergo

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Task
type Task struct {
	Name   string
	IdName string // "A", "90"...
	ID     string // "abc123_a", "typical90_cl"...
}

// Tasks gets tasks of the contest.
func (c *Client) Tasks(contestID string) ([]*Task, error) {
	if !c.loggedin {
		return nil, newNeedAuthError("Tasks()")
	}

	u := BASE_URL.tasks(contestID)
	resp, err := c.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer readAllClose(resp.Body)

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("no such contest: %s", contestID)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cant access tasks, status not ok: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	tasks := make([]*Task, 0, 10)
	doc.Find("table > tbody >tr").Each(func(i int, s *goquery.Selection) {
		as := s.Find("td > a")
		col1 := as.Eq(0)

		href, _ := col1.Attr("href")
		id := strings.Trim(href, "/")
		id = id[strings.LastIndex(id, "/")+1:]
		idName := col1.Text()

		col2 := as.Eq(1)
		name := col2.Text()

		tasks = append(tasks, &Task{
			Name:   name,
			IdName: idName,
			ID:     id,
		})
	})

	return tasks, nil
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

	return nil
}

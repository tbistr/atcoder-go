package atcodergo

import (
	"fmt"
	"io"
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

// TestCase for Task
type TestCase struct{ Input, Output string }

// Tasks gets tasks of the contest.
func (c *Client) Tasks(contestID string) ([]*Task, error) {
	u := BASE_URL.tasks(contestID)
	resp, err := c.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	tasks := make([]*Task, 0, 10)
	doc.Find("table > tbody >tr").Each(func(i int, s *goquery.Selection) {
		as := s.Find("td > a")
		td1 := as.First()

		href, _ := td1.Attr("href")
		id := strings.Trim(href, "/")
		id = id[strings.LastIndex(id, "/")+1:]
		idName := td1.Text()

		td2 := as.Last()
		name := td2.Text()

		tasks = append(tasks, &Task{
			Name:   name,
			IdName: idName,
			ID:     id,
		})
	})

	return tasks, nil
}

// TestCases gets testcases of the task.
func (c *Client) TestCases(contestID, taskID string) ([]*TestCase, error) {
	u := BASE_URL.task(contestID, taskID)
	resp, err := c.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	tcs := make([]*TestCase, 0, 5)
	doc.Find("div.part > section").Each(func(i int, s *goquery.Selection) {
		if strings.HasPrefix(strings.TrimSpace(s.Find("h3").Text()), "入力例") {
			tcs = append(tcs, &TestCase{Input: s.Find("pre").Text()})
		} else if strings.HasPrefix(strings.TrimSpace(s.Find("h3").Text()), "出力例") {
			tcs[len(tcs)-1].Output = s.Find("pre").Text()
		}
	})

	return tcs, nil
}

// Submit answer program for the task.
func (c *Client) Submit(contest *Contest, task *Task, program io.Reader, languageID string) error {
	u := BASE_URL.submit(contest.ID)
	v := url.Values{}
	v.Set("data.TaskScreenName", task.ID)
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
	b, _ = io.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(string(b))

	return nil
}

func (c *Client) Languages() (map[string]string, error) {
	return map[string]string{"4006": "Python (3.4.3)"}, nil
}

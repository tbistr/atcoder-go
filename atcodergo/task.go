package atcodergo

import (
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
	if !c.loggedin {
		return nil, newNeedAuthError("Tasks()")
	}

	u := BASE_URL.tasks(contestID)
	resp, err := c.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer readAllClose(resp.Body)

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

// TestCases gets testcases of the task.
func (c *Client) TestCases(contestID, taskID string) ([]*TestCase, error) {
	if !c.loggedin {
		return nil, newNeedAuthError("TestCases()")
	}

	u := BASE_URL.task(contestID, taskID)
	resp, err := c.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer readAllClose(resp.Body)
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

// Language stand for
// "<option value={Language.Value} data-mime={Language.Datamime}>{Language.Text}</option>"
type Language struct {
	Value    string
	Datamime string
	Text     string
}

// Languages lists up acceptable languages.
func (c *Client) Languages() ([]Language, error) {
	if !c.loggedin {
		return nil, newNeedAuthError("Languages()")
	}

	resp, err := c.Get(BASE_URL.submit("practice").String())
	if err != nil {
		return nil, err
	}
	defer readAllClose(resp.Body)

	Languages := []Language{}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	doc.Find("#select-lang-practice_1 > select > option").Each(func(i int, s *goquery.Selection) {
		// index-0 == "<option></option>"
		if i == 0 {
			return
		}
		// else is "<option value="4001" data-mime="text/x-csrc">C (GCC 9.2.1)</option>"
		value, _ := s.Attr("value")
		datamime, _ := s.Attr("data-mime")
		text := s.Text()
		Languages = append(Languages, Language{Value: value, Datamime: datamime, Text: text})
	})

	return Languages, nil
}

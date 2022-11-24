package atcodergo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Task struct {
	Name   string
	IdName string // "A", "90"...
	ID     string // "abc123_a", "typical90_cl"...
}

func (c *Client) Tasks(contestID string) ([]*Task, error) {
	u := BASE_URL.contestTasks(contestID)
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

package parse

import (
	"strings"
	"time"

	"github.com/tbistr/atcoder-go/atcodergo/model"
	"github.com/tbistr/pig"
)

func Task(doc pig.Node) []*model.Task {
	tasks := []*model.Task{}
	doc.FindDescendant(pig.Tag("table"), pig.Tag("tbody"), pig.Tag("tr")).
		Each(func(i int, tr pig.Node) {
			as := tr.FindDescendant(pig.Tag("td"), pig.Tag("a"))
			col1, _ := as.Index(0)

			href, _ := col1.AttrVal("href")
			id := strings.TrimPrefix(href, "/")
			id = id[strings.LastIndex(id, "/")+1:]
			idName := col1.Text()

			col2, _ := as.Index(1)
			name := col2.Text()

			tasks = append(tasks, &model.Task{
				Name:   name,
				IdName: idName,
				ID:     id,
			})
		})

	return tasks
}

func TaskInfo(doc pig.Node) *model.TaskInfo {
	info := &model.TaskInfo{}
	doc.FindDescendant(
		pig.And(pig.Tag("div"), pig.ID("task-statement")),
		pig.And(pig.Tag("span"), pig.Cls("lang-ja")),
		pig.Tag("section"),
	).Each(func(i int, section pig.Node) {

		h3 := section.Find(pig.Tag("h3"))
		text := strings.TrimSpace(section.Text())
		// html, _ := section.Html()
		// html = strings.TrimSpace(html)
		pre := strings.TrimSpace(section.Find(pig.Tag("pre")).Text())
		p := strings.TrimSpace(section.Find(pig.Tag("p")).Text())

		switch {
		case strings.Contains(h3.Text(), "問題文"):
			info.ProblemStatement = text
			info.ProblemStatementHTML = ""
			// info.ProblemStatementHTML = html
		case strings.Contains(h3.Text(), "制約"):
			info.Constraints = text
			info.ConstraintsHTML = ""
			// info.ConstraintsHTML = html

		// TODO: consider if len(inputs) != len(outputs).
		case strings.Contains(h3.Text(), "入力例"):
			info.TestCases = append(info.TestCases,
				&model.TestCase{Input: pre},
			)
		case strings.Contains(h3.Text(), "出力例"):
			info.TestCases[len(info.TestCases)-1].Output = pre

		case strings.Contains(h3.Text(), "入力"):
			info.IoStyle.InputSig = pre
			info.IoStyle.InputDesc = p
		case strings.Contains(h3.Text(), "出力"):
			info.IoStyle.OutputSig = pre
			info.IoStyle.OutputDesc = p
		}
	})

	return info
}

func Submissions(doc pig.Node) []*model.Submission {
	submissions := []*model.Submission{}
	doc.FindDescendant(pig.Tag("table"), pig.Tag("tbody"), pig.Tag("tr")).
		Each(func(i int, tr pig.Node) {
			td := tr.Find(pig.Tag("td"))

			inWJ := len(td.Children()) == 8
			if len(td.Children()) != 10 && !inWJ {
				return
			}

			s := &model.Submission{}
			td.Each(func(i int, td pig.Node) {
				switch {
				case i == 0: // Submission Time
					var err error
					s.Time, err = time.Parse("2006-01-02 15:04:05-0700", td.Text())
					if err != nil {
						s.Time, _ = time.Parse("2006-01-02 15:04:05", td.Text())
					}
				case i == 1: // Task
					// assumes href="/contests/practice/tasks/practice_1"
					a, _ := td.Find(pig.Tag("a")).Index(0)
					if href, exists := a.AttrVal("href"); exists {
						ss := strings.Split(href, "/")
						if len(ss) == 5 {
							// contestID = ss[2]
							s.TaskName = ss[4]
						}
					}
				case i == 2: // User
					// assumes href="/users/tbistr"
					a, _ := td.Find(pig.Tag("a")).Index(0)
					if href, exists := a.AttrVal("href"); exists {
						ss := strings.Split(href, "/")
						if len(ss) == 3 {
							s.UserID = ss[2]
						}
					}
				case i == 3:
					s.LanguageName = td.Text()
				case i == 4:
					s.Score = td.Text()
				case i == 5:
					s.CodeSize = td.Text()
				case i == 6:
					s.Status = td.Text()

				// If status == WJ, ExecTime and Memory is empty.
				case !inWJ && i == 7:
					s.ExecTime = td.Text()
				case !inWJ && i == 8:
					s.Memory = td.Text()

				case !inWJ && i == 9: // Details
				case inWJ && i == 7:
					// assumes href="/contests/practice/submissions/00000000"
					a := td.Find(pig.Tag("a"))
					if href, exists := a.AttrVal("href"); exists {
						ss := strings.Split(href, "/")
						if len(ss) == 5 {
							s.ID = ss[4]
						}
					}
				}
			})
		})

	return submissions
}

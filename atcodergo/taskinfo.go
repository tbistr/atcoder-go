package atcodergo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TaskInfo is detailed information of a task.
// It is target to express [commonest task page] neither too much nor little.
type TaskInfo struct {
	ProblemStatement     string
	ProblemStatementHTML string
	Constraints          string
	ConstraintsHTML      string
	IoStyle              IoStyle
	TestCases            []*TestCase
}

// IoStyle represents input and output signeture.
// Input and Output are machine readable sections.
// ~Desc is Description of ones.
type IoStyle struct {
	InputSig, OutputSig   string
	InputDesc, OutputDesc string
}

// TestCase.
// Can be used to ascertain (.Input > `user program` == .Output).
type TestCase struct{ Input, Output string }

func (c *Client) TaskInfo(contestID, taskID string) (*TaskInfo, error) {
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

	ti := &TaskInfo{}
	// tcsIndex := -1
	doc.Find("div#task-statement span.lang-ja section").
		Each(func(i int, section *goquery.Selection) {
			h3 := section.Find("h3").Remove()
			text := strings.TrimSpace(section.Text())
			html, _ := section.Html()
			html = strings.TrimSpace(html)

			switch {
			case strings.Contains(h3.Text(), "問題文"):
				ti.ProblemStatement = text
				ti.ProblemStatementHTML = html
			case strings.Contains(h3.Text(), "制約"):
				ti.Constraints = text
				ti.ConstraintsHTML = html

			// TODO: consider if len(inputs) != len(outputs).
			case strings.Contains(h3.Text(), "入力例"):
				ti.TestCases = append(ti.TestCases, &TestCase{Input: section.Find("pre").Text()})
			case strings.Contains(h3.Text(), "出力例"):
				ti.TestCases[len(ti.TestCases)-1].Output = section.Find("pre").Text()

			case strings.Contains(h3.Text(), "入力"):
				ti.IoStyle.InputSig = strings.TrimSpace(section.Find("pre").Remove().Text())
				ti.IoStyle.InputDesc = strings.TrimSpace(section.Text())
			case strings.Contains(h3.Text(), "出力"):
				ti.IoStyle.OutputSig = strings.TrimSpace(section.Find("pre").Remove().Text())
				ti.IoStyle.OutputDesc = strings.TrimSpace(section.Text())
			}
		})

	return ti, nil
}

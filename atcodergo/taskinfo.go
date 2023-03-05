package atcodergo

import (
	"github.com/tbistr/atcoder-go/atcodergo/model"
	"github.com/tbistr/atcoder-go/atcodergo/parse"
	"github.com/tbistr/pig"
)

// TaskInfo is detailed information of a task.
// It is target to express [commonest task page] neither too much nor little.
type TaskInfo = model.TaskInfo

// IoStyle represents input and output signature.
// Input and Output are machine readable sections.
// ~Desc is Description of ones.
type IoStyle = model.IoStyle

// TestCase.
// Can be used to ascertain (.Input > `user program` == .Output).
type TestCase = model.TestCase

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
	doc, err := pig.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return parse.TaskInfo(doc), nil
}

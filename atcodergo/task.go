package atcodergo

import (
	"fmt"
	"net/http"

	"github.com/tbistr/atcoder-go/atcodergo/model"
	"github.com/tbistr/atcoder-go/atcodergo/parse"
	"github.com/tbistr/pig"
)

// Task
type Task = model.Task

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

	doc, err := pig.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return parse.Task(doc), nil
}

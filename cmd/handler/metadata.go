package handler

import "github.com/tbistr/atcoder-go/atcodergo"

const (
	CONTEST_META_FILE = "contest.json"
	TASK_META_FILE    = "task.json"
)

type ContestMeta struct {
	ContestID string            `json:"contest_id"`
	Tasks     []*atcodergo.Task `json:"tasks"`
}

type TaskMeta struct {
	ContestID string          `json:"contest_id"`
	Task      *atcodergo.Task `json:"task"`
	MainFile  string          `json:"main_file"`
}

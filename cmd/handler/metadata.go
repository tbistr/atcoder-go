package handler

const (
	CONTEST_META_FILE = "contest.json"
	TASK_META_FILE    = "task.json"
)

type ContestMeta struct {
	ContestID string `json:"contest_id"`
}

type TaskMeta struct {
	ContestID string `json:"contest_id"`
	TaskID    string `json:"task_id"`
}

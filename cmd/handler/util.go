package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/tbistr/atcoder-go/atcodergo"
)

// prompt is default promptui.Prompt{}.
func prompt(label string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:  label,
		Stdout: os.Stderr,
	}
}

// chooseTask chooses task which user select.
func chooseTask() (taskMeta *TaskMeta, taskDir string, err error) {
	inContestDir, inTaskDir := false, false

	fs, err := os.ReadDir(".")
	if err != nil {
		return nil, "", err
	}
	for _, f := range fs {
		inTaskDir = f.Name() == TASK_META_FILE
		inContestDir = f.Name() == CONTEST_META_FILE
		if inTaskDir || inContestDir {
			break
		}
	}

	if inTaskDir {
		taskDir = "."
	} else if inContestDir {
		// parse contest.json
		b, err := os.ReadFile(CONTEST_META_FILE)
		if err != nil {
			return nil, "", err
		}
		contestMeta := &ContestMeta{}
		if err := json.Unmarshal(b, contestMeta); err != nil {
			return nil, "", err
		}

		// user prompt
		task, err := selectTaskPrompt(contestMeta.Tasks)
		if err != nil {
			return nil, "", err
		}
		taskDir = filepath.Join(".", task.IdName)
	} else {
		return nil, "", fmt.Errorf("metadata (%s or %s) not found", CONTEST_META_FILE, TASK_META_FILE)
	}

	b, err := os.ReadFile(filepath.Join(taskDir, TASK_META_FILE))
	if err != nil {
		return nil, "", err
	}
	taskMeta = &TaskMeta{}
	return taskMeta, taskDir, json.Unmarshal(b, taskMeta)
}

// selectTaskPrompt shows prompt and hears which contest does user select.
func selectTaskPrompt(tasks []*atcodergo.Task) (*atcodergo.Task, error) {
	items := make([]string, 0, len(tasks))
	for _, task := range tasks {
		items = append(items, fmt.Sprintf("%s: %s", task.IdName, task.Name))
	}
	i, _, err := (&promptui.Select{
		Label: "Select Task",
		Items: items,
	}).Run()
	return tasks[i], err
}

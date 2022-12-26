package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tbistr/atcoder-go/atcodergo"
)

func (h *Handler) NewContest(contestID, templateFile string) error {
	tasks, err := h.atcoder.Tasks(contestID)
	if err != nil {
		return err
	}

	_, err = os.Stat(contestID)
	if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("./%s is already exists", contestID)
	}

	// make contest.json
	cMeta := ContestMeta{ContestID: contestID}
	b, err := json.Marshal(cMeta)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(contestID, 0755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(contestID, "contest.json"), b, 0644); err != nil {
		return err
	}

	for _, task := range tasks {
		if err := h.mkTaskDir(contestID, task); err != nil {
			return err
		}
	}

	return nil
}

// mkTaskDir makes
// - contestDir/{taskName}
// - contestDir/{taskName}/main.go
// - contestDir/{taskName}/testcase{1..n}.input
// - contestDir/{taskName}/testcase{1..n}.output
func (h *Handler) mkTaskDir(contestID string, task *atcodergo.Task) error {
	taskDir := filepath.Join(contestID, task.IdName)
	if err := os.Mkdir(taskDir, 0755); err != nil {
		return err
	}

	template, err := os.ReadFile(h.config.TemplateFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read template: %s: %s \n", h.config.TemplateFile, err)
	}
	mainFile := h.config.MainFileName
	if mainFile == "" {
		mainFile = "main.go"
	}
	if err := os.WriteFile(filepath.Join(taskDir, mainFile), template, 0644); err != nil {
		return err
	}

	tcs, err := h.atcoder.TestCases(contestID, task.ID)
	if err != nil {
		// テストケースのパースがダメだった可能性
		// パースに自信がないので、正常終了
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	for i, tc := range tcs {
		if err := os.WriteFile(
			filepath.Join(taskDir, fmt.Sprintf("testcase%d.input", i+1)),
			[]byte(tc.Input), 0644,
		); err != nil {
			return err
		}
		if err := os.WriteFile(
			filepath.Join(taskDir, fmt.Sprintf("testcase%d.output", i+1)),
			[]byte(tc.Output), 0644,
		); err != nil {
			return err
		}
	}

	return nil
}

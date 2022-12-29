package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

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
	if err := os.WriteFile(filepath.Join(contestID, CONTEST_META_FILE), b, 0644); err != nil {
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

	ti, err := h.atcoder.TaskInfo(contestID, task.ID)
	if err != nil {
		// テストケースのパースがダメだった可能性
		// パースに自信がないので、正常終了
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	if err := h.mkTemplateFile(taskDir, ti); err != nil {
		return err
	}

	taskMeta := &TaskMeta{
		ContestID: contestID,
		TaskID:    task.ID,
	}
	b, err := json.Marshal(taskMeta)
	if err != nil {
		return err
	}
	if err := os.WriteFile(
		filepath.Join(taskDir, TASK_META_FILE),
		b, 0644,
	); err != nil {
		return err
	}

	for i, tc := range ti.TestCases {
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

// mkTemplateFile makes template file for each task.
// [input signature or task info(json)] > [template cmd] > [template file]
func (h *Handler) mkTemplateFile(taskDir string, info *atcodergo.TaskInfo) error {
	mainFile := h.config.MainFileName
	if mainFile == "" {
		mainFile = "main.txt"
	}

	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmdExe := exec.CommandContext(timeout, h.config.TemplateCmdName, h.config.TemplateCmdArgs...)

	if h.config.TemplateCmdJsonInput {
		b, err := json.Marshal(info)
		if err != nil {
			return err
		}
		cmdExe.Stdin = strings.NewReader(string(b))
	} else {
		cmdExe.Stdin = strings.NewReader(info.IoStyle.InputSig)
	}

	template, err := cmdExe.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed or timeout(10s) to run template generator: %s\nusing template file: %s\n", err, h.config.TemplateFile)
		template, err = os.ReadFile(h.config.TemplateFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read template file: %s \n", err)
		}
	}

	return os.WriteFile(filepath.Join(taskDir, mainFile), template, 0644)
}

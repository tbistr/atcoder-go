package handler

import (
	"fmt"
	"os"
	"path/filepath"
)

func (h *Handler) Submit() error {
	taskMeta, taskDir, err := chooseTask()
	if err != nil {
		return err
	}

	f, err := os.Open(filepath.Join(taskDir, taskMeta.MainFile))
	if err != nil {
		return err
	}
	defer f.Close()

	submission, err := h.atcoder.Submit(taskMeta.ContestID, taskMeta.Task.ID, h.config.DefaultLanguage.Value, f)
	if err != nil {
		return err
	}

	fmt.Println("submit success.")
	fmt.Printf("see: https://atcoder.jp/contests/%s/submissions/%s\n", taskMeta.ContestID, submission.ID)

	return nil
}

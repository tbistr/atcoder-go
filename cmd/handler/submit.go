package handler

import (
	"encoding/json"
	"errors"
	"os"
)

func (h *Handler) Submit() error {
	f, err := os.Open(h.config.MainFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	taskMeta := &TaskMeta{}
	dirs, err := os.ReadDir(wd)
	if err != nil {
		return err
	}
	for _, d := range dirs {
		if d.IsDir() {
			continue
		}
		if d.Name() == TASK_META_FILE {
			b, err := os.ReadFile(d.Name())
			if err != nil {
				return err
			}
			if err := json.Unmarshal(b, taskMeta); err != nil {
				return err
			}
			break
		}
	}

	if taskMeta.ContestID == "" {
		return errors.New("current dir is not task dir")
	}

	return h.atcoder.Submit(taskMeta.ContestID, taskMeta.TaskID, h.config.DefaultLanguage.Value, f)
}

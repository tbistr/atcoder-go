package handler

import (
	"os"
	"path/filepath"
)

func (h *Handler) Login() error {
	// hear username, password
	u, err := prompt("enter username").Run()
	if err != nil {
		return err
	}

	prompt := prompt("enter password")
	prompt.Mask = '*'
	p, err := prompt.Run()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(h.config.SessionFile), 0755); err != nil {
		return err
	}
	return h.atcoder.LoginWithNewSession(u, p, h.config.SessionFile)
}

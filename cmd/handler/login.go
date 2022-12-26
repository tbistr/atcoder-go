package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/term"
)

func (h *Handler) Login() error {
	// hear username, password
	var u, p string
	fmt.Fprint(os.Stderr, "enter username:")
	fmt.Scanln(&u)
	fmt.Fprint(os.Stderr, "enter password:")
	b, _ := term.ReadPassword(int(syscall.Stdin))
	p = string(b)
	fmt.Fprint(os.Stderr, "\n")

	if err := os.MkdirAll(filepath.Dir(h.config.SessionFile), 0755); err != nil {
		return err
	}
	return h.atcoder.LoginWithNewSession(u, p, h.config.SessionFile)
}

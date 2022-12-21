package handler

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func (h *Handler) Login(sessionFile string) {
	if h.atcoder.LoginWithSession(sessionFile) != nil {
		// hear username, password
		var u, p string
		fmt.Fprint(os.Stderr, "enter username:")
		fmt.Scanln(&u)
		fmt.Fprint(os.Stderr, "enter password:")
		b, _ := term.ReadPassword(int(syscall.Stdin))
		p = string(b)
		fmt.Fprint(os.Stderr, "\n")

		if err := h.atcoder.LoginWithNewSession(u, p, sessionFile); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

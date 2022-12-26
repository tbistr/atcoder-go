package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tbistr/atcoder-go/atcodergo"
)

// configDir gets config dir's and child item's path.
// ex. ~/.config/atgo
// ex. ~/.config/atgo/{elem...}
func configDir(elem ...string) string {
	home, err := os.UserHomeDir()
	exit1withE(err)
	return filepath.Join(append([]string{home, ".config", "atgo"}, elem...)...)
}

// exit1withE output e to stderr and exit(1).
// If e == nil, noop.
func exit1withE(e error) {
	needAuthErr := &atcodergo.NeedAuthError{}
	if errors.As(e, &needAuthErr) {
		fmt.Println("Please login. Try run `atgo login`.")
		os.Exit(0)
	}

	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

package handler

import (
	"os"

	"github.com/manifoldco/promptui"
)

// prompt is default promptui.Prompt{}.
func prompt(label string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:  label,
		Stdout: os.Stderr,
	}
}

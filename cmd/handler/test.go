package handler

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Test runs program with execution template.
// Show assertion results.
// Returns early if err (like compile err).
func (h *Handler) Test() error {
	_, taskDir, err := chooseTask()
	if err != nil {
		return err
	}

	files, err := os.ReadDir(taskDir)
	if err != nil {
		return err
	}
	for i := 0; i < len(files)/2; i++ {
		// NOTE: add testcase's filename list to task.json?
		in, err := os.ReadFile(filepath.Join(taskDir, fmt.Sprintf("testcase%d.input", i+1)))
		if err != nil {
			break
		}
		out, err := os.ReadFile(filepath.Join(taskDir, fmt.Sprintf("testcase%d.output", i+1)))
		if err != nil {
			break
		}

		// Exec program runner.
		for j, a := range h.config.RunCmdArgs {
			if a == "{{main}}" {
				h.config.RunCmdArgs[j] = filepath.Join(taskDir, h.config.MainFileName)
			}
		}
		c := exec.Command(h.config.RunCmdName, h.config.RunCmdArgs...)
		c.Stdin = bytes.NewReader(in)
		var stdout, stderr bytes.Buffer
		c.Stdout, c.Stderr = &stdout, &stderr
		c.Run()
		if c.ProcessState.ExitCode() != 0 {
			fmt.Println("Failed to exec!")
			fmt.Print(stderr.String())
			break
		}

		// Show assertion results.
		assumed := strings.TrimSpace(string(out))
		actual := strings.TrimSpace(stdout.String())
		if assumed == actual {
			fmt.Printf("test[%d] passed!\n", i+1)
		} else {
			fmt.Printf("test[%d] failed!\n", i+1)
			fmt.Println("===Assumed===")
			fmt.Println(assumed)
			fmt.Println("=============")
			fmt.Println("===Actual====")
			fmt.Println(actual)
			fmt.Println("=============")
			fmt.Println()
		}
	}
	return nil
}

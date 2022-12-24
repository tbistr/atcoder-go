package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/k0kubun/pp"
	"github.com/tbistr/atcoder-go/atcodergo"
	"golang.org/x/term"
)

func main() {
	c, _ := atcodergo.NewClient()

	// Some example has too much output.
	// Please comment out function `~Example(c)`, if you like.

	// login
	loginExample(c)

	// lists acceptable languages
	// l, _ := c.Languages()
	// pp.Println(l)

	// lists contests (page->1)
	// listContestsExample(c)

	// lists tasks of practice contest
	listTasksExample(c)

	// lists testcases of practice contest, task A
	listTestcasesExample(c)

	// submits program to practice contest, task A
	// ⚠️ THIS EXAMPLE WILL ACTUALY SUBMIT PROGRAM ⚠️
	// submitExample(c)

	// logout
	c.Logout()

	// error handling
	_, err := c.Tasks("hoge")
	var needAuth *atcodergo.NeedAuthError
	if errors.As(err, &needAuth) {
		fmt.Println(needAuth)
	}
}

func loginExample(c *atcodergo.Client) {
	session := "./.atcoder_session"
	err := c.LoginWithSession(session)
	if err != nil {
		fmt.Println(err)
		u := ""
		fmt.Fprint(os.Stderr, "username:")
		fmt.Scanln(&u)
		fmt.Fprint(os.Stderr, "password:")
		b, _ := term.ReadPassword(int(syscall.Stdin))
		fmt.Fprint(os.Stderr, "\n")
		if err := c.LoginWithNewSession(u, string(b), session); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func listContestsExample(c *atcodergo.Client) {
	pager := c.NewContestsPager()
	for contests, ok := pager.Next(); ok; contests, ok = pager.Next() {
		for _, contest := range contests {
			pp.Println(contest)
		}
	}
}

func listTasksExample(c *atcodergo.Client) {
	pager := c.NewContestsPager()
	contests, _ := pager.Next()
	practice := contests[0]
	pp.Println(practice)

	tasks, err := c.Tasks(practice.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pp.Println(tasks)
}

func listTestcasesExample(c *atcodergo.Client) {
	pager := c.NewContestsPager()
	contests, _ := pager.Next()
	practice := contests[0]
	pp.Println(practice)

	tasks, err := c.Tasks(practice.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	testcases, err := c.TestCases(practice.ID, tasks[0].ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pp.Println(testcases)
}

func submitExample(c *atcodergo.Client) {
	program := `package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var s string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)
	fmt.Printf("%d %s\n", a+b+c, s)
}`

	pager := c.NewContestsPager()
	contests, _ := pager.Next()
	practice := contests[0]
	pp.Println(practice)

	tasks, err := c.Tasks(practice.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := c.Submit(practice.ID, tasks[0].ID, "4026", strings.NewReader(program)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

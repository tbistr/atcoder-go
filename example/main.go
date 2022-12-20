package example

import (
	"fmt"
	"os"
	"syscall"

	"github.com/k0kubun/pp"
	"github.com/tbistr/atcoder-go/atcodergo"
	"golang.org/x/term"
)

func main() {
	c, _ := atcodergo.NewClient()

	session := "./.atcoder_session"
	if c.LoginWithSession(session) != nil {
		fmt.Fprint(os.Stderr, "password?:")
		b, _ := term.ReadPassword(int(syscall.Stdin))
		fmt.Fprint(os.Stderr, "\n")
		if err := c.LoginWithNewSession("your_username", string(b), session); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		b = nil
	}

	pager := c.NewContestsPager()
	contests, _ := pager.Next()
	prac := contests[0]

	ts, err := c.Tasks(prac.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pp.Println(ts)
}

package main

import (
	"fmt"

	atcodergo "github.com/tbistr/atcoder-go"
)

func main() {
	c, _ := atcodergo.NewClient()
	pager := c.NewContestsPager()

	for contests, ok := pager.Next(); ok; contests, ok = pager.Next() {
		for _, v := range contests {
			fmt.Printf("%s\n", v.Name)
		}
		fmt.Println(ok)
	}
}

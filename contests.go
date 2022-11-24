package atcodergo

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Contest struct {
	// id   string // like "abc123"
	// kind string
}

type ContestsPager struct {
	client *Client
	index  int
}

func (c *Client) NewContestsPager() *ContestsPager {
	return &ContestsPager{
		client: c,
		index:  0,
	}
}

func (pager *ContestsPager) Next() (contests []*Contest, ok bool) {
	u := BASE_URL.contests()

	if pager.index != 0 {
		u = BASE_URL.contestsArchive()
		q := url.Values{}
		q.Set("page", strconv.Itoa(pager.index))
		u.RawQuery = q.Encode()
	}

	resp, err := pager.client.Get(u.String())
	if err != nil {
		return nil, false
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, false
	}

	if pager.index == 0 {
		pager.nowContests(doc)
		pager.oldContests(doc)
	} else {
	}

	pager.index++
	return nil, true
}

func (pager *ContestsPager) nowContests(doc *goquery.Document) (contests []*Contest, err error) {
	selection := doc.Find("div#contest-table-permanent > div > div > table > tbody")
	h, _ := selection.Html()
	fmt.Println(h)

	return nil, nil
}

func (pager *ContestsPager) oldContests(doc *goquery.Document) (contests []*Contest, err error) {
	selection := doc.Find("div#contest-table-upcoming > div > div > table > tbody")
	h, _ := selection.Html()
	fmt.Println(h)

	return nil, nil
}

// func parseContests(c string) *Contest {
// 	return nil
// }

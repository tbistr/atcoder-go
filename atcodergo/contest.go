package atcodergo

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/tbistr/atcoder-go/atcodergo/model"
	"github.com/tbistr/atcoder-go/atcodergo/parse"
	"github.com/tbistr/pig"
)

// Contest
//
//	type Contest struct {
//	    Name  string
//	    ID    string // like "abc123"
//	    Kind  string // like "Algorithm", "Heuristics"...
//	    State string // "permanent", "upcoming", "archive"
//	}
type Contest = model.Contest

// ContestsPager is pager for contests.
// Atcoder's website serves contests list with pagination.
type ContestsPager struct {
	client *Client
	page   int
}

// NewContestsPager creates new ContestsPager.
func (c *Client) NewContestsPager() *ContestsPager {
	return &ContestsPager{
		client: c,
		page:   0,
	}
}

// Next returns next page's contests.
func (pager *ContestsPager) Next() (contests []*Contest, ok bool) {
	u := BASE_URL.contests()

	if pager.page != 0 {
		u = BASE_URL.contestsArchive()
		q := url.Values{}
		q.Set("page", strconv.Itoa(pager.page))
		u.RawQuery = q.Encode()
	}

	resp, err := pager.client.Get(u.String())
	if err != nil {
		return nil, false
	}
	defer readAllClose(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, false
	}

	doc, err := pig.Parse(resp.Body)
	if err != nil {
		return nil, false
	}

	var cs []*Contest
	if pager.page == 0 {
		cs = append(parse.Permanents(doc), parse.Upcomings(doc)...)
	} else {
		cs = parse.Archives(doc)
	}

	pager.page++
	return cs, len(cs) != 0
}

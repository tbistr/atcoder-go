package atcodergo

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Contest struct {
	Name  string
	href  string // like "/contests/abc123"
	Kind  string // like "Algorithm", "Heuristics"...
	State string // "permanent", "upcoming",
	// TODO: StateをEnumに
}

type ContestsPager struct {
	client *Client
	page   int
}

func (c *Client) NewContestsPager() *ContestsPager {
	return &ContestsPager{
		client: c,
		page:   0,
	}
}

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
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, false
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, false
	}

	cs := make([]*Contest, 0, 20)
	if pager.page == 0 {
		perms, err := permanents(doc)
		if err != nil {
			return nil, false
		}
		ups, err := upcomings(doc)
		if err != nil {
			return nil, false
		}
		cs = append(cs, perms...)
		cs = append(cs, ups...)
	} else {
		arc, err := archives(doc)
		if err != nil {
			return nil, false
		}
		cs = append(cs, arc...)
	}

	pager.page++
	return cs, true
}

// permanents 常時開催のコンテスト情報テーブルを辿ってパース
func permanents(doc *goquery.Document) (contests []*Contest, err error) {
	contests = make([]*Contest, 0, 10)
	doc.Find("div#contest-table-permanent > div > div > table > tbody > tr").
		Each(func(i int, s *goquery.Selection) {
			// 各行の内容をパース
			a := s.Find("td > a")
			name := a.Text()
			href, _ := a.Attr("href")
			span := s.Find("td > span[title]")
			kind, _ := span.Attr("title")
			contests = append(contests, &Contest{
				Name:  name,
				href:  href,
				Kind:  kind,
				State: "permanent",
			})
		})
	return contests, nil
}

// upcomings 開催予定のコンテスト情報テーブルを辿ってパース
func upcomings(doc *goquery.Document) (contests []*Contest, err error) {
	contests = make([]*Contest, 0, 10)
	doc.Find("div#contest-table-upcoming > div > div > table > tbody > tr").
		Each(func(i int, s *goquery.Selection) {
			// 各行の内容をパース
			a := s.Find("td > a").Last()
			name := a.Text()
			href, _ := a.Attr("href")
			span := s.Find("td > span[title]")
			kind, _ := span.Attr("title")
			contests = append(contests, &Contest{
				Name:  name,
				href:  href,
				Kind:  kind,
				State: "upcoming",
			})
		})
	return contests, nil
}

// upcomings 開催済みのコンテスト情報テーブルを辿ってパース
func archives(doc *goquery.Document) (contests []*Contest, err error) {
	contests = make([]*Contest, 0, 10)
	doc.Find("table > tbody > tr").
		Each(func(i int, s *goquery.Selection) {
			// 各行の内容をパース
			a := s.Find("td > a").Last()
			name := a.Text()
			href, _ := a.Attr("href")
			span := s.Find("td > span[title]")
			kind, _ := span.Attr("title")
			contests = append(contests, &Contest{
				Name:  name,
				href:  href,
				Kind:  kind,
				State: "archive",
			})
		})
	return contests, nil
}

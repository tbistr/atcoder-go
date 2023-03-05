package parse

import (
	"strings"

	"github.com/tbistr/atcoder-go/atcodergo/model"
	"github.com/tbistr/pig"
)

func parseContest(doc pig.Node, tableMatch pig.Match, tdIndex int, state string) []*model.Contest {
	return pig.Map(
		doc.FindDescendant(tableMatch, pig.Tag("tbody"), pig.Tag("tr")),
		func(tr pig.Node) *model.Contest {
			td, _ := tr.Find(pig.Tag("td")).Index(tdIndex)
			a, _ := td.Find(pig.Tag("a")).Index(0)
			name := a.Text()
			href, _ := a.AttrVal("href")
			span, _ := td.Find(pig.Tag("span")).Index(0)
			kind, _ := span.AttrVal("title")
			return &model.Contest{
				Name:  name,
				ID:    strings.TrimPrefix(href, "/contests/"),
				Kind:  kind,
				State: state,
			}
		})
}

// Actives 開催中のコンテスト情報テーブルを辿ってパース
// TODO 未検証
func Actives(doc pig.Node) []*model.Contest {
	return parseContest(doc, pig.ID("contest-table-action"), 0, "active")
}

// Permanents 常時開催のコンテスト情報テーブルを辿ってパース
func Permanents(doc pig.Node) []*model.Contest {
	return parseContest(doc, pig.ID("contest-table-permanent"), 0, "permanent")
}

// Upcomings 開催予定のコンテスト情報テーブルを辿ってパース
func Upcomings(doc pig.Node) []*model.Contest {
	return parseContest(doc, pig.ID("contest-table-upcoming"), 1, "upcoming")
}

// Archives 開催済みのコンテスト情報テーブルを辿ってパース
func Archives(doc pig.Node) []*model.Contest {
	return parseContest(doc, pig.Tag("table"), 1, "archive")
}

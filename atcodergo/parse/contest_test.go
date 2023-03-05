package parse

import (
	"os"
	"reflect"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/tbistr/atcoder-go/atcodergo/model"
	"github.com/tbistr/pig"
)

func TestPermanents(t *testing.T) {
	minimum := `
<div id="contest-table-permanent"><table><tbody><tr>
	<td>
		<span title="Algorithm">Ⓐ</span>
		<span>◉</span>
		<a href="/contests/practice">practice contest</a>
	</td>
	<td>-</td>
</tr></tbody></table></div>`

	for _, tt := range []struct {
		name string
		doc  pig.Node
		want []*model.Contest
	}{{
		"MinimumSample",
		fromString(t, minimum),
		[]*model.Contest{
			{Name: "practice contest", ID: "practice", Kind: "Algorithm", State: "permanent"},
		}}, {
		"UseReal",
		fromFile(t, "testdata/contest.html"),
		[]*model.Contest{
			{Name: "practice contest", ID: "practice", Kind: "Algorithm", State: "permanent"},
			{Name: "AtCoder Library Practice Contest", ID: "practice2", Kind: "Algorithm", State: "permanent"},
		}},
	// TODO Need "real" data(from internet) test?
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permanents(tt.doc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permanents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpcomings(t *testing.T) {
	minimum := `
<div id="contest-table-upcoming"><table><tbody><tr>
	<td class="text-center">
		<a><time class="fixtime fixtime-full">2023-03-11 21:00:00+0900</time></a>
	</td>
	<td>
		<span title="Algorithm">Ⓐ</span>
		<span class="user-blue">◉</span>
		<a href="/contests/abc293">AtCoder Beginner Contest 293</a>
	</td>
	<td class="text-center">01:40</td>
	<td class="text-center">- 1999</td>
</tr></tbody></table></div>`

	for _, tt := range []struct {
		name string
		doc  pig.Node
		want []*model.Contest
	}{{
		"MinimumSample",
		fromString(t, minimum),
		[]*model.Contest{
			{Name: "AtCoder Beginner Contest 293", ID: "abc293", Kind: "Algorithm", State: "upcoming"},
		}}, {
		"UseReal",
		fromFile(t, "testdata/contest.html"),
		[]*model.Contest{
			{Name: "AtCoder Beginner Contest 293", ID: "abc293", Kind: "Algorithm", State: "upcoming"},
			{Name: "AtCoder Regular Contest 158", ID: "arc158", Kind: "Algorithm", State: "upcoming"},
			{Name: "Toyota Programming Contest 2023 Spring Final", ID: "toyota2023spring-final", Kind: "Algorithm", State: "upcoming"},
			{Name: "MC Digital Programming Contest 2023（AtCoder Heuristic Contest 019）", ID: "ahc019", Kind: "Heuristic", State: "upcoming"},
		}},
	// TODO Need "real" data(from internet) test?
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := Upcomings(tt.doc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Upcomings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArchives(t *testing.T) {
	minimum := `
<div id="contest-table-recent"><table><tbody><tr>
	<td class="text-center">
		<a><time class="fixtime fixtime-full">2023-03-04 21:00:00+0900</time></a>
	</td>
	<td>
		<span title="Algorithm">Ⓐ</span>
		<span class="user-blue">◉</span>
		<a href="/contests/abc292">AtCoder Beginner Contest 292</a>
	</td>
	<td class="text-center">01:40</td>
	<td class="text-center">~ 1999</td>
</tr></tbody></table></div>`

	t.Run("MinimumSample", func(t *testing.T) {
		doc := fromString(t, minimum)
		want := []*model.Contest{{Name: "AtCoder Beginner Contest 292", ID: "abc292", Kind: "Algorithm", State: "archive"}}
		if got := Archives(doc); !reflect.DeepEqual(got, want) {
			t.Errorf("Archives() = %v, want %v", got, want)
		}
	})

	for _, tt := range []struct {
		name string
		doc  pig.Node
	}{
		{
			"UseReal1",
			fromFile(t, "testdata/contest_page1.html"),
		},
		{
			"UseReal2",
			fromFile(t, "testdata/contest_page2.html"),
		},
		// TODO Need "real" data(from internet) test?
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := Archives(tt.doc)
			if len(got) == 0 {
				t.Fatal("Archives() = empty")
			}
			pp.Println(got)
			for _, c := range got {
				if c.State != "archive" {
					t.Errorf("Archives().state = %v, want %v", c.State, "archive")
				}
				if c.Name == "" || c.ID == "" || c.Kind == "" {
					t.Errorf("Archives() = %v, want all members are not empty", c)
				}
			}
		})
	}
}

func fromString(t *testing.T, s string) pig.Node {
	doc, err := pig.ParseS(s)
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

func fromFile(t *testing.T, name string) pig.Node {
	b, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	doc, err := pig.ParseB(b)
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

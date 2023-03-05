package parse

import (
	"github.com/tbistr/atcoder-go/atcodergo/model"
	"github.com/tbistr/pig"
)

func Language(doc pig.Node) []model.Language {
	return pig.Map(
		doc.FindDescendant(
			pig.ID("select-lang"),
			pig.Tag("select"),
			pig.And(
				pig.Tag("option"),
				pig.HasAttr("value"),
			)),
		func(n pig.Node) model.Language {
			v, _ := n.AttrVal("value")
			datamime, _ := n.AttrVal("data-mime")
			return model.Language{
				Value:    v,
				Datamime: datamime,
				Text:     n.Text(),
			}
		},
	)
}

func GetCsrf(doc pig.Node) (string, bool) {
	input, ok := doc.Find(pig.And(pig.Tag("input"), pig.HasAttrVal("name", "csrf_token"))).Index(0)
	if !ok {
		return "", false
	}
	return input.AttrVal("value")
}

package atcodergo

import (
	"errors"

	"github.com/tbistr/atcoder-go/atcodergo/model"
	"github.com/tbistr/atcoder-go/atcodergo/parse"
	"github.com/tbistr/pig"
)

// Language stand for
// "<option value={Language.Value} data-mime={Language.Datamime}>{Language.Text}</option>"
type Language = model.Language

// Languages lists up acceptable languages.
func (c *Client) Languages() ([]Language, error) {
	if !c.loggedin {
		return nil, newNeedAuthError("Languages()")
	}

	resp, err := c.Get(BASE_URL.submit("practice").String())
	if err != nil {
		return nil, err
	}
	defer readAllClose(resp.Body)

	doc, err := pig.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	langs := parse.Language(doc)
	if len(langs) == 0 {
		// return nil, newParseError("Languages()")
		return nil, errors.New("parse error: Languages() = empty")
	} else {
		return langs, nil
	}
}

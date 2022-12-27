package atcodergo

import "github.com/PuerkitoBio/goquery"

// Language stand for
// "<option value={Language.Value} data-mime={Language.Datamime}>{Language.Text}</option>"
type Language struct {
	Value    string
	Datamime string
	Text     string
}

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

	Languages := []Language{}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	doc.Find("#select-lang-practice_1 > select > option").Each(func(i int, s *goquery.Selection) {
		// index-0 == "<option></option>"
		if i == 0 {
			return
		}
		// else is "<option value="4001" data-mime="text/x-csrc">C (GCC 9.2.1)</option>"
		value, _ := s.Attr("value")
		datamime, _ := s.Attr("data-mime")
		text := s.Text()
		Languages = append(Languages, Language{Value: value, Datamime: datamime, Text: text})
	})

	return Languages, nil
}

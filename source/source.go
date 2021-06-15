package source

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var sanitize = regexp.MustCompile(`\[[0-9,;:\.\/\\\w\s]*\]`)

func Lookup(url string) (Page, error) {
	res, err := http.Get(url)
	if err != nil {
		return Page{}, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return Page{}, err
	}

	languages := []Language{}
	doc.Find(".interlanguage-link-target").Each(func(i int, s *goquery.Selection) {
		langURL, ok := s.Attr("href")
		if !ok {
			return
		}
		lang, ok := s.Attr("lang")
		if !ok {
			return
		}
		languages = append(languages, Language{
			Link:     langURL,
			Language: lang,
		})
	})

	str := strings.Builder{}

	doc.Find(".mw-parser-output p").EachWithBreak(func(i int, s *goquery.Selection) bool {
		str.WriteString(s.Text())
		return i < 3
	})

	return Page{
		URL:       url,
		Abstract:  sanitize.ReplaceAllLiteralString(str.String(), ""),
		Title:     doc.Find("#firstHeading").Text(),
		Languages: languages,
	}, err
}

type Page struct {
	URL       string
	Abstract  string
	Title     string
	Languages []Language
}

type Language struct {
	Link     string
	Language string
}

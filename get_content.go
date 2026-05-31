package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	h1 := doc.Find("h1").First()
	headingText := h1.Text()
	if headingText == "" {
		h2 := doc.Find("h2").First()
		headingText = h2.Text()
	}
	return headingText
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	main := doc.Find("main")
	var paragraphText string
	if main.Length() > 0 {
		paragraphText = main.Find("p").First().Text()
	} else {
		paragraphText = doc.Find("p").First().Text()
	}
	return strings.TrimSpace(paragraphText)
}

package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	imageURLs := []string{}

	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		// For each '<img src>' it finds, it will run this function.
		src, ok := s.Attr("src")
		if !ok || strings.TrimSpace(src) == "" {
			return
		}

		u, err := url.Parse(src)
		if err != nil {
			fmt.Printf("couldn't parse src %q: %v\n", src, err)
			return
		}

		absolute := baseURL.ResolveReference(u)
		imageURLs = append(imageURLs, absolute.String())
	})
	return imageURLs, nil
}

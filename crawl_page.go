package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()
	pagesCount := len(cfg.pages)
	cfg.mu.Unlock()
	if pagesCount >= cfg.maxPages {
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	if parsedCurrentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v\n", err)
		return
	}

	if !cfg.addPageVisit(normalizedURL) {
		return
	}

	fmt.Printf("Getting html from: %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	// Extract all the data we care about and store it
	pageData := extractPageData(html, rawCurrentURL)
	cfg.setPageData(normalizedURL, pageData)

	// Recurse using the already-extracted outgoing links
	for _, nextURL := range pageData.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}

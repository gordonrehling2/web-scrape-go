package webscraper

import (
	"bytes"

	"golang.org/x/net/html"
)

// GetLinksWithDivClass searches the page for div with class that matches divClass e.g. productInfo
// returning a slice of URLs that match the criteria
// (Note assumption that first and only url in the div is returned)
func GetLinksWithDivClass(page []byte, divClass string) []string {
	// initialise return slice
	urls := make([]string, 0)

	// set up tokenizer
	tok := html.NewTokenizer(bytes.NewReader(page))

	inProductInfo := false
	for {
		tagToken := tok.Next()

		switch tagToken {
		case html.ErrorToken:
			//
			return urls
		case html.StartTagToken: //, html.EndTagToken:
			tagName, _ := tok.TagName()

			switch string(tagName) {
			case "div":
				searchValues := []string{divClass}
				match := attrKeyValuesMatch(tok, "class", searchValues)

				switch match {
				case divClass:
					inProductInfo = true
				}
			case "a":
				if url, matched := getKeyValue(tok, "href"); matched {
					if inProductInfo {
						urls = append(urls, url)
						// Assumes first and only url in the div is returned
						inProductInfo = false
					}
				}
			}
		}
	}
}

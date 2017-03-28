package webscraper

import (
	"bytes"

	"golang.org/x/net/html"
)

// attrKeyValueMatch searches Tag Attributes, returning true if it finds a match e.g. class productInfo
func attrKeyValueMatch(tok *html.Tokenizer, searchKey, searchValue string) bool {
	// Note: need to declare these because can't use ':= tok.TagAttr' with more declared in for loop
	var key []byte
	var value []byte
	// Search each key/value pair in the Attribute
	for more := true; more != false; {
		key, value, more = tok.TagAttr()
		if string(key) == searchKey && string(value) == searchValue {
			return true
		}
	}
	return false
}

// attrKeyMatch searches Tag Attributes
// returning value and true if it finds a match e.g. href
func getKeyValue(tok *html.Tokenizer, searchKey string) (string, bool) {
	// Note: need to declare these because can't use ':= tok.TagAttr' with more declared in for loop
	var key []byte
	var value []byte
	// Search each key in the Attribute
	for more := true; more != false; {
		key, value, more = tok.TagAttr()
		if string(key) == searchKey {
			return string(value), true
		}
	}
	return "", false
}

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
				if matched := attrKeyValueMatch(tok, "class", divClass); matched {
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

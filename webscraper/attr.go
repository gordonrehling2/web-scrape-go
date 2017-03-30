package webscraper

import (
	"golang.org/x/net/html"
)

// attrKeyValueMatches searches Tag Attributes, returning one of the searchValues if it finds a match e.g. class productText
// This is needed when the 'parser' needs to check multiple possible values of Attr
func attrKeyValuesMatch(tok *html.Tokenizer, searchKey string, searchValues []string) string {
	// Note: need to declare these because can't use ':= tok.TagAttr' with more declared in for loop
	var key []byte
	var value []byte
	// Search each key/value pair in the Attribute
	for more := true; more != false; {
		key, value, more = tok.TagAttr()
		if string(key) == searchKey {
			for _, searchValue := range searchValues {
				if string(value) == searchValue {
					return searchValue
				}
			}
		}
	}
	return ""
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

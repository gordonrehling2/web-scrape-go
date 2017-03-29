package webscraper

import (
	"bytes"
	"strconv"
	"strings"

	"log"

	"golang.org/x/net/html"
)

// ProductData contains the data we want to scrape from the page
type ProductData struct {
	title       string
	size        string
	unitPrice   float64
	description string
}

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

// GetPageProductData returns ProductData from the page
// The following patterns are matched:
// Title
//	<div class="productTitleDescriptionContainer">
//		<h1>Sainsbury's Avocado, Ripe & Ready x2</h1>
//	</div>
// Price Per Unit
//	<p class="pricePerUnit">
//		£1.90
// Description
//	<div class="productText">
//	<p>Avocados</p>
func GetPageProductData(page []byte) ProductData {
	productData := ProductData{}
	productData.size = strconv.FormatFloat(float64(len(page)/1000), 'f', -1, 64) + "kb"

	// set up tokenizer
	tok := html.NewTokenizer(bytes.NewReader(page))

	// Title
	inProductTitle := false
	inProductTitleH1 := false

	// PricePerUnit
	inProductSummary := false
	inPricePerUnit := false

	// ProductText
	notSeenOne := true
	inProductText := false
	inProductTextP := false

	for {
		tagToken := tok.Next()

		switch tagToken {
		case html.ErrorToken:
			//
			return productData
		case html.StartTagToken: //, html.EndTagToken:
			tagName, _ := tok.TagName()

			switch string(tagName) {
			case "div":
				searchValues := []string{"productSummary", "productTitleDescriptionContainer", "productText", "mainProductInfoWrapper"}
				match := attrKeyValuesMatch(tok, "class", searchValues)

				switch match {
				case "productSummary":
					inProductSummary = true
				case "productTitleDescriptionContainer":
					inProductTitle = true
				case "productText":
					inProductText = true
				case "mainProductInfoWrapper":
					inProductSummary = false
				}
			case "h1":
				if inProductTitle {
					inProductTitleH1 = true
				}
			case "p":
				if inProductText {
					inProductTextP = true
					continue
				}
				if matched := attrKeyValueMatch(tok, "class", "pricePerUnit"); matched {
					inPricePerUnit = true
				}
			}
		case html.TextToken:
			if inProductTitle && inProductTitleH1 {
				productData.title = string(tok.Text())
				inProductTitle = false
				inProductTitleH1 = false
				continue
			}
			if inProductSummary && inPricePerUnit {
				// Split token on '£'
				tmp := strings.Split(string(tok.Text()), "£")
				// Use the bit after the '£'
				f, err := strconv.ParseFloat(tmp[1], 64)
				if err != nil {
					log.Fatal(err)
				}
				productData.unitPrice = f
				inPricePerUnit = false
				continue
			}
			if inProductTextP && notSeenOne {
				productData.description = string(tok.Text())
				inProductText = false
				inProductTextP = false
				notSeenOne = false
				continue
			}
		}
	}
	return productData
}

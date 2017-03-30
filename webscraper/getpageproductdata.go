package webscraper

import (
	"bytes"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// ProductData contains the data we want to scrape from the page
type ProductData struct {
	Title       string  `json:"title"`
	Size        string  `json:"size"`
	UnitPrice   float64 `json:"unit_price"`
	Description string  `json:"description"`
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
	productData.Size = strconv.FormatFloat(float64(len(page)/1024), 'f', -1, 64) + "kb"

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
				productData.Title = string(tok.Text())
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
				productData.UnitPrice = f
				inPricePerUnit = false
				continue
			}
			if inProductTextP && notSeenOne {
				productData.Description = string(tok.Text())
				inProductText = false
				inProductTextP = false
				notSeenOne = false
				continue
			}
		}
	}
	return productData
}

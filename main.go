package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gordonrehling2/web-scrape-go/webscraper"
)

// USER STORY
//
// As a software developer
// I want to consume product item data from a web page and recompose it in JSON
// So that it can be more easily re-purposed

// GOL was updated on Mon 27-Mar-17 and target URL changed
//var targetURL = "http://www.sainsburys.co.uk/webapp/wcs/stores/servlet/CategoryDisplay?listView=tr ue&orderBy=FAVOURITES_FIRST&parent_category_rn=12518&top_category=125 18&langId=44&beginIndex=0&pageSize=20&catalogId=10137&searchTerm=&categ oryId=185749&listId=&storeId=10151&promotionId=#langId=44&storeId=10151&cat alogId=10137&categoryId=185749&parent_category_rn=12518&top_category=1251 8&pageSize=20&orderBy=FAVOURITES_FIRST&searchTerm=&beginIndex=0&hide Filters=true"
var targetURL = "http://www.sainsburys.co.uk/shop/gb/groceries/fruit-veg/ripe---ready#langId=44&storeId=10151&catalogId=10241&categoryId=185749&parent_category_rn=12518&top_category=12518&pageSize=20&orderBy=FAVOURITES_ONLY%7CTOP_SELLERS&searchTerm=&beginIndex=0"

// Result is JSON output for the test
type Result struct {
	Results []webscraper.ProductData `json:"results"`
	Total   float64                  `json:"total"`
}

func fatalBadURL(url string, err error) {
	log.Printf("can't open %s due to error '%s'\n", url, err.Error())
	log.Fatal("giving up...")
}

func processURLs(urls []string) Result {
	// Create result structure
	result := Result{}

	// Use int to sum unit prices, to avoid rounding errors
	pence := 0
	for _, url := range urls {
		// get product page
		page, err := webscraper.GetWebPage(url)
		if err != nil {
			fatalBadURL(url, err)
		}
		// get the required product data from the product page
		product := webscraper.GetPageProductData(page)
		// keep running total of unit prices
		pence += int(product.UnitPrice * 100)
		// append the product as a slice in the result
		result.Results = append(result.Results, product)
	}
	// convert pence to pounds
	result.Total = float64(pence) / 100

	return result
}

func main() {
	// get the target page
	page, err := webscraper.GetWebPage(targetURL)
	if err != nil {
		fatalBadURL(targetURL, err)
	}

	// get all the product URLs from the page
	urls := webscraper.GetLinksWithDivClass(page, "productInfo")

	// process the URLs
	result := processURLs(urls)

	// Create the actual JSON
	// actualJSON, err := json.Marshal(result)

	// Create pretty JSON, with 3 space indent
	prettyJSON, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	// Output pretty JSON
	fmt.Println(string(prettyJSON))
}

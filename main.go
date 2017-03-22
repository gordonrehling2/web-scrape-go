package main

import (
	"github.com/gordonrehling2/web-scrape-go/webscraper"
	"log"
)

// USER STORY
//
// As a software developer
// I want to consume product item data from a web page and recompose it in JSON
// So that it can be more easily re-purposed

var targetUrl = "http://www.sainsburys.co.uk/webapp/wcs/stores/servlet/CategoryDisplay?listView=tr ue&orderBy=FAVOURITES_FIRST&parent_category_rn=12518&top_category=125 18&langId=44&beginIndex=0&pageSize=20&catalogId=10137&searchTerm=&categ oryId=185749&listId=&storeId=10151&promotionId=#langId=44&storeId=10151&cat alogId=10137&categoryId=185749&parent_category_rn=12518&top_category=1251 8&pageSize=20&orderBy=FAVOURITES_FIRST&searchTerm=&beginIndex=0&hide Filters=true"

func main() {
	page, err := webscraper.GetWebPage(targetUrl)
	if err != nil {
		log.Printf("can't open %s due to error '%s'\n", targetUrl, err.Error())
	}

	log.Println(string(page))
}

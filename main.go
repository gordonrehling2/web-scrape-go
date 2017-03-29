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
var cookie = "_ga=GA1.3.1970836756.1488187879; Apache=10.173.10.13.1490193304969911; SESSION_COOKIEACCEPT=true; WC_SESSION_ESTABLISHED=true; WC_PERSISTENT=4yyGYA6T%2B1vyUboZ8etusrnfm9E%3D%0A%3B2017-03-22+14%3A35%3A04.979_1490193304972-34312_10151; WC_ACTIVEPOINTER=44%2C10151; WC_USERACTIVITY_-1002=-1002%2C10151%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2CbGBpl3Qui0C98mq56pJydYv6G3T2uqiEt4H6OT3YLzbP5I3c%2F8WPGKX2vn2NxzbrCYo%2B6N7rOnvKwrO5wCWln3H2lTC2IaS1v7W1GTqW9TglJ5oOcMe30I5VWg7XHPDa7b%2FE9RZOIBkzuhDCJLC%2FRKzodgubDSBY3GBwWpXk7NGAtUG8Ec%2FuJ%2BvcHZYwxv8CsvE2hCxYXZaFfzt7kHNtDQ%3D%3D; WC_GENERIC_ACTIVITYDATA=[17334505011%3Atrue%3Afalse%3A0%3Asc%2BhJxcA2DeeBSkyrdlQwdAr7Tc%3D][com.ibm.commerce.context.audit.AuditContext|1490193304972-34312][com.ibm.commerce.store.facade.server.context.StoreGeoCodeContext|null%26null%26null%26null%26null%26null][CTXSETNAME|Store][com.sol.commerce.context.SolBusinessContext|null%26null%26null%26null%26null%26null%26null%26null%26false%26false%26false%26null%26false%26null%26false%26false%26null%26false%26false%26false%26null%26false%26false][com.ibm.commerce.context.globalization.GlobalizationContext|44%26GBP%2644%26GBP][com.ibm.commerce.catalog.businesscontext.CatalogContext|10241%26null%26false%26false%26false][com.ibm.commerce.context.preview.PreviewContext|null%26null%26false][com.ibm.commerce.context.base.BaseContext|10151%26-1002%26-1002%26-1][com.ibm.commerce.context.experiment.ExperimentContext|null][com.ibm.commerce.context.entitlement.EntitlementContext|10502%2610502%26null%26-2000%26null%26null%26null]; sbrycookie1=630263751; bt_sc_serialize=14901933295941380659; _msuuid_529av24112=D6DCA46B-19EA-454F-A1AB-B82BCBB1C8C1; LIST_VIEW=true; TS017d4e39_77=08036bb980ab280017dbdfe5ded5fadb2fe57af8087ccef9e0d02cf5376199b48cb2b81708cc38b3760d381fff0761940896387c80823800c41e22b1b4bdf7b34b04bb8fdf0b0a2456f1dc985c71296b0e7c2c9760a0220c0d1fd2a1fe4d4a5e43c5c02cc6b802868f155b8b36d1d795; s_sq=; c=undefinedDirect%20LoadDirect%20Load; JSESSIONID=0000FHLZI3qYIgk_9XhAl-YlaaY:17cj2mr6e; mmapi.store.p.0=%7B%22mmparams.d%22%3A%7B%7D%2C%22mmparams.p%22%3A%7B%22pd%22%3A%221522234591939%7C%5C%22-1805000548%7CKQAAAAoBQvndG8mVDpMgQQAEAHqD0AnJddRIDwAAAFkB%2FKMwcdRIAAAAAP%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FABR3d3cuc2FpbnNidXJ5cy5jby51awKVDgQAAAAAAAAAAAAAanwBAGp8AQBqfAEAAAAAAAABRQ%3D%3D%5C%22%22%2C%22srv%22%3A%221522234591946%7C%5C%22fravwcgeu04%5C%22%22%7D%7D; mmapi.store.s.0=%7B%22mmparams.d%22%3A%7B%7D%2C%22mmparams.p%22%3A%7B%7D%7D; oa_products=|; oa_visited=|; oa_revenueincdel=|; oa_time=|; loginClicked=|; bt_product_click=; bt_hplinks=; WRUIDAWS=1148893853491278; __CT_Data_194=gpv=15&apv_194_www16=15&cpv_194_www16=15&rpv_194_www16=15; bt_espot_pageCtr=2; s_fid=38081712F5DEBF5D-165E578396B960A9; s_vnum=1491001200097%26vn%3D4; s_invisit=true; gpv_page=groceries%3Afruit%20%26%20veg%3Afresh%20fruit%3Aripe%20%26%20ready; s_cc=true; sbrycookie2=LlrHbjOczkHbzvDkHbQKUC; TS017d4e39=01a4f6ca7c817d13306ad79c873a7164205ba6fc90c93223d8edd508e6205e7b64846853597e0c8dd95df98e3df5078c1b9296e5966bcc3596f04aa050d511069f7c1b47c752a0b776e558e01eef3c3b8a1b2276e5c0f342a19f294489e4716b71cc0d398c69b74b75e56826d426caf7ffb7add74ba15653782ec3481b64c68ad9e435cddec35e2a0402254c142239dcc3d3c6cc59404a1a9ae6da46c980c43033c570b2042c4dd144026de6abc5a925c46e1e1cdd531bb1fc1e39eff3827d261ac1e524b11b7a498aa16c2286d6bbc5cf284d8872; TS01cd64fc=01a4f6ca7cee59522859a1140d035e70763c8523be12a52d562e7645fbbe0404a101dadd5bb446eeed199dd01e77af3779fc9cfad9; ctm={'pgv':2827812988202127|'vst':6661944893441759|'vstr':2243543643446053|'intr':1490699143234|'v':1|'lvst':1274}; bt_productclickfrom="

// Result is JSON output for the test
type Result struct {
	Results []webscraper.ProductData `json:"results"`
	Total   float64                  `json:"total"`
}

func fatalBadURL(url string, err error) {
	log.Printf("can't open %s due to error '%s'\n", url, err.Error())
	log.Fatal("giving up...")
}

func main() {
	// get the target page
	page, err := webscraper.GetWebPage(targetURL, cookie)
	//fmt.Println(string(page))
	if err != nil {
		fatalBadURL(targetURL, err)
	}

	// get all the product URLs
	urls := webscraper.GetLinksWithDivClass(page, "productInfo")

	// Create result structure
	result := Result{}
	// Use int to sum unit prices, to avoid rounding errors
	pence := 0
	for _, url := range urls {
		// get product page
		page, err := webscraper.GetWebPage(url, cookie)
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

	// This is actualJSON
	// actualJSON, err := json.Marshal(result)

	// Create prettyJSON with 3 space indent
	prettyJSON, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	// Output prettyJSON
	fmt.Println(string(prettyJSON))
}

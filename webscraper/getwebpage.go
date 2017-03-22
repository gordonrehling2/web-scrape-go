package webscraper

import (
	"io/ioutil"
	"net/http"
)

func GetWebPage(url string) (page []byte, err error) {
	// Get page
	response, err := http.Get(url)
	if err != nil {
		// http.Get error pass upwards
		return
	}
	defer response.Body.Close()

	page, err = ioutil.ReadAll(response.Body)
	if err != nil {
		// ioutil.ReadAll error pass upwards
		return
	}

	return
}

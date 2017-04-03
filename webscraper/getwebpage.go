package webscraper

import (
	"io/ioutil"
	"net/http"
	"time"
)

// GetWebPage gets the page for the given url, with a valid cookie
func GetWebPage(url string) (page []byte, err error) {
	// use a client and use the CookieJar interface
	client := &http.Client{nil, nil, jar, time.Duration(10) * time.Second}
	request, err := http.NewRequest("GET", url, nil)

	response, err := client.Do(request)
	if err != nil {
		// http.Get error pass upwards
		return
	}
	defer response.Body.Close()

	page, err = ioutil.ReadAll(response.Body)
	if err != nil {
		// pass ioutil.ReadAll error upwards
		return
	}

	return
}

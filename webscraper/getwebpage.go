package webscraper

import (
	"io/ioutil"
	"net/http"
)

// GetWebPage gets the page for the given url, with a valid cookie
func GetWebPage(url, cookie string) (page []byte, err error) {
	// Get page
	//response, err := http.Get(url)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	// When GOL updated on Mon 27-Mar-17, link changed and need to add this cookie Header to get data back
	request.Header.Set("Cookie", cookie)
	response, err := client.Do(request)
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

package xkcd

import (
	"errors"
	"io/ioutil"
	"regexp"

	"net/http"
)

// GetXKCDImageLink gets a random image from XKCD by calling the /random/comic endpoint and parsing the actual image URL
func GetXKCDImageLink() (imageURL string, err error) {
	imageURL = ""

	// the URL and endpoint are fixed for XKCD
	response, err := http.Get("https://c.xkcd.com/random/comic/")
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	imageURL, err = parseImageLink(&body)

	return
}

func parseImageLink(body *[]byte) (link string, err error) {
	re := regexp.MustCompile(`Image URL \(for hotlinking/embedding\):.+(https://.*(png|jpg|gif))`)

	regexMatch := re.FindSubmatch(*body)
	if regexMatch == nil {
		err = errors.New("No image has been found")
		return // we have to return, else regexMatch[1] will cause a panic
	}
	link = string(regexMatch[1])

	return
}

package xkcd

import (
	"errors"
	"regexp"

	"github.com/imroc/req"
)

// GetXKCDImageLink gets a random image from XKCD by calling the /random/comic endpoint and parsing the actual image URL
func GetXKCDImageLink() (imageURL string, err error) {
	imageURL = ""

	// the URL and endpoint are fixed for XKCD
	response, err := req.Get("https://c.xkcd.com/random/comic/")
	if err != nil {
		return
	}
	resp := response.String()

	imageURL, err = parseImageLink(&resp)

	return
}

func parseImageLink(body *string) (link string, err error) {
	re := regexp.MustCompile(`Image URL \(for hotlinking/embedding\):.+(https://.*(png|jpg|gif))`)

	regexMatch := re.FindStringSubmatch(*body)
	if regexMatch == nil {
		err = errors.New("No image has been found")
		return // we have to return, else regexMatch[1] will cause a panic
	}
	link = regexMatch[1]

	return
}

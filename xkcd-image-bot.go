package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/imroc/req"
)

// JSONPayload actually generates our message payload
type JSONPayload struct {
	PostText string `json:"text"`
}

func main() {
	var help = flag.Bool("help", false, "print this help and exit")
	var webhookURL = flag.String("webhook", "", "the URL of the webhook to post the image to")
	var customMessage = flag.String("message", "", "(optional) an additional message to post before the image")

	flag.Parse()
	if *help {
		fmt.Print("Usage:\n\n")
		flag.PrintDefaults()
	}

	imageURL := getXKCDImageLink()
	payload := buildPayload(imageURL, *customMessage)
	postToWebhook(*webhookURL, payload)
}

// getXKCDImageLink gets a random image from XKCD by calling the /random/comic endpoint and parsing the actual image URL
func getXKCDImageLink() (imageURL string) {
	imageURL = ""

	// the URL and endpoint are fixed for XKCD
	response, err := req.Get("https://c.xkcd.com/random/comic/")
	if err != nil {
		PrintErrorAndExit("Error getting the random comic", err)
	}
	resp := response.String()

	// parse the image URL from the returned HTML body
	re := regexp.MustCompile("Image URL.* (https://.*png)")
	imageURL = re.FindStringSubmatch(string(resp))[1]

	return
}

// buildPayload builds the payload from the custom message and the image URL
func buildPayload(imgURL string, message string) (payload string) {
	var textStringBuilder strings.Builder
	if message != "" {
		// fix newline errors by replacing escaped versions of the string with an actual "\ + n"
		messageEscaped := strings.Replace(message, `\n`, "\n", -1)
		textStringBuilder.WriteString(messageEscaped)
		textStringBuilder.WriteString("\n")
	}
	textStringBuilder.WriteString(imgURL)

	text := textStringBuilder.String()

	// Slack only needs text in its payload
	payloadStruct := JSONPayload{
		PostText: text,
	}

	payloadJSON, err := json.Marshal(payloadStruct)
	if err != nil {
		PrintErrorAndExit("marshalling payload", err)
	}
	payload = string(payloadJSON)

	return
}

// postToWebhook actually posts the payload to the webhook
func postToWebhook(webhookURL string, payload string) {

	// set the appropriate header for posting to the API
	header := req.Header{
		"Accept": "application/json",
	}

	// post the image to the app and check if there was an error
	// does not, however, check the HTTP status code!
	_, err := req.Post(webhookURL, header, payload)
	if err != nil {
		PrintErrorAndExit("posting the payload", err)
	}
}

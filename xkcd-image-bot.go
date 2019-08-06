package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/imroc/req"
)

// custom struct for the Mattermost payload
type JsonPayloadMattermost struct {
	MattermostUsername string `json:"username"`
	MattermostChannel  string `json:"channel"`
	PostText           string `json:"text"`
}

// custom struct for the Slack payload
type JsonPayloadSlack struct {
	PostText string `json:"text"`
}

func main() {
	// these are all the command line flags we declare
	// webhookURL is definitely required, as well as one of useSlack or useMattermost
	// mattermostChannel and mattermostUsername are only available if using mattermost, as slack does not support this feature
	var help = flag.Bool("help", false, "print this help and exit")
	var webhookURL = flag.String("webhook-url", "", "the URL of the webhook to post the image to")
	var useSlack = flag.Bool("slack", false, "post to Slack")
	var useMattermost = flag.Bool("mattermost", false, "post to Mattermost")
	var mattermostChannel = flag.String("channel", "town-square", "(optional) which channel to post to; only available when using Mattermost")
	var mattermostUsername = flag.String("username", "xkdc-image-bot", "(optional) which username to post as; only available when using Mattermost")

	// parse all given flags and do some basic validity checks on them
	flag.Parse()
	if *help {
		fmt.Print("Usage:\n\n")
		flag.PrintDefaults()
	}
	CheckValidityOfFlags(*webhookURL, *useMattermost, *useSlack)

	// get a random Image and then pass it along to be posted to the app of choice
	imageURL := ParseXKCD()
	PostToApp(imageURL, *webhookURL, *useMattermost, *mattermostChannel, *mattermostUsername)
}

// CheckValidityOfFlags performs some basic validity checks on the user-provided input
func CheckValidityOfFlags(webhookURL string, useMattermost bool, useSlack bool) {
	var isNotValid = false

	// a webhook should begin with https:// and contain at leas 4 more characters after that,
	// although an actual webhook URL would likely contain more characters
	webhookRegex := regexp.MustCompile("https://.{12,}")
	if !webhookRegex.MatchString(webhookURL) {
		fmt.Println("Webhook URL must contain 'https' and be of length 12, at least")
		isNotValid = true
	}
	// it does not make sense to post to Slack and Mattermost at the same time,
	if useMattermost && useSlack {
		fmt.Println("You cannot post to Mattermost and Slack at the same time!")
		isNotValid = true
	}
	// ...or not at all
	if !useMattermost && !useSlack {
		fmt.Println("You must specify either -mattermost or -slack")
		isNotValid = true
	}
	// if the basic validity checks fail, exit the program and print the usage
	if isNotValid {
		fmt.Println("\nUsage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

// ParseXKCD gets a random image from XKCD by calling the /random/comic endpoint and parsing the actual Image URL
func ParseXKCD() string {
	// the URL and endpoint are fixed for XKCD
	response, err := req.Get("https://c.xkcd.com/random/comic/")
	if err != nil {
		PrintErrorAndExit("getting the random comic", err)
	}
	resp := response.String()

	// parse the image URL from the returned HTML body
	re := regexp.MustCompile("Image URL.* (https://.*png)")
	imageURL := re.FindStringSubmatch(string(resp))[1]
	fmt.Println("[I] Posting this image: ", imageURL)

	return imageURL
}

// BuildPayloadSlack builds the payload specific for Slack
func BuildPayloadSlack(imgURL string) string {
	var imgURLString strings.Builder
	imgURLString.WriteString(imgURL)
	text := imgURLString.String()

	// Slack only needs text in its payload
	slackPayload := JsonPayloadSlack{
		PostText: text,
	}

	payloadJson, err := json.Marshal(slackPayload)
	if err != nil {
		PrintErrorAndExit("marshalling payload", err)
	}
	payload := string(payloadJson)

	return payload
}

// BuildPayloadMattermost builds the payload specific for Mattermost
func BuildPayloadMattermost(imgURL string, channel string, username string) string {
	var imgURLString strings.Builder
	imgURLString.WriteString(imgURL)
	text := imgURLString.String()

	// Mattermost is specific in that it requires "usnername" and "channel" inside the payload
	mattermostPayload := JsonPayloadMattermost{
		MattermostChannel:  channel,
		MattermostUsername: username,
		PostText:           text,
	}

	payloadJson, err := json.Marshal(mattermostPayload)
	if err != nil {
		PrintErrorAndExit("marshalling payload", err)
	}
	payload := string(payloadJson)

	return payload
}

// PrintErrorAndExit takes a custom message and error, prints it and then exits the program
func PrintErrorAndExit(customMessage string, errorMessage error) {
	fmt.Println("[E] Error ", customMessage, ":")
	fmt.Println(errorMessage)
	os.Exit(1)
}

// PostToApp builds the payload with the image URL an posts it to the specified app
func PostToApp(imgURL string, webhookURL string, useMattermost bool, mattermostUsername string, mattermostChannel string) {
	var payload string = ""

	// depending on the user's choice, we either build our payload for Mattermost or for Slack
	if useMattermost {
		payload = BuildPayloadMattermost(imgURL, mattermostUsername, mattermostChannel)
	} else {
		payload = BuildPayloadSlack(imgURL)
	}

	// set the appropriate header for posting to the API
	header := req.Header{
		"Accept": "application/json",
	}

	// post the image to the app and check if there was an error
	// does not, however, check the HTTP status code!
	response, err := req.Post(webhookURL, header, payload)
	if response != nil {
		fmt.Println("[I] Response from server: ", response)
	} else {
		PrintErrorAndExit("posting the payload", err)
	}
}

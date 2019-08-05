package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/imroc/req"
)

func main() {
	// these are all command line flags we declare
	// webhookURL is definitely required, as well as one of useSlack or useMattermost
	// mattermostChannel and mattermostUsername are only needed if using mattermost, as slack does not support this feature
	var help = flag.Bool("help", false, "print this help and exit")
	var webhookURL = flag.String("webhook-url", "", "the URL of the webhook to post the image to")
	var useSlack = flag.Bool("slack", false, "post to Slack")
	var useMattermost = flag.Bool("mattermost", false, "post to Mattermost")
	var mattermostChannel = flag.String("channel", "town-square", "which channel to post to (only available when using Mattermost)")
	var mattermostUsername = flag.String("username", "xkdc-image-bot", "which username to post as (only available when using Mattermost)")

	// parse all given flags and do some basic validity checks on them
	flag.Parse()
	if *help {
		fmt.Println("Usage:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	CheckValidityOfFlags(*webhookURL, *useMattermost, *useSlack)

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
	// or not at all
	if useMattermost && useSlack {
		fmt.Println("You cannot post to Mattermost and Slack at the same time!")
		isNotValid = true
	}
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
	// the URL and endpoint are fixed
	response, err := req.Get("https://c.xkcd.com/random/comic/")
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	resp := response.String()

	// parse the image URL from the returned HTML body
	re := regexp.MustCompile("Image URL.* (https://.*png)")
	imageURL := re.FindStringSubmatch(string(resp))[1]
	fmt.Print("Posting this image: ")
	fmt.Println(imageURL)

	return imageURL
}

// BuildPayloadSlack builds the payload specific for Slack
func BuildPayloadSlack(imgURL string) string {
	var imgURLString strings.Builder
	imgURLString.WriteString(imgURL)
	text := imgURLString.String()

	// Manually building the JSON gets pretty ugly, but using encoding/json has thrown errors
	// TODO: check encoding/json again and refactor code to use that module
	var payloadStringBuilder strings.Builder
	payloadStringBuilder.WriteString("{")
	payloadStringBuilder.WriteString("\"text\": \"")
	payloadStringBuilder.WriteString(text)
	payloadStringBuilder.WriteString("\"}")

	payload := payloadStringBuilder.String()

	return payload
}

// BuildPayloadMattermost builds the payload specific for Mattermost
func BuildPayloadMattermost(imgURL string, channel string, username string) string {
	var imgURLString strings.Builder
	imgURLString.WriteString(imgURL)
	text := imgURLString.String()

	// Mattermost is specific in that it requires "usnername" and "channel" inside the payload
	// Manually building the JSON gets pretty ugly, but using encoding/json has thrown errors
	// TODO: check encoding/json again and refactor code to use that module
	var payloadStringBuilder strings.Builder
	payloadStringBuilder.WriteString("{")
	payloadStringBuilder.WriteString("\"username\": \"")
	payloadStringBuilder.WriteString(username)
	payloadStringBuilder.WriteString("\",")
	payloadStringBuilder.WriteString("\"channel\": \"")
	payloadStringBuilder.WriteString(channel)
	payloadStringBuilder.WriteString("\",")
	payloadStringBuilder.WriteString("\"text\": \"")
	payloadStringBuilder.WriteString(text)
	payloadStringBuilder.WriteString("\"}")

	payload := payloadStringBuilder.String()

	return payload
}

// PostToApp builds the payload with the image URL an posts it to the specified app
func PostToApp(imgURL string, webhookURL string, useMattermost bool, mattermostUsername string, mattermostChannel string) {
	var payload string = ""

	if useMattermost {
		payload = BuildPayloadMattermost(imgURL, mattermostUsername, mattermostChannel)
	} else {
		payload = BuildPayloadSlack(imgURL)
	}

	header := req.Header{
		"Accept": "application/json",
	}

	response, err := req.Post(webhookURL, header, payload)
	if response != nil {
		fmt.Println(response)
	} else {
		fmt.Println("There was an error:")
		fmt.Println(err)
	}
}

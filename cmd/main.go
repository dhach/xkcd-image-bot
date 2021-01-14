package main

import (
	"flag"
	"fmt"
	"os"

	"xkcd-image-bot/pkg/helpers"
	"xkcd-image-bot/pkg/post"
	"xkcd-image-bot/pkg/xkcd"
)

func main() {
	webhookURL, customMessage := parseFlags()
	imageURL, errGetLink := xkcd.GetXKCDImageLink()
	if errGetLink != nil {
		helpers.PrintErrorAndExit(errGetLink)
	}

	errPost := post.ToWebhook(webhookURL, imageURL, customMessage)
	if errPost != nil {
		helpers.PrintErrorAndExit(errPost)
	}

	fmt.Printf("Posted '%s' to '%s'\n", imageURL, webhookURL)
}

func parseFlags() (string, string) {
	// TODO: rewrite to flag.StringVar to get rid of pointers
	help := flag.Bool("help", false, "print this help and exit")
	webhook := flag.String("webhook", "", "the URL of the webhook to post the image to")
	message := flag.String("message", "", "(optional) an additional message to post before the image")

	flag.Parse()
	if *help {
		fmt.Print("Usage:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *webhook == "" {
		fmt.Println("'-webhook' cannot be empty!")
		flag.PrintDefaults()
		os.Exit(0)
	}

	return *webhook, *message
}

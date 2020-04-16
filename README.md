# xkcd-image-bot

Build status: ![Go](https://github.com/dhach/xkcd-image-bot/workflows/Go/badge.svg)

A program to post a random XKCD image to either Mattermost or Slack.

Written for personal use and for learing Go.

## Usage

You have to choose between posting to Mattermost or Slack. Both Apps need a webhook URL.

You can pass the reguired arguments via command line flags:


**Supported command line flags**
```
Usage:
  -channel string
    	which channel to post to (only available when using Mattermost) (default "town-square")
  -help
    	print this help and exit
  -mattermost
    	post to Mattermost
  -slack
    	post to Slack
  -username string
    	which username to post as (only available when using Mattermost) (default "xkdc-image-bot")
  -webhook-url string
    	the URL of the webhook to post the image to
```

**Example Usage:**
```
./xkcd-image-bot -webhook-url https://my.private-mattermost.xyz/hooks/8t072k984xt7hjhjdahf\
 -mattermost -username xkcd-image-bot -channel test123456789
```

## Issues
If you find any bugs, please report them via Github

## License
[MIT License](LICENSE)


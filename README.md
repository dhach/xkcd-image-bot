# xkcd-image-bot

Build status:  ![Go](https://github.com/dhach/xkcd-image-bot/workflows/Go/badge.svg)

A program to post a random XKCD image to Slack or Slack-compatible webhooks (currently tested and verfified: Mattermost, Discord).

Written for personal use and for learing Go.

## Usage

```raw
Usage:

  -help
    	print this help and exit
  -message string
    	(optional) an additional message to post before the image
  -webhook string
    	the URL of the webhook to post the image to
```

## Example Usages

```bash

./xkcd-image-bot -webhook-url https://my.private-mattermost.xyz/hooks/{webhook-id}
./xkcd-image-bot -webhook-url https://my.private-mattermost.xyz/hooks/{webhook-id} -message "good morning!"
./xkcd-image-bot -webhook-url https://hooks.slack.com/services/{webhook-id}
./xkcd-image-bot -webhook-url https://discord.com/api/webhooks/{webhook.id}/{webhook.token}/slack -message "hello, world!"
```

## Issues

If you find any bugs, please report them via Github.

## License

[MIT License](LICENSE)

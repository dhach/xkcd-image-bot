# xkcd-image-bot

Build status:  ![Go](https://github.com/dhach/xkcd-image-bot/workflows/Go/badge.svg)

A program to post a random XKCD image Slack or service compatible to Slacks webhook format.

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
./xkcd-image-bot -webhook-url https://my.private-mattermost.xyz/hooks/8t072k984xt7hjhjdahf
./xkcd-image-bot -webhook-url https://my.private-mattermost.xyz/hooks/8t072k984xt7hjhjdahf -message "good morning!"
```

## Issues

If you find any bugs, please report them via Github

## License

[MIT License](LICENSE)

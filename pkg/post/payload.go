package post

import (
	"encoding/json"
	"strings"
)

// JSONPayload defines the structure of the final payload sent to the webhook
type JSONPayload struct {
	PostText string `json:"text"`
}

// buildPayload builds the payload from the custom message and the image URL
func buildPayload(imgURL string, message string) (payload []byte, err error) {
	var textStringBuilder strings.Builder
	if message != "" {
		// fix newline errors by replacing runes with an actual string "\n"
		messageEscaped := strings.Replace(message, `\n`, "\n", -1)
		textStringBuilder.WriteString(messageEscaped)
		textStringBuilder.WriteString("\n")
	}
	textStringBuilder.WriteString(imgURL)

	text := textStringBuilder.String()

	payloadStruct := JSONPayload{
		PostText: text,
	}
	payload, err = json.Marshal(payloadStruct)

	return
}

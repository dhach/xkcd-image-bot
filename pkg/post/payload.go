package post

import (
	"encoding/json"
	"strings"
)

// buildPayload builds the payload from the custom message and the image URL
func buildPayload(imgURL string, message string) (payload []byte, err error) {
	type JSONPayload struct {
		PostText string `json:"text"`
	}
	var textStringBuilder strings.Builder
	if message != "" {
		// fix newline errors by replacing escaped versions of the string with an actual "\n"
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

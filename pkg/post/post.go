package post

// postToWebhook actually posts the payload to the webhook
import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

// ToWebhook posts the constructed payload to the specified URL
func ToWebhook(webhookURL string, imgURL string, customMessage string) error {
	payload, payloadErr := buildPayload(imgURL, customMessage)
	if payloadErr != nil {
		return payloadErr
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewReader(payload))
	if err == nil && resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("Post to webhook returned status code %d", resp.StatusCode)
		err = errors.New(errMsg)
	}

	return err
}

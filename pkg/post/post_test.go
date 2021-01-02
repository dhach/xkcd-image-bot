package post

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToWebhook(t *testing.T) {
	type args struct {
		webhookURL    string
		imgURL        string
		customMessage string
	}
	tests := map[string]struct {
		args    args
		wantErr bool
	}{
		"invalid_webhook_URL": {
			args:    args{"http://invalid", "https://imgs.xkcd.com/comics/hack.png", "foo"},
			wantErr: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			if err := ToWebhook(tt.args.webhookURL, tt.args.imgURL, tt.args.customMessage); (err != nil) != tt.wantErr {
				t.Errorf("ToWebhook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMocked_ToWebhook(t *testing.T) {
	tests := map[string]struct {
		argImgURL        string
		argCustomMessage string
		wantErr          bool
	}{
		"regular":    {"https://imgs.xkcd.com/comics/hack.png", "foo", false},
		"whitespace": {"https://imgs.xkcd.com/comics/hack.png", " ", false},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			mockedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				reqBody, _ := ioutil.ReadAll(r.Body)
				wantPayload := fmt.Sprintf(`{"text":"%s\n%s"}`, tt.argCustomMessage, tt.argImgURL)
				w.WriteHeader(200)
				fmt.Fprintf(w, "all good!")

				assert.Equal(t, string(reqBody), wantPayload)
			}))
			err := ToWebhook(mockedServer.URL, tt.argImgURL, tt.argCustomMessage)

			if tt.wantErr {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestMockedFailure_ToWebhook(t *testing.T) {
	mockedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintf(w, "nope!")
	}))
	err := ToWebhook(mockedServer.URL, "https://nonexistent", "")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Post to webhook returned status code 404")

}

package post

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test_buildPayload(t *testing.T) {
	type args struct {
		imgURL  string
		message string
	}
	tests := map[string]struct {
		args    args
		wantErr bool
	}{
		"regular": {
			args:    args{"https://imgs.xkcd.com/comics/hack.png", "hello!"},
			wantErr: false,
		},
		"empty": {
			args:    args{" ", " "},
			wantErr: false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			wantPayload, _ := json.Marshal(JSONPayload{PostText: fmt.Sprintf("%s\n%s", tt.args.message, tt.args.imgURL)})

			gotPayload, err := buildPayload(tt.args.imgURL, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPayload, wantPayload) {
				t.Errorf("buildPayload() = %s, want %s", gotPayload, wantPayload)
			}
		})
	}
}

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
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "regular",
			args:    args{"https://imgs.xkcd.com/comics/hack.png", "hello!"},
			wantErr: false,
		},
		{
			name:    "empty",
			args:    args{" ", " "},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

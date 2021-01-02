package xkcd

import "testing"

func Test_parseImageLink(t *testing.T) {
	tests := map[string]struct {
		body     []byte
		wantLink string
		wantErr  bool
	}{
		"Good": {
			[]byte(`Image URL (for hotlinking/embedding): https://imgs.xkcd.com/comics/hack.png`),
			"https://imgs.xkcd.com/comics/hack.png",
			false,
		},
		"Error": {
			[]byte(`some garbage here`),
			"",
			true,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			gotLink, err := parseImageLink(&tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseImageLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLink != tt.wantLink {
				t.Errorf("parseImageLink() = %v, want %v", gotLink, tt.wantLink)
			}
		})
	}
}

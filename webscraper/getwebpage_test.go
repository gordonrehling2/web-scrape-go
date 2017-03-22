package webscraper

import (
	"testing"
)

func TestGetWebPage(t *testing.T) {
	tests := [...]struct {
		url         string
		wantPageLen int
		wantErr     error
	}{
		// TestNo  TestData
		0: {
			url:         "http://geo.groupkt.com/ip/172.217.3.14/json",
			wantPageLen: 371,
			wantErr:     nil,
		},
	}

	for i, td := range tests {
		page, err := GetWebPage(td.url)

		if td.wantErr != err {
			t.Errorf("Test %d failed, expected '%s', got '%s'", i, td.wantErr, err)
		}

		if td.wantPageLen != len(page) {
			t.Errorf("Test %d failed, expected '%d', got '%d'", i, td.wantPageLen, len(page))
		}
	}

}

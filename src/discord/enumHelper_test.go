package discord

import "testing"

func TestParseText(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"Single word lowercase", "pending", "PENDING"},
		{"Single word uppercase", "PENDING", "PENDING"},
		{"Single word mixed case", "PeNdInG", "PENDING"},
		{"Two words lowercase", "test word", "TEST_WORD"},
		{"Two words uppercase with space", "TEST WORD", "TEST_WORD"},
		{"Two words uppercase with underscore", "TEST_WORD", "TEST_WORD"},
		{"Two words mixed case with space", "TeSt WoRd", "TEST_WORD"},
		{"Two words mixed case with underscore", "TeSt_WoRd", "TEST_WORD"},
		{"Multiple words lowercase", "test multiple word", "TEST_MULTIPLE_WORD"},
		{"Multiple words uppercase with space", "TEST MULTIPLE WORD", "TEST_MULTIPLE_WORD"},
		{"Multiple words uppercase with underscore", "TEST_MULTIPLE_WORD", "TEST_MULTIPLE_WORD"},
		{"Multiple words mixed case with space", "TeSt MuLtIpLe WoRd", "TEST_MULTIPLE_WORD"},
		{"Multiple words mixed case with underscore", "TeSt_MuLtIpLe_WoRd", "TEST_MULTIPLE_WORD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseText(tt.input); got != tt.expected {
				t.Errorf("parseText(%s) = %s, expected %s", tt.input, got, tt.expected)
			}
		})
	}
}

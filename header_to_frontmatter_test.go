package main

import (
	"os"
	"testing"
)

func TestHasT2MDHeader(t *testing.T) {
	tests := []struct {
		filePath string
		want     bool
	}{
		{"test/has header simple.md", true},
	}

	for _, test := range tests {
		content, err := os.ReadFile(test.filePath)
		if err != nil {
			t.Fatalf("can't read test file %s. %s", test.filePath, err)
		}
		got := hasT2MDHeader(string(content))
		if got != test.want {
			t.Fatalf("\"%s\", expected: %v, got: %v", test.filePath, test.want, got)
		}
	}
}

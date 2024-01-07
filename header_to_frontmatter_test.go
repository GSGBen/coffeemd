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

	// test:
	//		file with it as expected
	//		file without it
	//		file with it but no "Original URL:"
	// 		file with the --- lower for another purpose, and "Original URL:" (false positive)
	//		a # further down

	for _, test := range tests {
		content, err := os.ReadFile(test.filePath)
		if err != nil {
			t.Fatalf("can't read test file %s. %s", test.filePath, err)
		}

		got, _ := findT2MDHeader(string(content))

		if got != test.want {
			t.Fatalf("\"%s\", expected: %v, got: %v", test.filePath, test.want, got)
		}
	}
}

package main

import (
	"os"
	"testing"
)

// test findT2MDHeader can find headers with no false positives.
func TestFindT2MDHeaderFinding(t *testing.T) {
	tests := []struct {
		filePath string
		want     bool
	}{
		{"test/has header simple.md", true},
		{"test/no header simple.md", false},
		{"test/has header and frontmatter.md", false},
		{"test/has header and duplicate bits 1.md", true},
		{"test/has almost correct bits 1.md", false},
		{"test/has almost correct bits 2.md", false},
	}

	// test:
	//		file with it but no "Original URL:"
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

// test findT2MDHeader doesn't extend the match too far.
func TestFindT2MDHeaderCorrectScope(t *testing.T) {
	tests := []struct {
		filePath string
		// path to the file containing the expected output
		wantPath string
	}{
		{
			"test/has header and duplicate bits 1.md",
			"test/has header and duplicate bits 1.md.expected",
		},
	}

	for _, test := range tests {
		content, err := os.ReadFile(test.filePath)
		if err != nil {
			t.Fatalf("can't read test file %s. %s", test.filePath, err)
		}
		want, err := os.ReadFile(test.wantPath)
		if err != nil {
			t.Fatalf("can't read expected output file %s. %s", test.wantPath, err)
		}

		_, got := findT2MDHeader(string(content))

		if got != string(want) {
			t.Fatalf("\"%s\", expected: %v, got: %v", test.filePath, string(want), got)
		}
	}
}

func TestParseHeader(t *testing.T) {
	headerString := `# 🥅 Full Card Name

Original URL: https://trello.com/example

---`

	want := parsedHeader{
		Title:       "🥅 Full Card Name",
		OriginalURL: "https://trello.com/example",
	}

	got, err := parseHeader(headerString)
	if err != nil {
		t.Fatalf("error parsing header string. %s", err)
	}
	if got.Title != want.Title {
		t.Fatalf("expected: %v, got: %v", want.Title, got.Title)
	}
	if got.OriginalURL != want.OriginalURL {
		t.Fatalf("expected: %v, got: %v", want.OriginalURL, got.OriginalURL)
	}
}

func TestParsedHeaderToYamlString(t *testing.T) {
	ph := parsedHeader{
		Title:       "🥅 Full Card Name",
		OriginalURL: "https://trello.com/example",
	}

	want := `---
title: 🥅 Full Card Name
original_url: https://trello.com/example
---
`

	got := ph.toYamlString()

	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

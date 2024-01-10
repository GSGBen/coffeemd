package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/urfave/cli"
)

// the regex used to find the header.
// e.g.
//
//	# (emoji) Full Card Name
//
//	Original URL: https://trello.com/example
//
//	---
var headerRegex = regexp.MustCompile(
	// no (?s) here because we want to explicitly control newline matching to narrow down the
	// search.
	// exactly at the start: #<space><any non-whitespace character><the rest of the title line>
	`^# \S.+` +
		// Original URL: in the exactly expected place (\r's for windows)
		`\r?\n\r?\n` +
		// and the exactly expected format
		`Original URL: https://trello\.com/.*` +
		// the line break in the exactly expected place
		`\r?\n\r?\n` +
		// and format
		`---\r?\n`)

// the regex that extracts the title and
var parsedHeaderRegex = regexp.MustCompile(
	// no (?s) here because we want to explicitly control newline matching to narrow down the
	// search.
	// the title is the rest of the line after the "# "
	`^# (.*?)\n` +
		// the original URL is the rest of the line after "Original URL: "
		`\nOriginal URL: (.*?)\n`)

// headerSearchResult keeps the file path, search result and (if relevant) the resulting
// header snippet together so that we don't have to run a regex search again to find the
// header when we need it.
type headerSearchResult struct {
	// the path to the markdown file
	filePath string
	// whether or not the search found the header
	containsHeader bool
	// if containsHeader is true, the content of the header (from the # to the end of
	// the ---)
	header string
}

// parsed header contains the title and Original URL extracted from a header.
type parsedHeader struct {
	Title       string `yaml:"title"`
	OriginalURL string `yaml:"original_url"`
}

// headerToFrontmatter is the entrypoint of the header-to-frontmatter action. it either shows the
// files it will change, or makes the change - converting the T2MD-created header to
// Obsidian/markdown frontmatter.
func headerToFrontmatter(cCtx *cli.Context, vaultPath string, apply bool) error {
	searchResults, err := findFilesWithT2MDHeader(vaultPath)
	if err != nil {
		return err
	}

	if apply {
		for _, hsr := range searchResults {
			convertHeaderInPlace(hsr)
			fmt.Printf("Converted the header in \"%s\"\n", hsr.filePath)
		}
		fmt.Println()
		fmt.Printf("Converted headers in %d files\n", len(searchResults))
		fmt.Println()
	} else {
		fmt.Println("The below files will have their headers (# Title, Original URL: value, --- separator) migrated to frontmatter:")
		fmt.Println()
		for _, hsr := range searchResults {
			fmt.Println(hsr.filePath)
		}
		fmt.Println()
		fmt.Printf("%d files total\n", len(searchResults))
		fmt.Println()
		fmt.Println("The above files will have their headers (# Title, Original URL: value, --- separator) migrated to frontmatter.")
		fmt.Println()
	}
	return nil
}

// findFilesWithHeader searches through all markdown files in vaultPath and returns
// those that have a header created by T2MD.
func findFilesWithT2MDHeader(vaultPath string) ([]headerSearchResult, error) {
	var results []headerSearchResult
	err := filepath.WalkDir(
		vaultPath,
		func(path string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() && strings.HasSuffix(path, ".md") {
				content, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				headerFound, header := findT2MDHeader(string(content))
				if headerFound {
					results = append(
						results,
						headerSearchResult{path, headerFound, header},
					)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// findT2MDHeader checks if the given markdown content starts with a header created by T2MD,
// e.g:
//
//	# (emoji) Full Card Name
//
//	Original URL: https://trello.com/example
//
//	---
//
//	(content...)
//
// If it does, `found` is true. If `found` is true, `header` will contain the content of the header.
//
// This is a key function because finding all the files with this and avoiding false positives is
// important. This one should have a lot of tests.
func findT2MDHeader(markdownContent string) (found bool, header string) {
	header = headerRegex.FindString(markdownContent)
	found = header != ""
	return
}

// convertHeaderInPlace converts the header in the given file to frontmatter.
func convertHeaderInPlace(hsr headerSearchResult) error {
	if !hsr.containsHeader {
		return errors.New("hsr.containsHeader == false")
	}

	ph, err := parseHeader(hsr.header)
	if err != nil {
		return err
	}

	yamlBytes, err := yaml.Marshal(&ph)
	if err != nil {
		return err
	}
	yamlString := "---\n" + string(yamlBytes) + "---\n"

	originalContent, err := os.ReadFile(hsr.filePath)
	if err != nil {
		return err
	}
	updatedContent := strings.Replace(string(originalContent), hsr.header, yamlString, 1)
	fileDetails, err := os.Stat(hsr.filePath)
	if err != nil {
		return err
	}
	err = os.WriteFile(hsr.filePath, []byte(updatedContent), fileDetails.Mode())
	if err != nil {
		return err
	}

	return nil
}

// parseHeader extracts the title and Original URL from a T2MD-generated header.
func parseHeader(header string) (parsedHeader, error) {
	submatches := parsedHeaderRegex.FindStringSubmatch(header)
	// standard setup - first index [0] is the full match, second [1] is the first submatch, etc
	if len(submatches) != 3 {
		return parsedHeader{}, errors.New("header and title not exactly matched")
	}

	return parsedHeader{
		Title:       submatches[1],
		OriginalURL: submatches[2],
	}, nil
}

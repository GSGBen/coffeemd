package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli"
)

// the regex used to find the # Title ... Original URL: ... --- header.
// e.g.
//
//	# (emoji) Full Card Name
//
//	Original URL: https://trello.com/example
//
//	---
var headerRegex = regexp.MustCompile(
	// no (?m) here because we want to explicitly control newline matching to narrow down the
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

// headerToFrontmatter is the entrypoint of the header-to-frontmatter action.
func headerToFrontmatter(cCtx *cli.Context, vaultPath string, apply bool) error {
	pathsOfFilesToChange, err := findFilesWithT2MDHeader(vaultPath)
	if err != nil {
		return err
	}

	if apply {
		// if check mode: output that the file will be changed
		// if not check mode:
		// 		convert the format
		//		change in place
		//		output that it was changed
	} else {
		fmt.Println("The below files will have their headers (# Title, Original URL: value, --- separator) migrated to frontmatter:")
		fmt.Println()
		fmt.Println(strings.Join(pathsOfFilesToChange, "\n"))
		fmt.Println()
		fmt.Printf("%d files total\n", len(pathsOfFilesToChange))
		fmt.Println()
		fmt.Println("The above files will have their headers (# Title, Original URL: value, --- separator) migrated to frontmatter.")
		fmt.Println()
	}
	return nil

	// test:
	//		file with it as expected
	//		file without it
	//		file with it but no "Original URL:"
	// 		file with the --- lower for another purpose, and "Original URL:" (false positive)
	//		a # further down
}

// findFilesWithHeader searches through all markdown files in vaultPath and returns the path of
// those that have a header created by T2MD.
func findFilesWithT2MDHeader(vaultPath string) ([]string, error) {
	var markdownFilePaths []string
	err := filepath.WalkDir(
		vaultPath,
		func(path string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() && strings.HasSuffix(path, ".md") {
				content, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				if hasT2MDHeader(string(content)) {
					markdownFilePaths = append(markdownFilePaths, path)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return markdownFilePaths, nil
}

// hasT2MDHeader returns true if the given markdown content starts with a header created by T2MD,
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
// This is a key function because finding all the files with this and avoiding false positives is
// important. This one should have a lot of tests.
func hasT2MDHeader(markdownContent string) bool {

	return headerRegex.MatchString(markdownContent)
}

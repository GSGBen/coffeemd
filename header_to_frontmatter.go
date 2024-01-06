package main

import (
	"fmt"
	"regexp"

	"github.com/urfave/cli"
)

func headerToFrontmatter(cCtx *cli.Context, vaultPath string, apply bool) error {
	fmt.Println(apply)
	return nil
}

// findFilesWithHeader searches through all markdown files in vaultPath and returns the path of
// those that have a header created by T2MD.
func findFilesWithT2MDHeader(vaultPath string) []string {
	return []string{}
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
	regex := regexp.MustCompile(
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

	return regex.MatchString(markdownContent)
}

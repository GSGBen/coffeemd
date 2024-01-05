package main

import (
	"fmt"

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
// important.
func hasT2MDHeader(markdownContent string) bool {
	return false
}

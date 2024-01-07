package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	vaultPath := ""
	apply := false

	app := &cli.App{
		Name:  "coffeemd",
		Usage: "a tool to make changes to files in an an Obsidian vault after you've created it with t2md and have already started using it.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "vault-path",
				Value:       "",
				Usage:       "the path to your Obsidian vault (or a folder of markdown files)",
				Required:    true,
				Destination: &vaultPath,
			},
			&cli.BoolFlag{
				Name:        "apply",
				Usage:       "Specify this flag to make the actual changes. Without this it'll just be a dry run",
				Destination: &apply,
			},
		},
		Commands: []cli.Command{
			{
				Name:  "header-to-frontmatter",
				Usage: `Convert the too-verbose header + Original URL: to yml frontmatter`,
				Description: `convert the manual and
too-heavy converted card headers to frontmatter. From:

	# (emoji) Full Card Name

	Original URL: https://trello.com/example

	---

	(content...)

to

	---
	title: (emoji) Full Card Name
	original_url: https://trello.com/example
	---

	(content...)
				`,
				Action: func(cCtx *cli.Context) error {
					return headerToFrontmatter(cCtx, vaultPath, apply)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

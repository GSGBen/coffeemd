package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	vaultPath := ""

	app := &cli.App{
		Name:  "coffeemd",
		Usage: "a tool to make changes to files in an an Obsidian vault after you've created it with t2md and have already started using it.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "vault-path",
				Value: "",
				Usage: "the path to your Obsidian vault (or a folder of markdown files)",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	// get the path to the obsidian vault (folder with markdown files) (global option)

	// take an apply option to actually make the change, not just list it (check mode by default)
	// (global option)

	// take a subcommand indicating which function to apply

	// find all files matching the issue

	// if check mode: output that the file will be changed
	// if not check mode:
	// 		convert the format
	//		change in place
	//		output that it was changed

	// test:
	//		file with it as expected
	//		file without it
	//		file with it but no "Original URL:"
	// 		file with the --- lower for another purpose, and "Original URL:" (false positive)
}

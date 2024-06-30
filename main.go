package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
	"github.com/jwalton/gchalk"
)

func main() {
	// Creates the 
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "filepath",
				Aliases: []string{"f"},
				Usage:   "Input the file path",
				Action:  readFile,
			},
			{
				Name:    "format",
				Aliases: []string{"ft"},
				Usage:   "Format type",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func readFile(cli_ctx *cli.Context) error {
	// Check if there's at least one argument after the command
	if cli_ctx.NArg() < 1 {
		return fmt.Errorf("file path not provided")
	}
	// filepath := cli_ctx.Args().Get(0) // Get the first argument as the file path
	// contents, err := os.ReadFile(filepath)

	filepaths := cli_ctx.Args().Slice()
	for i := 0; i < cli_ctx.Args().Len(); i++ {
		file := filepaths[i]

		contents, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
		return err
		}
	output := gchalk.WithRed().WithBgGreen().Bold(file) + "\n" + string(contents)
	PrintBoxedText(output)
	}

	return nil
}

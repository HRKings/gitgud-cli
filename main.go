package main

import (
	"github.com/HRKings/gitgud-cli/commit"
	"github.com/HRKings/gitgud-cli/flow"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "gitgud",
		Usage: "A cross-platform command line interface for the GitGud modular Git model",
		Commands: []*cli.Command{
			&commit.Command,
			&flow.Command,
		},
	}
	app.UseShortOptionHandling = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

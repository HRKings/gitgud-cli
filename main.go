package main

import (
	"github.com/HRKings/gitgud-cli/commit"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			&commit.Command,
		},
	}
	app.UseShortOptionHandling = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

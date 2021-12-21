package commit

import (
	"github.com/HRKings/gitgud-cli/utils"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var Command = cli.Command{
	Name:    "commit",
	Aliases: []string{"c"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "message",
			Aliases: []string{"m"},
			Usage:   "Use the given `<msg>` as the commit Subject.",
		},
		&cli.BoolFlag{
			Name:  "amend",
			Usage: "Replace the tip of the current branch by creating a new commit.",
		},
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.",
		},
		&cli.BoolFlag{
			Name:    "quick",
			Aliases: []string{"q"},
			Usage:   "Don't ask for missing parts",
		},
		&cli.StringFlag{
			Name:    "scope",
			Aliases: []string{"s"},
			Usage:   "`<scope>` is the area of the code that the commit affects.",
		},
	},
	Action: func(context *cli.Context) error {
		return ExecCommand(context, false)
	},
	Subcommands: []*cli.Command{
		{
			Name:    "full",
			Aliases: []string{"f"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "body",
					Aliases: []string{"b"},
					Usage:   "The `<body>` of this commit.",
				},
				&cli.StringFlag{
					Name:    "closes",
					Aliases: []string{"c"},
					Usage:   "`<#issue1, #issueN>` is a comma separated list of issues that this commit closes.",
				},
				&cli.StringFlag{
					Name:    "see",
					Aliases: []string{"r"},
					Usage:   "`<#issue1, #issueN>` is a comma separated list of issues that this commit references.",
				},
			},
			Action: func(context *cli.Context) error {
				return ExecCommand(context, true)
			},
		},
		{
			Name:    "generate",
			Aliases: []string{"gen"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "body",
					Aliases: []string{"b"},
					Usage:   "The `<body>` of this commit.",
				},
				&cli.StringFlag{
					Name:    "closes",
					Aliases: []string{"c"},
					Usage:   "`<#issue1, #issueN>` is a comma separated list of issues that this commit closes.",
				},
				&cli.StringFlag{
					Name:    "see",
					Aliases: []string{"r"},
					Usage:   "`<#issue1, #issueN>` is a comma separated list of issues that this commit references.",
				},
			},
			Action: func(context *cli.Context) error {
				message, err := BuildFullCommitMessage(context.String("message"),
					context.String("scope"), context.Bool("quick"), context.String("body"),
					context.String("closes"), context.String("see"))
				if err != nil {
					return err
				}

				color.Blue(message)
				return nil
			},
		},
	},
}

func ExecCommand(context *cli.Context, full bool) error {
	// Build the commit message
	var message string
	var err error

	// Build a full message or a simple one
	if !full {
		message, err = BuildCommitMessage(context.String("message"),
			context.String("scope"), context.Bool("quick"))
	} else {
		message, err = BuildFullCommitMessage(context.String("message"),
			context.String("scope"), context.Bool("quick"), context.String("body"),
			context.String("closes"), context.String("see"))
	}
	if err != nil {
		return err
	}

	// Execute the commit is everything is all right
	stdout, stderr, err := DoCommit(message, context.Bool("all"), context.Bool("amend"))

	// Print the stdout and stderr in case of an error
	if err != nil && stderr != "" {
		color.Red(stderr)
		return err
	} else if err != nil && stdout != "" {
		color.Red(stdout)
		return err
	}

	// Print the stdout of git
	color.Green(stdout)
	return nil
}

func DoCommit(message string, stageTracked bool, amend bool) (string, string, error) {
	var flags []string

	if stageTracked {
		flags = append(flags, "-a")
	}

	if amend {
		flags = append(flags, "--amend")
	}

	flags = append(flags, "-m", message)
	return utils.ExecGit("commit", flags...)
}

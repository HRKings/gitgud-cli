package commit

import "github.com/AlecAivazis/survey/v2"

func EnterSubject() (string, error) {
	subjectPrompt := &survey.Input{
		Message: "What is the commit subject?",
	}

	var subject string
	err := survey.AskOne(subjectPrompt, &subject, survey.WithValidator(survey.Required), survey.WithValidator(survey.MaxLength(80)))

	return subject, err
}

func EnterBody() (string, error) {
	bodyPrompt := &survey.Multiline{
		Message: "What is the commit body?",
	}

	var body string
	err := survey.AskOne(bodyPrompt, &body, survey.WithValidator(survey.MaxLength(100)))

	return body, err
}

func EnterCloses() (string, error) {
	closesPrompt := &survey.Input{
		Message: "What issues does this commits closes?",
	}

	var closes string
	err := survey.AskOne(closesPrompt, &closes)

	return closes, err
}

func EnterSeeAlso() (string, error) {
	alsoPrompt := &survey.Input{
		Message: "What issues does this commits references?",
	}

	var also string
	err := survey.AskOne(alsoPrompt, &also)

	return also, err
}

func EnterScope() (string, error) {
	scopePrompt := &survey.Input{
		Message: "What scope this commit affects?",
	}

	var scope string
	err := survey.AskOne(scopePrompt, &scope)

	return scope, err
}

func EnterTag() (string, error) {
	tagPrompt := &survey.Select{
		Message: "Select the commit tag:",
		Options: []string{
			"feature",
			"change",
			"fix",
			"style",
			"refactor",
			"test",
			"docs",
			"chore",
			"misc"},
		Default: "misc",
		Help: `Tag descriptions =
feature  : A new feature and small additions
change   : Any changes on existing functionality
fix      : A bugfix or hotfix
style    : Any change in styling, layout, css, design, etc
refactor : Any code refactoring, cleanup, formatting, improvements in code style and readability
test     : Adding, changing or refactoring tests, with no production code change
docs     : Changes in documentation, readme, guides, etc.
chore    : Updates/Removes/Adds dependencies, package manager configs, build tasks, etc.
misc     : Anything not covered by the above categories`,
		PageSize: 10,
	}

	var tag string
	err := survey.AskOne(tagPrompt, &tag, survey.WithValidator(survey.Required))

	return tag, err
}

func EnterFlags() ([]string, error) {
	flagsPrompt := &survey.MultiSelect{
		Message: "Select the flags for this commit:",
		Options: []string{
			"!!!",
			"db",
			"api",
			"ux",
			"dpc",
			"rm",
			"wip",
		},
		Help: `Flags descriptions =
!!! : Breaking change - Significant changes in software architecture and/or logic, that affects existing code.
db  : Changes that require database structure or data to be updated
api : Changes that modify the API usage, models or structure
ux  : Change in user experience - Anything that needs the user to relearn to use a feature
dpc : Deprecated - Commits with this flag deprecates existing code
rm  : Code Removal - Means that this commit removes old/legacy/deprecated code
wip : Work In Progress -  Commits marked as WIP can never be merged`,
		PageSize: 10,
	}

	var flags []string
	err := survey.AskOne(flagsPrompt, &flags)

	return flags, err
}

package flow_module

import "github.com/AlecAivazis/survey/v2"

func EnterBranchName() (string, error) {
	branchNamePrompt := &survey.Input{
		Message: "What is the branch name?",
	}

	var branchName string
	err := survey.AskOne(branchNamePrompt, &branchName, survey.WithValidator(survey.Required))

	return branchName, err
}

func EnterBranchType() (string, error) {
	branchTypePrompt := &survey.Select{
		Message: "Select the branch type:",
		Options: []string{
			"wip/",
			"fix/",
			"chore/",
			"hotfix/",
		},
		Default: "misc",
		Help: `Branch type descriptions =
wip/    : Receive direct commits and can be freely edited by the developers.
fix/    : Used to fix bugs or missing resources found in the master branch.
chore/  : Only used when updating dependencies, frameworks, build tasks and other updates required.
hotfix/ : A priority fix for when you find a serious bug in the stable branch that needs to be resolved ASAP.`,
	}

	var branchType string
	err := survey.AskOne(branchTypePrompt, &branchType, survey.WithValidator(survey.Required))

	return branchType, err
}
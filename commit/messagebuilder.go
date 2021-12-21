package commit

import (
	"fmt"
	"strings"
)

func BuildCommitMessage(subject string, domain string, quick bool) (string, error) {
	var err error

	// Ask for the subject if none is provided
	if subject == "" {
		subject, err = EnterSubject()
		if err != nil {
			return "", err
		}
	}

	// Ask for the flag if none is provided
	tag, err := EnterTag()
	if err != nil {
		return "", err
	}

	// Ask for the optional flags
	flags, err := EnterFlags()
	if err != nil {
		return "", err
	}

	// Ask for the domain if isn't in the quick mode
	if domain == "" && !quick {
		domain, err = EnterDomain()
		if err != nil {
			return "", err
		}
	}

	// Build the tag of the commit message
	message := fmt.Sprintf("[%s]", tag)

	// Add the flags if at least one is provided
	if len(flags) > 0 {
		message = fmt.Sprintf("%s{%s}", message, strings.Join(flags, "/"))
	}

	// Add the domain if one is provided
	if domain != "" {
		message = fmt.Sprintf("%s(%s)", message, domain)
	}

	// Add the subject and return the message
	return fmt.Sprintf("%s %s", message, subject), nil
}

func BuildFullCommitMessage(subject string, domain string, quick bool, body string, closes string, seeAlso string) (string, error) {
	// Build the base message
	baseMessage, err := BuildCommitMessage(subject, domain, quick)
	if err != nil {
		return "", err
	}

	// Ask for the body if one is not provided, and it isn't in the quick mode
	if body == "" && !quick {
		body, err = EnterBody()
		if err != nil {
			return "", err
		}
	}

	// Ask for the closed issues if one is not provided, and it isn't in the quick mode
	if closes == "" && !quick {
		closes, err = EnterCloses()
		if err != nil {
			return "", err
		}
	}

	// Ask for the referenced issues if one is not provided, and it isn't in the quick mode
	if seeAlso == "" && !quick {
		seeAlso, err = EnterSeeAlso()
		if err != nil {
			return "", err
		}
	}

	// Add the body to the complete message if one is provided
	if body != "" {
		baseMessage = fmt.Sprintf("%s\n~~~\n%s", baseMessage, body)
	}

	// Format the closed issues if any is provided
	if closes != "" {
		closes = fmt.Sprintf("Closes: %s", closes)
	}

	// Format the referenced issues if any is provided
	if seeAlso != "" {
		seeAlso = fmt.Sprintf("See also: %s", seeAlso)
	}

	// Concat and format the footer to its final form
	// if the footer is not empty, then put it on the complete message
	footer := strings.Trim(strings.Join([]string{closes, seeAlso}, "\n"), "\n")
	if footer != "" {
		baseMessage = fmt.Sprintf("%s\n~~~\n%s", baseMessage, footer)
	}

	// Return the complete message and no errors
	return baseMessage, nil
}

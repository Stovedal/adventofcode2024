package utils

import (
	"os"
	"strconv"
)

func ParseUserArguments(existingSolutions int) (int, string) {
	if len(os.Args) < 3 {
		panic("Arguments missing, provide solutionIdentifier to solve and session token")
	}

	solutionIdentifier, err := strconv.Atoi(os.Args[1])

	if err != nil || solutionIdentifier < 1 || solutionIdentifier > existingSolutions {
		panic("Provided solutionIdentifier argument is invalid: %v, solutionIdentifiers up to %v exists\n")
	}

	sessionToken := os.Args[2]

	if sessionToken == "" {
		panic("Session token is invalid")
	}

	return solutionIdentifier, sessionToken
}

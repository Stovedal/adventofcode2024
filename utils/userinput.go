package utils

import (
	"os"
	"strconv"
)

func ParseUserArguments(existingSolutions int) (int, string) {
	if len(os.Args) < 3 {
		panic("Arguments missing, provide solution identifier and session token")
	}

	solutionIdentifier, err := strconv.Atoi(os.Args[1])

	if err != nil || solutionIdentifier < 1 || solutionIdentifier > existingSolutions {
		panic("Provided solution identifier argument is out of range")
	}

	sessionToken := os.Args[2]

	if sessionToken == "" {
		panic("Session token is invalid")
	}

	return solutionIdentifier, sessionToken
}

package input

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func FetchInputForSolution(solutionIdentifier int, sessionToken string) string {
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", solutionIdentifier)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%v", sessionToken))

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	input := string(body)

	return input
}

package utils

import (
	"fmt"
	"strings"
)

// PromptRepoSelection processes a comma-separated repository list and returns unique repo names.
func PromptRepoSelection(reposInput []string) ([]string, error) {
	if len(reposInput) == 0 {
		return nil, fmt.Errorf("no repositories provided")
	}

	repoSet := make(map[string]bool)
	var uniqueRepos []string
	for _, r := range reposInput {
		trimmed := strings.TrimSpace(r)
		if trimmed != "" && !repoSet[trimmed] {
			repoSet[trimmed] = true
			uniqueRepos = append(uniqueRepos, trimmed)
		}
	}
	fmt.Println("Selected repositories:", uniqueRepos)
	return uniqueRepos, nil
}


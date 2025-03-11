// utils/prompt.go - Utility function for repository selection
package utils

import (
	"strings"
)

// PromptRepoSelection ensures unique and valid repositories
func PromptRepoSelection(reposInput []string) ([]string, error) {
	if len(reposInput) == 0 {
		return nil, nil
	}

	// Deduplicate repositories
	repoSet := make(map[string]bool)
	var uniqueRepos []string
	for _, r := range reposInput {
		trimmed := strings.TrimSpace(r)
		if trimmed != "" && !repoSet[trimmed] {
			repoSet[trimmed] = true
			uniqueRepos = append(uniqueRepos, trimmed)
		}
	}

	return uniqueRepos, nil
}


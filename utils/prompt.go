// utils/prompt.go - Interactive selection for repositories

package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PromptRepoSelection prompts the user to select repositories interactively
func PromptRepoSelection(repos []string) ([]string, error) {
	fmt.Println("Select repositories to create a branch (comma-separated indices):")

	// Display repository list
	for i, repo := range repos {
		fmt.Printf("[%d] %s\n", i+1, repo)
	}

	// Get user input
	fmt.Print("Enter the numbers of the repositories you want (e.g., 1,3,5): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// Process input
	input = strings.TrimSpace(input)
	indices := strings.Split(input, ",")

	var selectedRepos []string
	for _, index := range indices {
		idx := strings.TrimSpace(index)
		i, err := strconv.Atoi(idx)
		if err != nil || i < 1 || i > len(repos) {
			fmt.Printf("⚠️ Invalid input: %s (Skipping)\n", idx)
			continue
		}
		selectedRepos = append(selectedRepos, repos[i-1])
	}

	if len(selectedRepos) == 0 {
		return nil, fmt.Errorf("no repositories selected")
	}

	return selectedRepos, nil
}

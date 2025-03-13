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
// (This is used for branch creation.)
func PromptRepoSelection(repos []string) ([]string, error) {
	fmt.Println("Select repositories to create a branch (comma-separated indices):")
	for i, repo := range repos {
		fmt.Printf("[%d] %s\n", i+1, repo)
	}
	fmt.Print("Enter the numbers of the repositories you want (e.g., 1,3,5): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	indices := strings.Split(input, ",")
	var selectedRepos []string
	for _, index := range indices {
		idx, err := strconv.Atoi(strings.TrimSpace(index))
		if err != nil {
			fmt.Printf("⚠️ Invalid input: %s (Skipping)\n", index)
			continue
		}
		if idx >= 1 && idx <= len(repos) {
			selectedRepos = append(selectedRepos, repos[idx-1])
		} else {
			fmt.Printf("⚠️ Index out of range: %d (Skipping)\n", idx)
		}
	}
	if len(selectedRepos) == 0 {
		return nil, fmt.Errorf("no repositories selected")
	}
	return selectedRepos, nil
}

// PromptRepoSelectionForPR prompts the user to select repositories for pull request creation.
func PromptRepoSelectionForPR(repos []string) ([]string, error) {
	fmt.Println("Select repositories to create a pull request (comma-separated indices):")
	for i, repo := range repos {
		fmt.Printf("[%d] %s\n", i+1, repo)
	}
	fmt.Print("Enter the numbers of the repositories you want (e.g., 1,3,5): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	indices := strings.Split(input, ",")
	var selectedRepos []string
	for _, index := range indices {
		idx, err := strconv.Atoi(strings.TrimSpace(index))
		if err != nil {
			fmt.Printf("⚠️ Invalid input: %s (Skipping)\n", index)
			continue
		}
		if idx >= 1 && idx <= len(repos) {
			selectedRepos = append(selectedRepos, repos[idx-1])
		} else {
			fmt.Printf("⚠️ Index out of range: %d (Skipping)\n", idx)
		}
	}
	if len(selectedRepos) == 0 {
		return nil, fmt.Errorf("no repositories selected")
	}
	return selectedRepos, nil
}

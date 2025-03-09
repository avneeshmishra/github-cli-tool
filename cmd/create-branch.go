package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/avneeshmishra/go-github-cli/github"
	"github.com/avneeshmishra/go-github-cli/utils"
	"github.com/spf13/cobra"
)

var createBranchCmd = &cobra.Command{
	Use:   "create-branch",
	Short: "Create a new branch in selected repositories with rollback on error",
	Run: func(cmd *cobra.Command, args []string) {
		// Get GitHub credentials from environment variables.
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			log.Fatal("GITHUB_TOKEN environment variable is not set.")
		}
		owner := os.Getenv("GITHUB_OWNER")
		if owner == "" {
			log.Fatal("GITHUB_OWNER environment variable is not set.")
		}

		// Initialize the GitHub client.
		client := github.NewGitHubClient(token, owner)

		// Deduplicate the repo list.
		repoSet := make(map[string]bool)
		var repos []string
		for _, r := range repoNames {
			trimmed := strings.TrimSpace(r)
			if trimmed != "" && !repoSet[trimmed] {
				repoSet[trimmed] = true
				repos = append(repos, trimmed)
			}
		}

		// Optionally allow interactive selection.
		selectedRepos, err := utils.PromptRepoSelection(repos)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Selected repositories:", selectedRepos)

		// Track which repositories have been processed successfully.
		var createdRepos []string

		// Process each repository.
		for _, repo := range selectedRepos {
			err := client.CreateBranch(repo, branchName, baseBranch)
			if err != nil {
				log.Printf("Error creating branch in %s: %v\n", repo, err)
				if rollback && len(createdRepos) > 0 {
					log.Println("Rollback enabled: Deleting branches created so far...")
					for _, r := range createdRepos {
						delErr := client.DeleteBranch(r, branchName)
						if delErr != nil {
							log.Printf("Error rolling back branch in %s: %v\n", r, delErr)
						} else {
							fmt.Printf("Rolled back branch '%s' in %s\n", branchName, r)
						}
					}
					// Exit after rollback.
					return
				}
			} else {
				// Only add repository if branch creation succeeded and no error occurred.
				createdRepos = append(createdRepos, repo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createBranchCmd)
}


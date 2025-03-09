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

var createPRCmd = &cobra.Command{
	Use:   "create-pr",
	Short: "Create a pull request in selected repositories",
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

		// Initialize GitHub client.
		client := github.NewGitHubClient(token, owner)

		// Deduplicate repositories.
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

		// Create a PR in each selected repository.
		for _, repo := range selectedRepos {
			err := client.CreatePullRequest(repo, prTitle, prBody, branchName, baseBranch)
			if err != nil {
				log.Printf("Error creating PR in %s: %v\n", repo, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createPRCmd)
}


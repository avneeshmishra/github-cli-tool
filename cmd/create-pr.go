// cmd/create-pr.go - Handles pull request creation with rollback support

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
		// Fetch GitHub credentials from environment variables
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			log.Fatal("GITHUB_TOKEN is not set. Please set it before running this command.")
		}
		owner := os.Getenv("GITHUB_OWNER")
		if owner == "" {
			log.Fatal("GITHUB_OWNER is not set. Please set it before running this command.")
		}

		// Initialize GitHub client
		client := github.NewGitHubClient(token, owner)

		// Deduplicate repo list
		repoSet := make(map[string]bool)
		var repos []string
		for _, r := range repoNames {
			trimmed := strings.TrimSpace(r)
			if trimmed != "" && !repoSet[trimmed] {
				repoSet[trimmed] = true
				repos = append(repos, trimmed)
			}
		}

		// Let user confirm repo selection
		selectedRepos, err := utils.PromptRepoSelection(repos)
		if err != nil {
			log.Fatal(err)
		}

		// Print selected repositories only once
		fmt.Printf("Selected repositories: %v\n", selectedRepos)

		// ✅ Track created PRs for rollback if needed
		var createdPRs []struct {
			Repo     string
			PRNumber int
		}

		// ✅ Process each repository
		for _, repo := range selectedRepos {
			// Create PR and capture PR number
			prNumber, err := client.CreatePullRequest(repo, prTitle, prBody, branchName, baseBranch)
			if err != nil {
				log.Printf("Error creating PR in %s: %v\n", repo, err)

				// ✅ If rollback is enabled, delete created PRs
				if rollback && len(createdPRs) > 0 {
					log.Println("Rolling back created PRs...")
					for _, pr := range createdPRs {
						delErr := client.DeletePullRequest(pr.Repo, pr.PRNumber)
						if delErr != nil {
							log.Printf("Error rolling back PR in %s: %v\n", pr.Repo, delErr)
						} else {
							fmt.Printf("✅ Rolled back PR #%d in %s\n", pr.PRNumber, pr.Repo)
						}
					}
					return
				}
			} else {
				// ✅ Store created PR info for rollback
				createdPRs = append(createdPRs, struct {
					Repo     string
					PRNumber int
				}{Repo: repo, PRNumber: prNumber})
				fmt.Printf("✅ Pull request #%d created successfully in %s\n", prNumber, repo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createPRCmd)
}

// cmd/create-pr.go - Handles pull request creation with rollback support
// Author: Avneesh Mishra

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
		// Ensure required environment variables are set.
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			log.Fatal("GITHUB_TOKEN is not set.")
		}
		owner := os.Getenv("GITHUB_OWNER")
		if owner == "" {
			log.Fatal("GITHUB_OWNER is not set.")
		}

		client := github.NewGitHubClient(token, owner)
		var repos []string

		// Build the provided repository list after trimming spaces and quotes.
		providedRepos := make([]string, 0)
		for _, r := range repoNames {
			trimmed := strings.Trim(strings.TrimSpace(r), "\"")
			if trimmed != "" {
				providedRepos = append(providedRepos, trimmed)
			}
		}

		// Use provided repo list if available; otherwise, prompt interactively.
		if len(providedRepos) > 0 {
			repos = providedRepos
		} else {
			fmt.Println("Fetching available repositories from GitHub...")
			allRepos, err := client.ListRepositories()
			if err != nil {
				log.Fatalf("Error fetching repositories: %v", err)
			}
			if len(allRepos) == 0 {
				log.Fatal("No repositories found in your GitHub account.")
			}
			// Use the PR-specific prompt message.
			selectedRepos, err := utils.PromptRepoSelectionForPR(allRepos)
			if err != nil {
				log.Fatal(err)
			}
			repos = selectedRepos
		}

		if len(repos) == 0 {
			log.Fatal("No valid repositories selected.")
		}

		fmt.Printf("Selected repositories: %v\n", repos)

		// Track created PRs for rollback if needed.
		var createdPRs []struct {
			Repo     string
			PRNumber int
		}

		for _, repo := range repos {
			prNumber, err := client.CreatePullRequest(repo, prTitle, prBody, branchName, baseBranch)
			if err != nil {
				log.Printf("Error creating PR in %s: %v", repo, err)
				if rollback && len(createdPRs) > 0 {
					log.Println("Rolling back created PRs...")
					for _, pr := range createdPRs {
						client.DeletePullRequest(pr.Repo, pr.PRNumber)
					}
					return
				}
			} else {
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

// cmd/create-branch.go - Handles branch creation with rollback
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
		// Fetch API credentials from env vars
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			log.Fatal("GITHUB_TOKEN is not set. Please set it before running this command.")
		}
		owner := os.Getenv("GITHUB_OWNER")
		if owner == "" {
			log.Fatal("GITHUB_OWNER is not set. Please set it before running this command.")
		}

		// Initialize GitHub API client
		client := github.NewGitHubClient(token, owner)

		// Remove duplicates from repo list
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

		// Only print repo selection once
		fmt.Printf("Selected repositories: %v\n", selectedRepos)

		// Store created branches for rollback if needed
		var createdRepos []string

		// Process each repo
		for _, repo := range selectedRepos {
			err := client.CreateBranch(repo, branchName, baseBranch)
			if err != nil {
				log.Printf("Error creating branch in %s: %v\n", repo, err)

				// If rollback flag is set, delete already created branches
				if rollback && len(createdRepos) > 0 {
					log.Println("Rolling back created branches...")
					for _, r := range createdRepos {
						delErr := client.DeleteBranch(r, branchName)
						if delErr != nil {
							log.Printf("Error rolling back branch in %s: %v\n", r, delErr)
						} else {
							fmt.Printf("✅ Rolled back branch '%s' in %s\n", branchName, r)
						}
					}
					return
				}
			} else {
				// Add to rollback list
				createdRepos = append(createdRepos, repo)
				fmt.Printf("✅ Branch '%s' created successfully in %s\n", branchName, repo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createBranchCmd)
}


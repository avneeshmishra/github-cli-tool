// cmd/create-branch.go - Handles branch creation with interactive repo selection

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
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			log.Fatal("GITHUB_TOKEN is not set.")
		}
		owner := os.Getenv("GITHUB_OWNER")
		if owner == "" {
			log.Fatal("GITHUB_OWNER is not set.")
		}

		client := github.NewGitHubClient(token, owner)

		// ✅ Step 1: Fetch repositories from GitHub
		fmt.Println("Fetching available repositories from GitHub...")
		allRepos, err := client.ListRepositories()
		if err != nil {
			log.Fatalf("Error fetching repositories: %v", err)
		}
		if len(allRepos) == 0 {
			log.Fatal("No repositories found in your GitHub account.")
		}

		// ✅ Step 2: Prompt the user to select repositories
		selectedRepos, err := utils.PromptRepoSelection(allRepos)
		if err != nil {
			log.Fatal(err)
		}
		repoNames = selectedRepos

		// ✅ Step 3: Deduplicate and clean repo list
		repoSet := make(map[string]bool)
		var repos []string
		for _, r := range repoNames {
			trimmed := strings.TrimSpace(r)
			if trimmed != "" && !repoSet[trimmed] {
				repoSet[trimmed] = true
				repos = append(repos, trimmed)
			}
		}

		if len(repos) == 0 {
			log.Fatal("No valid repositories selected.")
		}

		fmt.Printf("✅ Selected repositories: %v\n", repos)

		// ✅ Step 4: Create branch in selected repositories
		var createdRepos []string

		for _, repo := range repos {
			// 🔹 Fetch the correct SHA for the base branch
			baseBranchSHA, err := client.GetBranchSHA(repo, baseBranch)
			if err != nil {
				log.Printf("⏩ Skipping %s: Failed to get base branch SHA: %v\n", repo, err)
				continue
			}

			err = client.CreateBranch(repo, branchName, baseBranchSHA)
			if err != nil {
				log.Printf("❌ Error creating branch in %s: %v\n", repo, err)

				// 🔹 Rollback if enabled
				if rollback && len(createdRepos) > 0 {
					log.Println("⚠️ Rolling back created branches...")
					for _, r := range createdRepos {
						client.DeleteBranch(r, branchName) // Handles rollback
					}
					return
				}
			} else {
				createdRepos = append(createdRepos, repo)
				fmt.Printf("✅ Branch '%s' created successfully in %s\n", branchName, repo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createBranchCmd)
}

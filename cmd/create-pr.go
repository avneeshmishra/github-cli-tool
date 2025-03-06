package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/avneeshmishra/go-github-cli/github"
)

// CreatePrCmd represents the create-pr command
var CreatePrCmd = &cobra.Command{
	Use:   "create-pr",
	Short: "Create a pull request from a branch to the base branch in one or more GitHub repositories",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the token from flags
		token, _ := cmd.Flags().GetString("token")

		// Get the list of repositories, PR title, and PR body from the flags
		repoNames, _ := cmd.Flags().GetString("repo")
		prTitle, _ := cmd.Flags().GetString("title")
		prBody, _ := cmd.Flags().GetString("body")
		prBranch, _ := cmd.Flags().GetString("branch")

		// Split the repo names by comma
		repos := strings.Split(repoNames, ",")

		// Create a new GitHub client
		client, err := github.NewGitHubClient(token)
		if err != nil {
			fmt.Println("Error initializing GitHub client:", err)
			return
		}

		// Loop through each repository and create the pull request
		for _, repo := range repos {
			err := client.CreatePR(repo, prBranch, prTitle, prBody)
			if err != nil {
				fmt.Printf("Error creating pull request in %s: %v\n", repo, err)
			} else {
				fmt.Printf("Pull request created successfully in %s\n", repo)
			}
		}
	},
}

func init() {
	// Define flags for create-pr command
	CreatePrCmd.Flags().String("repo", "", "Comma-separated list of repositories (e.g., user/repo1,user/repo2)")
	CreatePrCmd.Flags().String("title", "", "The title of the pull request")
	CreatePrCmd.Flags().String("body", "", "The body of the pull request")
	CreatePrCmd.Flags().String("branch", "", "The name of the branch to create the pull request from")
	CreatePrCmd.MarkFlagRequired("repo")
	CreatePrCmd.MarkFlagRequired("title")
	CreatePrCmd.MarkFlagRequired("body")
	CreatePrCmd.MarkFlagRequired("branch")
}


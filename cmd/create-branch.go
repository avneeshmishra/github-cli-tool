package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/avneeshmishra/go-github-cli/github"
)

// CreateBranchCmd represents the create-branch command
var CreateBranchCmd = &cobra.Command{
	Use:   "create-branch",
	Short: "Create a new branch in one or more GitHub repositories",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the token from flags
		token, _ := cmd.Flags().GetString("token")

		// Get the list of repositories and branch name from the flags
		repoNames, _ := cmd.Flags().GetString("repo")
		branchName, _ := cmd.Flags().GetString("branch")
		baseBranch, _ := cmd.Flags().GetString("base")

		// Split the repo names by comma
		repos := strings.Split(repoNames, ",")

		// Create a new GitHub client
		client, err := github.NewGitHubClient(token)
		if err != nil {
			fmt.Println("Error initializing GitHub client:", err)
			return
		}

		// Loop through each repository and create the branch
		for _, repo := range repos {
			err := client.CreateBranch(repo, branchName, baseBranch)
			if err != nil {
				fmt.Printf("Error creating branch in %s: %v\n", repo, err)
			} else {
				fmt.Printf("Branch '%s' created successfully in %s\n", branchName, repo)
			}
		}
	},
}

func init() {
	// Define flags for create-branch command
	CreateBranchCmd.Flags().String("repo", "", "Comma-separated list of repositories (e.g., user/repo1,user/repo2)")
	CreateBranchCmd.Flags().String("branch", "", "The name of the new branch")
	CreateBranchCmd.Flags().String("base", "main", "The base branch to create the new branch from")
	CreateBranchCmd.MarkFlagRequired("repo")
	CreateBranchCmd.MarkFlagRequired("branch")
}


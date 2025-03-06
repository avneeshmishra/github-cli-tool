package cmd

import (
	"fmt"
	"github.com/avneeshmishra/go-github-cli/github"
	"github.com/spf13/cobra"
	"os"
)

// createBranchCmd represents the create-branch command
var createBranchCmd = &cobra.Command{
	Use:   "create-branch",
	Short: "Creates a new branch in the specified repository",
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			fmt.Println("Error: GITHUB_TOKEN environment variable is required")
			return
		}

		client := github.NewGitHubClient(token)

		if err := client.CreateBranch(repo, branchName); err != nil {
			fmt.Println("Error creating branch:", err)
			return
		}

		fmt.Println("Branch created successfully.")
	},
}

var repo string
var branchName string

func init() {
	createBranchCmd.Flags().StringVarP(&repo, "repo", "r", "", "GitHub repository (owner/repo)")
	createBranchCmd.Flags().StringVarP(&branchName, "branch", "b", "", "Branch name")

	createBranchCmd.MarkFlagRequired("repo")
	createBranchCmd.MarkFlagRequired("branch")

	rootCmd.AddCommand(createBranchCmd)
}


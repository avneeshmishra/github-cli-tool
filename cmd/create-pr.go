package cmd

import (
	"fmt"
	"github.com/avneeshmishra/go-github-cli/github"
	"github.com/spf13/cobra"
	"os"
)

// createPRCmd represents the create-pr command
var createPRCmd = &cobra.Command{
	Use:   "create-pr",
	Short: "Creates a pull request for a new branch",
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			fmt.Println("Error: GITHUB_TOKEN environment variable is required")
			return
		}

		client := github.NewGitHubClient(token)

		if err := client.CreatePR(repo, branchName, prTitle, prBody); err != nil {
			fmt.Println("Error creating pull request:", err)
			return
		}

		fmt.Println("Pull request created successfully.")
	},
}

var prTitle string
var prBody string

func init() {
	createPRCmd.Flags().StringVarP(&repo, "repo", "r", "", "GitHub repository (owner/repo)")
	createPRCmd.Flags().StringVarP(&branchName, "branch", "b", "", "Branch name")
	createPRCmd.Flags().StringVarP(&prTitle, "title", "t", "", "Pull request title")
	createPRCmd.Flags().StringVarP(&prBody, "body", "d", "", "Pull request body")

	createPRCmd.MarkFlagRequired("repo")
	createPRCmd.MarkFlagRequired("branch")
	createPRCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(createPRCmd)
}


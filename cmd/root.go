// cmd/root.go - Defines the root CLI command and global flags
// Author: Avneesh Mishra

package cmd

import "github.com/spf13/cobra"

// Global flags - these are available for all commands
var (
	repoNames  []string // stores repo names passed in CLI
	branchName string   // branch to create
	baseBranch string   // base branch (default: main)
	prTitle    string   // PR title
	prBody     string   // PR description
	rollback   bool     // flag to rollback on failure
)

// rootCmd is the main command, everything starts from here
var rootCmd = &cobra.Command{
	Use:   "go-github-cli",
	Short: "CLI tool to manage GitHub branches and pull requests",
}

// Execute runs the CLI
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Defining flags so users can pass values from CLI
	rootCmd.PersistentFlags().StringSliceVarP(&repoNames, "repo", "r", nil, "Comma-separated list of repositories")
	rootCmd.PersistentFlags().StringVarP(&branchName, "branch", "b", "", "Branch name to create")
	rootCmd.PersistentFlags().StringVarP(&baseBranch, "base", "B", "main", "Base branch name")
	rootCmd.PersistentFlags().StringVarP(&prTitle, "title", "t", "", "Pull request title")
	rootCmd.PersistentFlags().StringVarP(&prBody, "body", "d", "", "Pull request description")
	rootCmd.PersistentFlags().BoolVarP(&rollback, "rollback", "R", false, "Enable rollback on failure")
}


// cmd/root.go - Defines the root CLI command and global flags
// Author: Avneesh Mishra

package cmd

import "github.com/spf13/cobra"

var (
	// Global flag variables
	repoNames  []string // repository list provided via flag
	branchName string   // branch name to create / use for PR
	baseBranch string   // base branch (default "main")
	prTitle    string   // pull request title
	prBody     string   // pull request description
	rollback   bool     // rollback flag on failure
)

var rootCmd = &cobra.Command{
	Use:   "go-github-cli",
	Short: "CLI tool to manage GitHub branches and pull requests",
	// When no subcommand is provided, show help.
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute runs the CLI
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Define persistent flags with a default empty slice for repoNames.
	rootCmd.PersistentFlags().StringSliceVarP(&repoNames, "repo", "r", []string{}, "Comma-separated list of repositories")
	rootCmd.PersistentFlags().StringVarP(&branchName, "branch", "b", "", "Branch name to create")
	rootCmd.PersistentFlags().StringVarP(&baseBranch, "base", "B", "main", "Base branch name")
	rootCmd.PersistentFlags().StringVarP(&prTitle, "title", "t", "", "Pull request title")
	rootCmd.PersistentFlags().StringVarP(&prBody, "body", "d", "", "Pull request description")
	rootCmd.PersistentFlags().BoolVarP(&rollback, "rollback", "R", false, "Enable rollback on failure")
}

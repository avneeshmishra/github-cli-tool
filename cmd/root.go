package cmd

import "github.com/spf13/cobra"

// Global variables for flags (shared across commands)
var (
	repoNames  []string // comma-separated list of repositories (e.g. "avi,gosample")
	branchName string   // branch name to create / use for PR
	baseBranch string   // base branch (default "main")
	prTitle    string   // pull request title
	prBody     string   // pull request description
	rollback   bool     // if true, rollback (delete) branches created earlier on error
)

// rootCmd is the base command for the CLI.
var rootCmd = &cobra.Command{
	Use:   "go-github-cli",
	Short: "CLI tool to manage GitHub branches and pull requests",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Define global persistent flags.
	rootCmd.PersistentFlags().StringSliceVarP(&repoNames, "repo", "r", nil, "Comma-separated list of repositories (format: owner/repo)")
	rootCmd.PersistentFlags().StringVarP(&branchName, "branch", "b", "", "Branch name to create (and for PRs)")
	rootCmd.PersistentFlags().StringVarP(&baseBranch, "base", "B", "main", "Base branch name")
	rootCmd.PersistentFlags().StringVarP(&prTitle, "title", "t", "", "Pull request title")
	rootCmd.PersistentFlags().StringVarP(&prBody, "body", "d", "", "Pull request description")
	rootCmd.PersistentFlags().BoolVarP(&rollback, "rollback", "R", false, "Enable rollback on error (delete previously created branches)")
}


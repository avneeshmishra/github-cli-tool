package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var githubToken string

var rootCmd = &cobra.Command{
	Use:   "go-github-cli",
	Short: "CLI tool for managing GitHub branches and PRs across multiple repositories",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	githubToken = os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		fmt.Println("Error: GITHUB_TOKEN environment variable is not set.")
		os.Exit(1)
	}
}

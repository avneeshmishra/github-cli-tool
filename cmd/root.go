package cmd

import (
	"github.com/spf13/cobra"
)

var githubToken string

// Root command
var RootCmd = &cobra.Command{
	Use:   "go-github-cli",
	Short: "A CLI tool for managing GitHub repositories",
}

func init() {
	// Persistent flag for GitHub token so that it's available in all subcommands
	RootCmd.PersistentFlags().StringVar(&githubToken, "token", "", "GitHub personal access token")
}


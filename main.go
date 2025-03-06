/*
Copyright © 2025 Github @avneeshmishra
*/

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/avneeshmishra/go-github-cli/cmd"
)

func main() {
	// Create the root command for the CLI tool
	var rootCmd = &cobra.Command{
		Use:   "go-github-cli",
		Short: "A CLI tool for managing GitHub repositories",
	}

	// Declare the token flag to be used across commands
	var token string

	// Define the token flag for the root command
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "GitHub Personal Access Token (PAT)")
	rootCmd.MarkPersistentFlagRequired("token")

	// Add the "create-branch" command to the root command
	rootCmd.AddCommand(cmd.CreateBranchCmd)
	// Add the "create-pr" command to the root command
	rootCmd.AddCommand(cmd.CreatePrCmd)

	// Pass the token to the CreateBranch and CreatePr commands
	cmd.CreateBranchCmd.PersistentFlags().StringVar(&token, "token", "", "GitHub Personal Access Token")
	cmd.CreatePrCmd.PersistentFlags().StringVar(&token, "token", "", "GitHub Personal Access Token")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		// Print the error and exit with status 1 if the command fails
		fmt.Println(err)
		os.Exit(1)
	}
}


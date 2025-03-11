// main.go - Entry point of the application
// Author: Avneesh Mishra

package main

import (
	"log"
	"os"

	"github.com/avneeshmishra/go-github-cli/cmd"
)

func main() {
	// Running the CLI command execution
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}


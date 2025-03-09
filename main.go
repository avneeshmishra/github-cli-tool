package main

import (
	"log"
	"os"

	"github.com/avneeshmishra/go-github-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}


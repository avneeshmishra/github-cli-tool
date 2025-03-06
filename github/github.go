package github

import (
	"context" // Import the context package
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GitHubClient is a custom wrapper around the GitHub client
type GitHubClient struct {
	client *github.Client
}

// NewGitHubClient initializes a new GitHub client using the provided token
func NewGitHubClient(token string) (*GitHubClient, error) {
	if token == "" {
		return nil, fmt.Errorf("GitHub token is required")
	}

	// Use OAuth2 to authenticate with the GitHub API
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(nil, ts)
	client := github.NewClient(tc)

	return &GitHubClient{client: client}, nil
}

// CreateBranch creates a new branch in the provided repository
func (c *GitHubClient) CreateBranch(repo string, branchName string, baseBranch string) error {
	// Get the latest commit on the base branch
	ctx := context.Background() // Create a background context
	commit, _, err := c.client.Repositories.GetCommit(ctx, strings.Split(repo, "/")[0], strings.Split(repo, "/")[1], baseBranch)
	if err != nil {
		return fmt.Errorf("failed to get commit for base branch: %v", err)
	}

	// Create the new branch by referencing the base branch's latest commit
	ref := &github.Reference{
		Object: &github.GitObject{
			SHA: commit.SHA,
		},
	}

	_, _, err = c.client.Git.CreateRef(ctx, strings.Split(repo, "/")[0], strings.Split(repo, "/")[1], &github.Reference{
		Ref:    github.String("refs/heads/" + branchName),
		Object: ref.Object,
	})
	if err != nil {
		return fmt.Errorf("failed to create branch: %v", err)
	}

	return nil
}

// CreatePR creates a new pull request from the given branch to the base branch
func (c *GitHubClient) CreatePR(repo string, prBranch string, prTitle string, prBody string) error {
	ctx := context.Background() // Create a background context
	// Create the pull request
	_, _, err := c.client.PullRequests.Create(ctx, strings.Split(repo, "/")[0], strings.Split(repo, "/")[1], &github.NewPullRequest{
		Title: github.String(prTitle),
		Head:  github.String(prBranch),
		Base:  github.String("main"),
		Body:  github.String(prBody),
	})
	if err != nil {
		return fmt.Errorf("failed to create pull request: %v", err)
	}

	return nil
}


package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

// GitHubClient implements GitHubAPIClient.
type GitHubClient struct {
	client *github.Client
	owner  string
}

// NewGitHubClient creates a GitHub client with authentication.
func NewGitHubClient(token, owner string) *GitHubClient {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)
	return &GitHubClient{
		client: github.NewClient(tc),
		owner:  owner,
	}
}

// CreateBranch checks if the branch exists and creates it if it doesn't.
func (g *GitHubClient) CreateBranch(repo, branchName, baseBranch string) error {
	owner, repoName := g.owner, repo

	// Check if the branch already exists.
	branch, _, err := g.client.Repositories.GetBranch(context.Background(), owner, repoName, branchName, false)
	if err == nil && branch != nil {
		fmt.Printf("⚠️  Branch '%s' already exists in %s. Skipping creation.\n", branchName, repo)
		return nil // Branch exists, skip creation.
	}

	// Get reference of the base branch.
	ref, _, err := g.client.Git.GetRef(context.Background(), owner, repoName, "refs/heads/"+baseBranch)
	if err != nil {
		return fmt.Errorf("failed to get base branch '%s' in %s: %v", baseBranch, repo, err)
	}

	sha := ref.Object.GetSHA()
	newRef := &github.Reference{
		Ref:    github.String("refs/heads/" + branchName),
		Object: &github.GitObject{SHA: &sha},
	}

	_, _, err = g.client.Git.CreateRef(context.Background(), owner, repoName, newRef)
	if err != nil {
		return fmt.Errorf("failed to create branch '%s' in %s: %v", branchName, repo, err)
	}

	fmt.Printf("✅ Branch '%s' created successfully in %s\n", branchName, repo)
	return nil
}

// CreatePullRequest creates a new pull request.
func (g *GitHubClient) CreatePullRequest(repo, prTitle, prBody, branchName, baseBranch string) error {
	owner, repoName := g.owner, repo

	newPR := &github.NewPullRequest{
		Title: github.String(prTitle),
		Body:  github.String(prBody),
		Head:  github.String(branchName),
		Base:  github.String(baseBranch),
	}

	pr, _, err := g.client.PullRequests.Create(context.Background(), owner, repoName, newPR)
	if err != nil {
		return fmt.Errorf("failed to create pull request in %s: %v", repo, err)
	}

	fmt.Printf("✅ Pull request created: %s\n", pr.GetHTMLURL())
	return nil
}

// DeleteBranch deletes a branch from a repository.
func (g *GitHubClient) DeleteBranch(repo, branchName string) error {
	owner, repoName := g.owner, repo
	ref := "refs/heads/" + branchName
	_, err := g.client.Git.DeleteRef(context.Background(), owner, repoName, ref)
	if err != nil {
		return fmt.Errorf("failed to delete branch '%s' in %s: %v", branchName, repo, err)
	}
	fmt.Printf("✅ Rolled back: Branch '%s' deleted from %s\n", branchName, repo)
	return nil
}


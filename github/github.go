// github/github.go - Handles GitHub API interactions (branches & PRs)
// Author: Avneesh Mishra

package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

// GitHubClient struct holds GitHub API client
type GitHubClient struct {
	client *github.Client
	owner  string
}

// ✅ Initialize GitHub Client
func NewGitHubClient(token, owner string) *GitHubClient {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)
	return &GitHubClient{
		client: github.NewClient(tc),
		owner:  owner,
	}
}

// ✅ Create a Branch
func (g *GitHubClient) CreateBranch(repo, branchName, baseBranch string) error {
	owner, repoName := g.owner, repo

	// Check if the branch already exists
	branch, _, err := g.client.Repositories.GetBranch(context.Background(), owner, repoName, branchName, false)
	if err == nil && branch != nil {
		fmt.Printf("⚠️  Branch '%s' already exists in %s. Skipping creation.\n", branchName, repo)
		return nil
	}

	// Get base branch reference
	ref, _, err := g.client.Git.GetRef(context.Background(), owner, repoName, "refs/heads/"+baseBranch)
	if err != nil {
		return fmt.Errorf("failed to get base branch '%s' in %s: %v", baseBranch, repo, err)
	}

	// Create new branch
	sha := ref.Object.GetSHA()
	newRef := &github.Reference{
		Ref:    github.String("refs/heads/" + branchName),
		Object: &github.GitObject{SHA: github.String(sha)},
	}

	_, _, err = g.client.Git.CreateRef(context.Background(), owner, repoName, newRef)
	if err != nil {
		return fmt.Errorf("failed to create branch '%s' in %s: %v", branchName, repo, err)
	}

	fmt.Printf("✅ Branch '%s' created successfully in %s\n", branchName, repo)
	return nil
}

// ✅ Delete a Branch (Rollback)
func (g *GitHubClient) DeleteBranch(repo, branchName string) error {
	owner, repoName := g.owner, repo
	ref := "refs/heads/" + branchName

	// Call GitHub API to delete the branch
	_, err := g.client.Git.DeleteRef(context.Background(), owner, repoName, ref)
	if err != nil {
		return fmt.Errorf("failed to delete branch '%s' in %s: %v", branchName, repo, err)
	}

	fmt.Printf("✅ Deleted branch '%s' in %s\n", branchName, repo)
	return nil
}

// ✅ Create a Pull Request (Returns PR Number)
func (g *GitHubClient) CreatePullRequest(repo, prTitle, prBody, branchName, baseBranch string) (int, error) {
	owner, repoName := g.owner, repo

	newPR := &github.NewPullRequest{
		Title: github.String(prTitle),
		Body:  github.String(prBody),
		Head:  github.String(branchName),
		Base:  github.String(baseBranch),
	}

	pr, _, err := g.client.PullRequests.Create(context.Background(), owner, repoName, newPR)
	if err != nil {
		return 0, fmt.Errorf("failed to create pull request in %s: %v", repo, err)
	}

	fmt.Printf("✅ Pull request #%d created successfully in %s\n", pr.GetNumber(), repo)
	return pr.GetNumber(), nil
}

// ✅ Delete a Pull Request (Rollback)
func (g *GitHubClient) DeletePullRequest(repo string, prNumber int) error {
	owner, repoName := g.owner, repo
	_, resp, err := g.client.PullRequests.Edit(context.Background(), owner, repoName, prNumber, &github.PullRequest{
		State: github.String("closed"),
	})

	if err != nil {
		return fmt.Errorf("failed to close PR #%d in %s: %v", prNumber, repo, err)
	}

	// Ensure PR was successfully closed
	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to close PR #%d in %s: unexpected status code %d", prNumber, repo, resp.StatusCode)
	}

	fmt.Printf("✅ Pull request #%d closed in %s\n", prNumber, repo)
	return nil
}


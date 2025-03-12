// github/github.go - GitHub API functions

package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GitHubClient represents a client for GitHub API
type GitHubClient struct {
	Token string
	Owner string
}

// NewGitHubClient initializes a new GitHub API client
func NewGitHubClient(token, owner string) *GitHubClient {
	return &GitHubClient{
		Token: token,
		Owner: owner,
	}
}

// ListRepositories fetches public and private repositories of the authenticated user
func (c *GitHubClient) ListRepositories() ([]string, error) {
	url := "https://api.github.com/user/repos?per_page=100"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	var repos []struct {
		Name string `json:"name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return nil, err
	}

	// Extract repository names
	var repoNames []string
	for _, repo := range repos {
		repoNames = append(repoNames, repo.Name)
	}

	return repoNames, nil
}

// GetBranchSHA fetches the latest commit SHA of a given branch in a repository
func (c *GitHubClient) GetBranchSHA(repo, branch string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/ref/heads/%s", c.Owner, repo, branch)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	// Parse response to get SHA
	var result struct {
		Object struct {
			SHA string `json:"sha"`
		} `json:"object"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Object.SHA, nil
}

// CreateBranch creates a new branch in the specified repository
func (c *GitHubClient) CreateBranch(repo, branchName, baseBranchSHA string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs", c.Owner, repo)

	// Payload for branch creation
	payload := map[string]string{
		"ref": fmt.Sprintf("refs/heads/%s", branchName),
		"sha": baseBranchSHA, // SHA must be the commit hash of the base branch
	}
	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create branch in %s: %s", repo, resp.Status)
	}

	return nil
}

// DeleteBranch deletes a branch from the specified repository
func (c *GitHubClient) DeleteBranch(repo, branchName string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s", c.Owner, repo, branchName)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete branch in %s: %s", repo, resp.Status)
	}

	return nil
}

// CreatePullRequest creates a pull request in the specified repository
func (c *GitHubClient) CreatePullRequest(repo, title, body, headBranch, baseBranch string) (int, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", c.Owner, repo)

	// Payload for PR creation
	payload := map[string]string{
		"title": title,
		"body":  body,
		"head":  fmt.Sprintf("%s:%s", c.Owner, headBranch),
		"base":  baseBranch,
	}
	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return 0, fmt.Errorf("failed to create PR in %s: %s", repo, resp.Status)
	}

	// Parse response to get PR number
	var response struct {
		Number int `json:"number"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return 0, err
	}

	return response.Number, nil
}

// DeletePullRequest closes and deletes a pull request from the specified repository
func (c *GitHubClient) DeletePullRequest(repo string, prNumber int) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls/%d", c.Owner, repo, prNumber)

	// Close the PR by sending a PATCH request to update its state
	payload := map[string]string{
		"state": "closed",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to close PR #%d in %s: %s", prNumber, repo, resp.Status)
	}

	return nil
}

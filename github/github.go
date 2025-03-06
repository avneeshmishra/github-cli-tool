package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const githubAPI = "https://api.github.com"

type GitHubClient struct {
	Token string
}

func NewGitHubClient(token string) *GitHubClient {
	return &GitHubClient{Token: token}
}

// CreateBranch creates a new branch in the repository
func (c *GitHubClient) CreateBranch(repo, branch string) error {
	url := fmt.Sprintf("%s/repos/%s/git/refs", githubAPI, repo)

	// Get the default branch reference
	ref, err := c.getDefaultBranchRef(repo)
	if err != nil {
		return fmt.Errorf("failed to get default branch: %w", err)
	}

	// Prepare request payload
	payload := map[string]string{
		"ref": "refs/heads/" + branch,
		"sha": ref, // Reference SHA of the default branch
	}

	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create branch: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	return nil
}

// CreatePR creates a pull request
func (c *GitHubClient) CreatePR(repo, branch, title, body string) error {
	url := fmt.Sprintf("%s/repos/%s/pulls", githubAPI, repo)

	payload := map[string]string{
		"title": title,
		"head":  branch,
		"base":  "main", // Assuming "main" is the default branch
		"body":  body,
	}

	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create pull request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	return nil
}

// getDefaultBranchRef retrieves the SHA of the default branch
func (c *GitHubClient) getDefaultBranchRef(repo string) (string, error) {
	url := fmt.Sprintf("%s/repos/%s", githubAPI, repo)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch repo details: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	var result struct {
		DefaultBranch string `json:"default_branch"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", errors.New("failed to parse response")
	}

	// Get SHA of the latest commit on default branch
	return c.getBranchSHA(repo, result.DefaultBranch)
}

// getBranchSHA retrieves the latest commit SHA for a given branch
func (c *GitHubClient) getBranchSHA(repo, branch string) (string, error) {
	url := fmt.Sprintf("%s/repos/%s/branches/%s", githubAPI, repo, branch)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch branch details: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	var result struct {
		Commit struct {
			SHA string `json:"sha"`
		} `json:"commit"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", errors.New("failed to parse response")
	}

	return result.Commit.SHA, nil
}


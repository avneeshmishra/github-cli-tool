package github

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockGitHubClient is a stub implementation for testing.
type MockGitHubClient struct{}

func (m *MockGitHubClient) CreateBranch(repo, branchName, baseBranch string) error {
	if repo == "fail/repo" {
		return errors.New("failed to create branch")
	}
	return nil
}

func (m *MockGitHubClient) CreatePullRequest(repo, prTitle, prBody, branchName, baseBranch string) error {
	if repo == "fail/repo" {
		return errors.New("failed to create PR")
	}
	return nil
}

func (m *MockGitHubClient) DeleteBranch(repo, branchName string) error {
	if repo == "fail/repo" {
		return errors.New("failed to delete branch")
	}
	return nil
}

func TestCreateBranchSuccess(t *testing.T) {
	mock := &MockGitHubClient{}
	err := mock.CreateBranch("test/repo", "feature", "main")
	assert.NoError(t, err)
}

func TestCreatePRFailure(t *testing.T) {
	mock := &MockGitHubClient{}
	err := mock.CreatePullRequest("fail/repo", "Test PR", "Body", "feature", "main")
	assert.Error(t, err)
}

func TestDeleteBranch(t *testing.T) {
	mock := &MockGitHubClient{}
	err := mock.DeleteBranch("test/repo", "feature")
	assert.NoError(t, err)
}


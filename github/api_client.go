package github

// GitHubAPIClient defines an interface for GitHub API interactions.
type GitHubAPIClient interface {
	CreateBranch(repo, branchName, baseBranch string) error
	CreatePullRequest(repo, prTitle, prBody, branchName, baseBranch string) error
	DeleteBranch(repo, branchName string) error
}


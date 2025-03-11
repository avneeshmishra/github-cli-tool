package github

// ✅ Define an interface for GitHub API interactions
type GitHubAPIClient interface {
	CreateBranch(repo, branchName, baseBranch string) error
	CreatePullRequest(repo, prTitle, prBody, branchName, baseBranch string) (int, error)
	DeleteBranch(repo, branchName string) error 
	DeletePullRequest(repo string, prNumber int) error
}


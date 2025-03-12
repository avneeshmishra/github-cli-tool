# go-github-cli 🚀

A CLI tool for managing GitHub repositories, automating branch creation and pull requests across multiple repositories with rollback support.

---

## ✅ Features
- ✅ **Create Git Branches** - Automates branch creation from a base branch across multiple repositories.
- ✅ **Create Pull Requests** - Generates pull requests across multiple repositories with rollback support.
- ✅ **Rollback on Failure** - Deletes created branches or PRs if an error occurs.
- ✅ **Interactive Repository Selection** - Allows users to manually select repositories instead of processing all.
- ✅ **Handles Repos Without Base Branch** - Skips repositories that don’t have a base branch instead of failing.
- ✅ **Handles Invalid User Inputs** - Prevents invalid repo selection and ensures only valid repositories are processed.
- ✅ **Error Logging** - Displays errors when GitHub API requests fail, with detailed output.
- ✅ **Modular Design** - Organized into `cmd/`, `github/`, and `utils/` packages for maintainability.
- ✅ **TDD & Interface Abstraction** - Uses an interface (`GitHubAPIClient`) to abstract API interactions for easy testing.

---

## 📌 Prerequisites
Before using the CLI, ensure the following:

1. **Go installed** (`go version`)
2. **GitHub Personal Access Token (PAT)** with required permissions.
3. **Set up environment variables**:
   ```sh
   export GITHUB_TOKEN="your_github_token_here"
   export GITHUB_OWNER="your_github_username_or_org"
   ```

---

## 📥 Installation
Clone the repository and build the binary:
```sh
git clone https://github.com/yourusername/go-github-cli.git
cd go-github-cli
go build -o go-github-cli
```

To use it globally, move it to `/usr/local/bin/`:
```sh
sudo mv go-github-cli /usr/local/bin/
```

---

## 🚀 Usage

### 1️⃣ Create a Branch in Multiple Repositories
Creates a branch from a base branch across multiple repositories.

#### Command:
```sh
./go-github-cli create-branch --repo="repo1,repo2" --branch="feature-update" --base="main"
```

#### Interactive Repo Selection (if `--repo` is not specified):
```sh
./go-github-cli create-branch --branch="feature-update"
```
_(The tool will prompt you to select repositories manually.)_

---

### 2️⃣ Create a Pull Request Across Multiple Repositories
Creates a PR from a feature branch to the base branch.

#### Command:
```sh
./go-github-cli create-pr --repo="repo1,repo2" --branch="feature-update" --base="main" --title="New Feature"
```

#### Rollback on PR Failure:
```sh
./go-github-cli create-pr --repo="repo1,repo2" --branch="feature-update" --rollback
```
_(If a PR fails, already created PRs are deleted to maintain consistency.)_

---

### 3️⃣ Rollback on Failure
If an error occurs while creating branches or PRs, the tool can automatically delete previously created branches/PRs.

#### Command:
```sh
./go-github-cli create-branch --repo="repo1,repo2" --branch="feature-update" --rollback
```

---

### 🔍 Error Handling
- **Handles missing base branches** - Skips repositories where the base branch is missing.
- **Prevents duplicate repository selection** - Ensures a repository is not processed twice.
- **Logs API failures** - Displays detailed errors when API requests fail.

---

## 🔧 Development & Contribution
To contribute:
1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-branch`
3. Commit changes: `git commit -m "Add new feature"`
4. Push changes: `git push origin feature-branch`
5. Open a Pull Request.

---

## 📜 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Author
Developed by **Avneesh Mishra**  
[GitHub](https://github.com/avneeshmishra) 

---


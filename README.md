# go-github-cli

**go-github-cli** is a command-line tool written in Go that automates GitHub repository management. It allows you to create branches and pull requests (PRs) across multiple repositories, supports rollback (abort/cleanup), and follows test-driven development (TDD) principles.

---

## **Table of Contents**
- [Overview](#overview)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Post-Installation Steps](#post-installation-steps)
- [Environment Variables](#environment-variables)
- [Usage](#usage)
  - [Create Branch](#create-branch)
  - [Create Pull Request](#create-pull-request)
  - [Rollback (Abort/Cleanup) Option](#rollback-abortcleanup-option)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [Future Enhancements](#future-enhancements)
- [License](#license)

---

## **Overview**
**go-github-cli** simplifies GitHub repository management by automating:
- **Branch creation across multiple repositories.**
- **Pull request creation with automated checks.**
- **Rollback (abort/cleanup) option:** If an error occurs, previously created branches are deleted.
- **Interactive repository selection** (removes duplicates).
- **Error handling for API failures, rate limits, and permission issues.**
- **Test-driven development (TDD):** Includes unit tests using mock APIs.
- **Modular and maintainable code design.**

---

## **Features**
- ✅ **Create Git Branches:** Automate branch creation from a base branch.
- ✅ **Create Pull Requests:** Generate pull requests across multiple repositories.
- ✅ **Rollback on Failure:** If an error occurs, previously created branches can be deleted.
- ✅ **Interactive Repository Selection:** Prevents duplicate repository processing.
- ✅ **TDD & Interface Abstraction:** Uses an interface (`GitHubAPIClient`) for API interactions.
- ✅ **Error Handling:** Includes API rate limit handling and permission checks.
- ✅ **Modular Design:** Organized into `cmd`, `github`, and `utils` packages.

---

## **Requirements**
- **Go 1.18+**
- **GitHub Personal Access Token (PAT)** with `repo` scope.
- **GitHub username or organization name**.

---

## **Installation**

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/go-github-cli.git
   cd go-github-cli
   ```

2. **Initialize Go modules and install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Build the CLI tool:**
   ```sh
   go build -o go-github-cli
   ```

---

## **Post-Installation Steps**

### **Verify Installation**
After building the tool, you can verify the installation by running:
```sh
./go-github-cli --help
```

If the installation was successful, you should see the list of available commands and options.

### **Set Up Environment Variables**
Before using the CLI, set the following environment variables:

```sh
export GITHUB_TOKEN="your_actual_github_token"
export GITHUB_OWNER="your_github_username_or_org"
```

To persist these settings, add them to your shell profile:
```sh
echo 'export GITHUB_TOKEN="your_actual_github_token"' >> ~/.bashrc
echo 'export GITHUB_OWNER="your_github_username_or_org"' >> ~/.bashrc
source ~/.bashrc
```
For **macOS/Linux (zsh users)**:
```sh
echo 'export GITHUB_TOKEN="your_actual_github_token"' >> ~/.zshrc
source ~/.zshrc
```

For **Windows (PowerShell users)**:
```powershell
[System.Environment]::SetEnvironmentVariable("GITHUB_TOKEN", "your_actual_github_token", "User")
[System.Environment]::SetEnvironmentVariable("GITHUB_OWNER", "your_github_username_or_org", "User")
```

---

## **Troubleshooting**

### **1. Command Not Found**
If you see an error like:
```sh
bash: ./go-github-cli: No such file or directory
```
Ensure you are in the correct directory and the file has execution permissions:
```sh
ls -la go-github-cli
chmod +x go-github-cli
./go-github-cli --help
```

### **2. GitHub API Authentication Errors**
If you get authentication errors, verify that your GitHub token is valid:
```sh
curl -H "Authorization: token $GITHUB_TOKEN" https://api.github.com/user
```
If this fails, generate a new token from **GitHub → Settings → Developer Settings → Personal Access Tokens**.

### **3. Permission Denied on Windows**
If you face permission errors, run the command as an administrator:
```powershell
Start-Process PowerShell -Verb RunAs
```
Then retry executing the `go-github-cli` binary.

---


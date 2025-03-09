# go-github-cli

**go-github-cli** is a command-line tool written in Go for managing GitHub repositories. It automates the creation of Git branches and corresponding pull requests (PRs) across multiple repositories. The tool supports interactive repository selection, a rollback (abort/cleanup) mechanism, and follows modular design with Test-Driven Development (TDD) principles.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
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

## Overview

**go-github-cli** simplifies the management of GitHub repositories by enabling you to create branches and pull requests across multiple repositories in one command. The tool ensures:
- **Unique repository selection** (removes duplicates)
- **Rollback functionality:** If an error occurs during branch creation, any branches created earlier in the run are automatically deleted.
- **Robust error handling** for API failures, rate limits, and permission issues.
- **Interface abstraction and TDD:** Uses an interface for GitHub API interactions with a stub implementation for unit testing.
- **Modular and maintainable design:** Separate packages for CLI commands, GitHub API interactions, and utility functions.

---

## Features

- **Create Git Branches:** Automatically create new branches from a base branch across multiple repositories.
- **Create Pull Requests:** Generate pull requests using the created branches.
- **Interactive Repository Selection:** Deduplicates and (optionally) allows interactive selection of repositories.
- **Rollback (Abort/Cleanup):** If an error occurs during branch creation, previously created branches can be rolled back (deleted).
- **TDD & Interface Abstraction:** The tool is designed with unit tests and an interface-based API client for easy testing and future extension.
- **Robust Error Handling:** Clear messaging for API errors, rate limits, and permission issues.
- **Modular Design:** Organized into separate packages (`cmd`, `github`, and `utils`) for maintainability.

---

## Requirements

- Go 1.18 or higher.
- A GitHub Personal Access Token (PAT) with sufficient permissions (e.g., `repo` scope).
- Your GitHub username or organization name.

---

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/go-github-cli.git
   cd go-github-cli


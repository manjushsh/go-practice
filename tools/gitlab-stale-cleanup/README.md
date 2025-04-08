# GitLab Stale Branch Cleaner

A simple command-line tool built in Go to automatically delete stale branches in a GitLab repository. The tool uses GitLab's REST API to interact with your repository and removes branches that haven't been updated in a specified number of days. It also supports a dry-run mode to show which branches would be deleted without actually performing the deletion.

## TODO
- [] Release as tool or package
- [] Implement CI for build
- [] Implement test cases (optional)

## Features

- Deletes stale branches in GitLab repositories based on a specified days threshold.
- Supports a dry-run mode to preview which branches would be deleted.
- Protects specific branches (e.g., `main`, `master`) from deletion.
- Lightweight and written in Go with concurrency support for future scalability.
- CLI-based tool with configurable options.

## Table of Contents

1. [Installation](#installation)
2. [Configuration](#configuration)
3. [Usage](#usage)
4. [Flags](#flags)
5. [Example](#example)

## Installation

To install the tool, you can either build it from source or install it directly using a pre-built binary.

### Option 1: Build from Source

1. Clone the repository:

```bash
git clone https://github.com/yourusername/gitlab-stale-branch-cleaner.git
cd gitlab-stale-branch-cleaner
```

### 1. Download packages
```bash
go mod tidy && go mod download
```

### 2. Build App
```bash
rm ./gitlab-stale-branch-cleaner && go build -o gitlab-stale-branch-cleaner
```
and make it executable
```bash
chmod +x gitlab-stale-branch-cleaner
```

### 3. Move the binary to a directory in your $PATH (optional):
```bash
sudo mv gitlab-stale-branch-cleaner /usr/local/bin/
```

### 4. Run the Application: Now, run the application with the appropriate flags:
```bash
./gitlab-stale-branch-cleaner --project-id="your_gitlab_project_id" --private-token="your_private_token" --stale-days-threshold=30 --dry-run --protected-branches="main master"
```

## Configuration

This tool requires the following input from the user:

- GitLab project ID: The ID of your GitLab project.
- GitLab private token: A GitLab personal access token with permissions to view and delete branches.
- Stale days threshold: The number of days after which a branch is considered stale.
- Protected branches: A comma or space-separated list of branches you want to protect from deletion (e.g., main,master).
- Dry run mode: Flag to simulate the deletions without actually removing any branches.

## GitLab Personal Access Token

To create a personal access token, follow these steps:
- Go to your GitLab account.
- Navigate to User Settings > Access Tokens.
- Create a token with api and read_repository scopes.


## Usage

Once installed, you can use the tool from the command line.
Basic Usage
```bash
gitlab-stale-branch-cleaner --project-id <PROJECT_ID> --private-token <PRIVATE_TOKEN> --stale-days-threshold <DAYS_THRESHOLD> --dry-run --protected-branches "main master"
```

## Flags

|Flag                    |	Description                                             |	Default Value  |
|------------------------|----------------------------------------------------------|------------------|
|   `--project-id`       |	The GitLab project ID where branches will be managed.   |	Required       |
|`--private-token`       |	Your GitLab personal access token.                      |	Required       |
|`--stale-days-threshold`|	The number of days after which a branch is considered stale.	30
|   `--dry-run`          |	Flag to preview which branches would be deleted without actually deleting them.|	false   |
|`--protected-branches`  |	A comma or space-separated list of branches to protect from deletion.|	`main master`   |

## Example

Here’s an example of how to run the tool with the necessary flags:
```bash
gitlab-stale-branch-cleaner \
  --project-id "12345678" \
  --private-token "your-private-token-here" \
  --stale-days-threshold 30 \
  --dry-run \
  --protected-branches "main master"
```
or
```bash
./gitlab-stale-branch-cleaner --project-id "12345678" --private-token "glpat-xxxxxxxxx" --stale-days-threshold 30 --dry-run --protected-branches "main master"
```

if errors,
```bash
./gitlab-stale-branch-cleaner --project-id "12345678" --private-token "glpat-xxxxxxxxx" --stale-days-threshold 30 --dry-run --protected-branches "main,master"
```

## Actual Deletion Example

To actually delete stale branches, remove the --dry-run flag:
```bash
gitlab-stale-branch-cleaner \
  --project-id "12345678" \
  --private-token "your-private-token-here" \
  --stale-days-threshold 60 \
  --protected-branches "main master"
```

## How It Works

1. Fetch Branches: The tool first fetches all branches in the GitLab project using the GitLab API.
2. Identify Stale Branches: It then checks the last commit date of each branch. If the branch has not been updated for more than the specified threshold of days, it is considered stale.
3. Protected Branches: Any branches that match the protected branch list are skipped and not considered for deletion.
4. Dry Run: If the --dry-run flag is provided, the tool only simulates the deletions and outputs which branches would be deleted.
5. Deletion: In actual run mode (without --dry-run), the tool deletes stale branches via the GitLab API.

## Development

If you'd like to contribute to the project, here are a few things you can do:
1. Fork the repository.
2. Make your changes or add new features.
3. Run the tests to ensure everything works correctly.
4. Create a pull request to the main repository.

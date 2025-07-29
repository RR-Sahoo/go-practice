# GitHub User Activity CLI

A simple command-line tool to fetch and display recent GitHub activity for any user.

## Features

- 🔍 Fetch recent GitHub events for any username
- 📊 Display different types of activities with emojis:
  - 🔼 Push events (commits)
  - 🐛 Issue events (created, closed, etc.)
  - 💬 Issue comment events
  - ⭐ Watch events (starring repositories)
  - 🔃 Pull request events
- 🎯 Clean, readable output format
- ⚡ Fast and lightweight

## Installation

1. Clone or download this repository
2. Navigate to the project directory
3. Run the application with Go:

```bash
go run main.go <username>
```

## Usage

### Basic Usage

```bash
go run main.go octocat
```

### Example Output

```
🔼 Pushed 3 commits to octocat/Hello-World
🐛 Opened an issue in octocat/test-repo
⭐ Starred microsoft/vscode
💬 Created a comment on an issue in octocat/Hello-World
```

## Supported Event Types

- **PushEvent**: Shows number of commits pushed to a repository
- **IssuesEvent**: Shows issue actions (opened, closed, etc.)
- **IssueCommentEvent**: Shows comment actions on issues
- **WatchEvent**: Shows when a repository was starred
- **PullRequestEvent**: Shows pull request actions

## Requirements

- Go 1.16 or higher
- Internet connection to access GitHub API

## API Rate Limits

This tool uses GitHub's public API which has rate limits:

- 60 requests per hour for unauthenticated requests
- 5000 requests per hour for authenticated requests

## Error Handling

The tool handles various error scenarios:

- Invalid usernames
- Network connectivity issues
- API rate limiting
- Empty activity results

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the MIT License.

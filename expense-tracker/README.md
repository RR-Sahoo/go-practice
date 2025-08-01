# Expense Tracker CLI

A simple command-line tool to track and manage your daily expenses. Store, view, and summarize your spending with ease.

## Features

- 💰 **Add expenses** with description and amount
- 📋 **List all expenses** with details
- 📊 **Summary view** showing total spending
- 💾 **Persistent storage** using CSV format
- 🆔 **Auto-generated IDs** for each expense
- 📅 **Automatic date tracking** for all entries

## Installation

1. Clone or download this repository
2. Navigate to the expense-tracker directory
3. Run the application with Go:

```bash
go run main.go <command>
```

## Usage

### Add an Expense

```bash
go run main.go add --description "Lunch" --amount 15
go run main.go add --description "Coffee" --amount 5
go run main.go add --description "Gas" --amount 45
```

### List All Expenses

```bash
go run main.go list
```

**Example Output:**

```
ID: 1, Description: Lunch, Amount: 15, Date: 2024-01-15
ID: 2, Description: Coffee, Amount: 5, Date: 2024-01-15
ID: 3, Description: Gas, Amount: 45, Date: 2024-01-15
```

### View Summary

```bash
go run main.go summary
```

**Example Output:**

```
Total expense: 65
```

## Data Storage

- **File Format**: CSV (Comma-Separated Values)
- **File Name**: `expenses.csv`
- **Location**: Same directory as the application
- **Structure**: Description, Date, Amount, ID

## Commands

| Command   | Description          | Example                                |
| --------- | -------------------- | -------------------------------------- |
| `add`     | Add a new expense    | `add --description "Food" --amount 20` |
| `list`    | Display all expenses | `list`                                 |
| `summary` | Show total spending  | `summary`                              |

## Requirements

- Go 1.16 or higher
- No external dependencies

## File Structure

```
expense-tracker/
├── main.go          # Main application code
├── expenses.csv     # Data storage file (auto-generated)
└── README.md        # This file
```

## Error Handling

The application handles various scenarios:

- Missing command arguments
- Invalid amount values
- File read/write errors
- CSV parsing errors

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the MIT License.

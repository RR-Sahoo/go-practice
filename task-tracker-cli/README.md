# Task Tracker CLI

A simple command-line tool written in Go to manage your daily tasks. It allows you to add, update, delete, and track task progress directly from the terminal. Tasks are stored in a local JSON file.

## 🔗 Project URL

[https://github.com/RR-Sahoo/go-practice.git](https://github.com/RR-Sahoo/go-practice.git)

---

## 📦 Features

- Add new tasks
- Update existing tasks
- Delete tasks
- Mark tasks as in-progress or done
- List all tasks or filter by status (`todo`, `in-progress`, `done`)
- Stores tasks in a local `tasks.json` file

---

## 🚀 Usage

```bash
# Add a task
go run main.go add "Buy groceries"

# Update task description
go run main.go update 1 "Buy groceries and cook dinner"

# Delete a task
go run main.go delete 1

# Mark task as in-progress or done
go run main.go mark-in-progress 2
go run main.go mark-done 2

# List all tasks
go run main.go list

# List tasks by status
go run main.go list todo
go run main.go list in-progress
go run main.go list done
```

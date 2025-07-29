package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var tasks []Task

func loadTasks() {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		tasks = []Task{}
		return
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		tasks = []Task{}
	}
}
func getNextID() int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}
func saveTasks() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	os.WriteFile("tasks.json", data, 0644)
}

func main() {
	loadTasks()

	if len(os.Args) < 2 {
		fmt.Println("Expected a command.")
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Expected a task description.")
			return
		}
		description := os.Args[2]
		task := Task{
			ID:          getNextID(),
			Description: description,
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		tasks = append(tasks, task)
		saveTasks()
		fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: update <id> <new description>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}
		for i, task := range tasks {
			if task.ID == id {
				task.Description = os.Args[3]
				task.UpdatedAt = time.Now()
				tasks[i] = task
				saveTasks()
				fmt.Printf("Updated task %d: %s\n", id, os.Args[3])
				return
			}
		}
		fmt.Printf("Task with ID %d not found.\n", id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Expected a task ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}
		for i, task := range tasks {
			if task.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				saveTasks()
				fmt.Printf("Deleted task %d\n", id)
				return
			}
		}
		fmt.Printf("Task with ID %d not found.\n", id)
	case "mark-in-progress", "mark-done":
		if len(os.Args) < 3 {
			println("Expected a task ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			println("Invalid task ID.")
			return
		}
		for i, task := range tasks {
			if task.ID == id {
				newStatus := "in-progress"
				if cmd == "mark-done" {
					newStatus = "done"
				}
				task.Status = newStatus
				task.UpdatedAt = time.Now()
				tasks[i] = task
				saveTasks()
				fmt.Printf("Marked task %d as %s\n", id, newStatus)
				return
			}
		}
	case "list":
		if len(os.Args) == 2 {
			printTasks(tasks)
		} else {
			filter := os.Args[2]
			var filteredTasks []Task
			for _, task := range tasks {
				if task.Status == filter {
					filteredTasks = append(filteredTasks, task)
				}
			}
			printTasks(filteredTasks)
		}

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Available commands: add, update, delete, mark-in-progress, mark-done, list, filter")
	}
}
func printTasks(taskList []Task) {
	if len(taskList) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range taskList {
		fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n",
			task.ID, task.Description, task.Status,
			task.CreatedAt.Format("2006-01-02 15:04:05"),
			task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

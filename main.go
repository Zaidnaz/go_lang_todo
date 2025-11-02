package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1. DATA STRUCTURE
// Todo defines the structure for a single to-do item.
// Note: The fields must be capitalized (exported) for the json package
// to be able to see and encode/decode them.
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// tasks is a slice to hold all our Todo items
var tasks []Todo

// todoFileName is the name of the file where we'll store our tasks
const todoFileName = ".todos.json"

// 2. MAIN FUNCTION (Entry Point)
func main() {
	// Load existing tasks from the file when the program starts.
	// We ignore the error here, as a missing file is okay on first run.
	loadTasks()

	// Get all arguments from the command line, except the program name
	args := os.Args[1:]

	// If no arguments are provided, show the task list by default
	if len(args) == 0 {
		handleList()
		return
	}

	// 3. COMMAND HANDLING
	// Use a switch statement to handle different commands (add, list, complete)
	switch args[0] {
	case "list":
		handleList()

	case "add":
		// Join all arguments after "add" to form the task text
		taskText := strings.Join(args[1:], " ")
		if taskText == "" {
			fmt.Println("Error: Missing task description. Usage: todo add \"your task here\"")
			return
		}
		handleAdd(taskText)

	case "complete":
		if len(args) < 2 {
			fmt.Println("Error: Missing task number. Usage: todo complete <task_number>")
			return
		}
		// Convert the task number (string) to an integer
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Invalid task number. Must be an integer.")
			return
		}
		handleComplete(id)

	default:
		fmt.Printf("Unknown command: %s\n", args[0])
		fmt.Println("Usage: todo [list|add|complete]")
	}
}

// 4. COMMAND LOGIC FUNCTIONS

// handleList prints all tasks to the console
func handleList() {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show. Add one with 'todo add \"my task\"'")
		return
	}

	for _, task := range tasks {
		status := " " // Not completed
		if task.Completed {
			status = "✔" // Completed
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Task)
	}
}

// handleAdd adds a new task to the list
func handleAdd(text string) {
	// Find the next available ID
	nextID := 1
	if len(tasks) > 0 {
		nextID = tasks[len(tasks)-1].ID + 1
	}

	// Create the new task
	newTask := Todo{
		ID:        nextID,
		Task:      text,
		Completed: false,
	}

	// Add it to the slice
	tasks = append(tasks, newTask)

	// Save the updated slice to the file
	if err := saveTasks(); err != nil {
		fmt.Printf("Error saving task: %v\n", err)
		return
	}

	fmt.Printf("Added task: \"%s\"\n", text)
}

// handleComplete marks a task as completed
func handleComplete(id int) {
	taskFound := false
	for i := range tasks {
		// Find the task with the matching ID
		if tasks[i].ID == id {
			tasks[i].Completed = true
			taskFound = true
			break
		}
	}

	if !taskFound {
		fmt.Printf("Error: Task with ID %d not found.\n", id)
		return
	}

	// Save the change to the file
	if err := saveTasks(); err != nil {
		fmt.Printf("Error saving task: %v\n", err)
		return
	}

	fmt.Printf("Completed task %d.\n", id)
}

// 5. FILE HELPER FUNCTIONS

// saveTasks writes the current 'tasks' slice to the JSON file
func saveTasks() error {
	// Encode the 'tasks' slice into JSON format with indentation
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	// Write the JSON data to the file, overwriting it
	return os.WriteFile(todoFileName, data, 0644)
}

// loadTasks reads the JSON file and loads tasks into the 'tasks' slice
func loadTasks() error {
	// Check if the file exists
	if _, err := os.Stat(todoFileName); os.IsNotExist(err) {
		// File doesn't exist, which is fine. Just return.
		return nil
	}

	// Read the file's content
	data, err := os.ReadFile(todoFileName)
	if err != nil {
		return err
	}

	// Decode the JSON data into the 'tasks' slice
	// We pass a pointer (&tasks) so the function can modify the slice.
	return json.Unmarshal(data, &tasks)
}


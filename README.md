Go To-Do CLI

A simple and efficient command-line (CLI) to-do list application written in Go. This project was built as a fun introduction to the Go programming language, covering concepts like file I/O, JSON encoding, and handling command-line arguments.

✨ Features

Add: Quickly add new tasks to your list.

List: View all your pending and completed tasks.

Complete: Mark tasks as complete by their ID.

Persistent: Your tasks are saved locally in a .todos.json file so they're always there when you come back.

📦 Installation

To get started, you'll need to have Go installed on your system (version 1.25+ recommended).

Clone this repository:

git clone [https://github.com/Zaidnaz/go_lang_todo](https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git)
cd YOUR_REPO_NAME


Build the application:

go build


This will create an executable file (e.g., go_project1.exe on Windows or go_project1 on macOS/Linux).

🚀 Usage

You can run the application directly from your terminal.

List All Tasks

Shows all current tasks. This is the default action if no command is given.

.\go_project1.exe list


or simply:

.\go_project1.exe


Output:

[✔] 1: Learn Go
[ ] 2: Build a CLI app


Add a New Task

Add a new, incomplete task to your list.

.\go_project1.exe add "Review my pull request"


Output:

Added task: "Review my pull request"


Complete a Task

Mark an existing task as completed using its ID.

.\go_project1.exe complete 2


Output:

Completed task 2.

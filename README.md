# Task Tracker (CLI)

A simple command-line ToDo application written in Go.
All tasks are stored in a JSON file located in the user’s home directory:
- Windows: C:\Users\<user>\.tasktracker\tasks.json
- Linux/macOS: /home/<user>/.tasktracker/tasks.json
The app automatically creates the directory and file if they do not exist.

# Features
- Add new tasks
- Update task description
- Delete tasks
- Mark tasks as in-progress or done
- List all tasks or filter by status

# Instalation (Linux)

- git clone https://github.com/HigorRamosM/task-tracker.git

- cd task-tracker

- go build -o task-tracker

- chmod +x task-tracker

# Instalation (Windows)

- git clone https://github.com/HigorRamosM/task-tracker.git

- cd task-tracker.exe

- go build -o task-tracker.exe

To run, use:

- task-tracker (linux)

- .\task-tracker.exe (windows)

# Available Commands

- add [task]            - Add a new task

- update [id] [task]    - Update an existing task

- delete [id]           - Delete a task

- mark-in-progress [id] - Mark a task as in progress

- mark-done [id]        - Mark a task as done

- list                  - List all tasks

- list done             - List all completed tasks

- list todo             - List all pending tasks

- list in-progress      - List all tasks in progress

# Examples
(On windows replace task-tracker with task-tracker.exe)

- task-tracker add "Buy groceries"

- task-tracker update 1 "Buy groceries and cook dinner"

- task-tracker delete 1

- task-tracker mark-in-progress 2

- task-tracker mark-done 3

- task-tracker list

- task-tracker list done

- task-tracker list todo

- task-tracker list in-progress


https://roadmap.sh/projects/task-tracker

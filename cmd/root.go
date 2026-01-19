package cmd

import (
	"fmt"
	"os"

	"github.com/HigorRamosM/task-tracker/internal/task"
)

const (
	add            string = "add"
	update         string = "update"
	delete         string = "delete"
	markInProgress string = "mark-in-progress"
	markDone       string = "mark-done"
	list           string = "list"
)

var commandsList = fmt.Sprintf(`Commands:
	%s [task] - Add a new task
	%s [id] [task] - Update an existing task
	%s [id] - Delete a task
	%s [id] - Mark a task as in progress
	%s [id] - Mark a task as done
	%s - List all tasks
	list done - List all completed tasks
	list todo - List all pending tasks
	list in-progress - List all tasks in progress`,
	add, update, delete, markInProgress, markDone, list,
)

var unknownCmdMsg = "Command not recognized. Use one of the followinf commands:\n" + commandsList

func RunCLI(t task.TasksList) {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println(commandsList)
		return
	}

	switch os.Args[1] {
	case add:
		RunAddTask(os.Args, t)
	case list:
		RunListTasks(os.Args, t)
	case update:
		RunUpdateTask(os.Args, t)
	case delete:
		RunDeleteTask(os.Args, t)
	case markInProgress, markDone:
		RanMarkTask(os.Args, t)
	default:
		fmt.Println(unknownCmdMsg)
	}
}

package cmd

import (
	"fmt"

	"github.com/HigorRamosM/task-tracker/internal/task"
)

func PrintTasks(id int, description, status, createdAt, updatedAt string) {
	fmt.Printf("    %d - %s - status: %s", id, description, status)
	if createdAt != updatedAt {
		fmt.Printf(" - Updated at %s\n", updatedAt)
	} else {
		fmt.Printf(" - Created at %s\n", createdAt)
	}
}

func RunListTasks(s []string, t task.TasksList) {
	if len(s) < 2 || len(s) > 3 {
		fmt.Printf("Unkown command %s, use list", s[1:])
		return
	}
	taskList, err := t.ListTasks(s)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	if len(taskList.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range taskList.Tasks {
		PrintTasks(task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

package cmd

import (
	"fmt"
	"strings"

	"github.com/HigorRamosM/task-tracker/internal/task"
)

func RunUpdateTask(s []string, t task.TasksList) {
	if len(s) != 4 {
		command := strings.Join(s[1:], " ")
		fmt.Printf("Unkown command %s, use update [id] [task]", command)
		return
	}
	id := s[2]
	newTask := s[3]
	err := t.UpdateTask(id, newTask)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}
	fmt.Print("Task updated successfully")
}

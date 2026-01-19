package cmd

import (
	"fmt"
	"strings"

	"github.com/HigorRamosM/task-tracker/internal/task"
)

func RunAddTask(s []string, t task.TasksList) {
	if len(s) != 3 {
		command := strings.Join(s[1:], " ")
		fmt.Printf("Unkown command %s, use add [task]", command)
		return
	}
	description := s[2]
	id, err := t.AddTask(description)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}
	fmt.Printf("Task added successfully (ID: %d)", id)
}

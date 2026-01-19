package cmd

import (
	"fmt"
	"strings"

	"github.com/HigorRamosM/task-tracker/internal/task"
)

func RunDeleteTask(s []string, t task.TasksList) {
	if len(s) != 3 {
		command := strings.Join(s[1:], " ")
		fmt.Printf("Unkown command %s, use delete [id]", command)
		return
	}
	id := s[2]
	err := t.DeleteTask(id)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Deleted task successfully")
}

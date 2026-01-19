package cmd

import (
	"fmt"
	"strings"

	"github.com/HigorRamosM/task-tracker/internal/task"
)

func RanMarkTask(s []string, t task.TasksList) {
	if len(s) != 3 {
		command := strings.Join(s[1:], " ")
		fmt.Printf("Unkown command %s, use update [id] [task]", command)
		return
	}
	id := s[2]
	cmdStatus := s[1]
	status := strings.TrimPrefix(cmdStatus, "mark-")
	err := t.MarkStatus(id, status)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Marked task %s as %s", id, status)
}

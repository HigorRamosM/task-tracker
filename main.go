package main

import (
	"github.com/HigorRamosM/task-tracker/cmd"
	"github.com/HigorRamosM/task-tracker/internal/task"
)

func main() {
	task := task.TasksList{}
	task.Load()
	cmd.RunCLI(task)
}

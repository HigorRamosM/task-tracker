package task

import (
	"fmt"
	"strconv"
	"time"
)

var ErrorEmptyDescription error = fmt.Errorf("task description cannot be empty")
var ErrorEmptyStatus error = fmt.Errorf("task status cannot be empty")
var ErrorInvalidTypeID = fmt.Errorf("invalid format, use an integer value for the id")
var ErrorTaskNotFound = fmt.Errorf("task not found")
var ErrorInvalidStatus = fmt.Errorf("invalid status, use todo, in-progress or done")

// Add a new task with the given description.
// Returns the ID of the newly created task or an error if the description is empty.
func (t *TasksList) AddTask(description string) (int, error) {
	if description == "" {
		return 0, ErrorEmptyDescription
	}

	var lastID int
	if len(t.Tasks) > 0 {
		lastID = t.Tasks[len(t.Tasks)-1].ID
	}

	newTask := TaskItem{
		ID:          lastID + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Tasks = append(t.Tasks, newTask)
	t.Save()

	return newTask.ID, nil
}

// Update the description of a task by its ID.
// Returns an error if the task is not found or if the new description is empty.
func (t *TasksList) UpdateTask(taskID, newDrescription string) error {
	if newDrescription == "" {
		return ErrorEmptyDescription
	}
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return ErrorInvalidTypeID
	}
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].Description = newDrescription
			t.Tasks[i].UpdatedAt = time.Now()
			t.Save()
			return nil
		}
	}
	return ErrorTaskNotFound
}

// Delete a task by its ID.
// Returns an error if the task is not found.
func (t *TasksList) DeleteTask(taskID string) error {
	id, err := strconv.Atoi(taskID)
	var deleted bool = false
	if err != nil {
		return ErrorInvalidTypeID
	}
	newTaskList := TasksList{}
	for _, task := range t.Tasks {
		if task.ID != id {
			newTaskList.Tasks = append(newTaskList.Tasks, task)
		} else {
			deleted = true
		}
	}
	t = &newTaskList
	t.Save()

	if !deleted {
		return ErrorTaskNotFound
	}
	return nil
}

// MarkStatus updates the status of a task by its ID.
// Returns an error if the task is not found or if the id is not numeric.
func (t *TasksList) MarkStatus(taskID, status string) error {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return ErrorInvalidTypeID
	}
	switch status {
	case "in-progress", "done":
		for i, task := range t.Tasks {
			if task.ID == id {
				t.Tasks[i].Status = status
				t.Tasks[i].UpdatedAt = time.Now()
				t.Save()
				return nil
			}
		}
	}
	return ErrorTaskNotFound
}

// ListTasks returns a list of tasks, optionally filtered by status.
// If no status is provided, all tasks are returned.
// Returns an error if the provided status is invalid.
func (t *TasksList) ListTasks(s []string) (TasksList, error) {
	if t == nil {
		return TasksList{}, nil
	}
	if len(s) == 2 {
		return *t, nil
	}
	status := s[2]
	if status != "todo" && status != "in-progress" && status != "done" {
		return TasksList{}, ErrorInvalidStatus
	}
	filteredTasks := TasksList{}

	for _, task := range t.Tasks {
		if task.Status == status {
			filteredTasks.Tasks = append(filteredTasks.Tasks, task)
		}
	}
	return filteredTasks, nil
}

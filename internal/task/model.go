package task

import "time"

type TaskItem struct {
	ID          int       `json:"ID"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TasksList struct {
	Tasks []TaskItem `json:"tasks"`
}

package api

import "web-rk2/internal/entities"

type Usecase interface {
	CreateTask(entities.Task) (*entities.Task, error)
	ListTasks() ([]*entities.Task, error)
	GetTaskByID(id int) (*entities.Task, error)
	UpdateTaskByID(id int, user entities.Task) (*entities.Task, error)
	DeleteTaskByID(id int) error
}

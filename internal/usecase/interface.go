package usecase

import "web-rk2/internal/entities"

type Provider interface {
	InsertTask(entities.Task) (*entities.Task, error)
	SelectAllTasks() ([]*entities.Task, error)

	SelectTaskByID(id int) (*entities.Task, error)
	SelectTaskByAssigneeName(name string) (*entities.Task, error)
	SelectTaskByAuthorName(author_name string) (*entities.Task, error)

	UpdateTaskByID(id int, user entities.Task) (*entities.Task, error)
	DeleteTaskByID(id int) error
}

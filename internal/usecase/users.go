package usecase

import (
	"fmt"
	"web-rk2/internal/entities"
)

func (u *Usecase) CreateTask(user entities.Task) (*entities.Task, error) {
	if user, err := u.p.SelectTaskByAuthorName(user.Author_name); err != nil {
		return nil, err
	} else if user != nil {
		return nil, entities.ErrUserEmailConflict
	}

	if user, err := u.p.SelectTaskByAssigneeName(user.Assignee_name); err != nil {
		return nil, err
	} else if user != nil {
		return nil, entities.ErrUserNameConflict
	}

	if user.Status != "new" && user.Status != "in progress" && user.Status != "done" {
		return nil, entities.ErrUserStatusConflict
	}

	createdUser, err := u.p.InsertTask(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *Usecase) ListTasks() ([]*entities.Task, error) {
	users, err := u.p.SelectAllTasks()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *Usecase) GetTaskByID(id int) (*entities.Task, error) {
	user, err := u.p.SelectTaskByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entities.ErrUserNotFound
	}

	return user, nil
}

func (u *Usecase) UpdateTaskByID(id int, user entities.Task) (*entities.Task, error) {
	oldUser, err := u.p.SelectTaskByID(id)
	if err != nil {
		return nil, err
	}

	if user, err := u.p.SelectTaskByAssigneeName(user.Assignee_name); err != nil {
		return nil, err
	} else if user != nil && user.ID != oldUser.ID {
		return nil, entities.ErrUserEmailConflict
	}

	if user, err := u.p.SelectTaskByAuthorName(user.Author_name); err != nil {
		return nil, err
	} else if user != nil && user.ID != oldUser.ID {
		return nil, entities.ErrUserNameConflict
	}
	fmt.Println(oldUser.Status + " " + user.Status)

	if oldUser.Status != user.Status {
		if oldUser.Status == "new" && (user.Status != "in progress" && user.Status != "done") {
			return nil, entities.ErrUserStatusConflict
		}
		if oldUser.Status == "in progress" && (user.Status != "done") {
			return nil, entities.ErrUserStatusConflict
		}
	}

	updatedUser, err := u.p.UpdateTaskByID(id, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *Usecase) DeleteTaskByID(id int) error {
	if err := u.p.DeleteTaskByID(id); err != nil {
		return err
	}

	return nil
}

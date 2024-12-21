package provider

import (
	"database/sql"
	"errors"

	"web-rk2/internal/entities"
)

func (p *Provider) InsertTask(user entities.Task) (*entities.Task, error) {
	var id int

	err := p.conn.QueryRow(`INSERT INTO "task" (author_name, assignee_name, created, resolved, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`, user.Author_name, user.Assignee_name, user.Created, user.Resolved, user.Status).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &entities.Task{
		ID:            id,
		Assignee_name: user.Assignee_name,
		Author_name:   user.Author_name,
		Created:       user.Created,
		Resolved:      user.Resolved,
		Status:        user.Status,
	}, nil
}

func (p *Provider) SelectAllTasks() ([]*entities.Task, error) {
	users := []*entities.Task{}

	rows, err := p.conn.Query(`SELECT id, author_name, assignee_name, created, resolved, status FROM "task"`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, nil
		}
		return nil, err
	}

	for rows.Next() {
		var user entities.Task
		if err := rows.Scan(&user.ID, &user.Assignee_name, &user.Author_name, &user.Created, &user.Resolved, &user.Status); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (p *Provider) SelectTaskByID(id int) (*entities.Task, error) {
	var user entities.Task
	err := p.conn.QueryRow(`SELECT id, assignee_name, author_name, created, resolved, status FROM "task" WHERE id = $1 LIMIT 1`, id).
		Scan(&user.ID, &user.Assignee_name, &user.Author_name, &user.Created, &user.Resolved, &user.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (p *Provider) SelectTaskByAssigneeName(name string) (*entities.Task, error) {
	var user entities.Task
	err := p.conn.QueryRow(`SELECT id, assignee_name, author_name, created, resolved, status FROM "task" WHERE assignee_name = $1 LIMIT 1`, name).
		Scan(&user.ID, &user.Assignee_name, &user.Author_name, &user.Created, &user.Resolved, &user.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (p *Provider) SelectTaskByAuthorName(author_name string) (*entities.Task, error) {
	var user entities.Task
	err := p.conn.QueryRow(`SELECT id, assignee_name, author_name, created, resolved, status FROM "task" WHERE author_name = $1 LIMIT 1`, author_name).
		Scan(&user.ID, &user.Assignee_name, &user.Author_name, &user.Created, &user.Resolved, &user.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (p *Provider) UpdateTaskByID(id int, user entities.Task) (*entities.Task, error) {
	var updatedUser entities.Task
	err := p.conn.QueryRow(`UPDATE "task" SET author_name = $1, assignee_name = $2, created = $3, resolved = $4, status = $5 WHERE id = $6 RETURNING id, author_name, assignee_name, created, resolved`,
		user.Author_name, user.Assignee_name, user.Created, user.Resolved, user.Status, id).
		Scan(&updatedUser.ID, &updatedUser.Assignee_name, &updatedUser.Author_name, &updatedUser.Created, &updatedUser.Resolved, &updatedUser.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (p *Provider) DeleteTaskByID(id int) error {
	_, err := p.conn.Exec(`DELETE FROM "task" WHERE id = $1`,
		id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.ErrUserNotFound
		}
		return err
	}

	return nil
}

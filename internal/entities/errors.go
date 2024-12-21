package entities

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("task already exists")
	ErrUserNameConflict   = errors.New("task name conflict")
	ErrUserEmailConflict  = errors.New("task author_name conflict")
	ErrUserStatusConflict = errors.New("task status conflict")

	ErrUserNotFound = errors.New("task not found")
)

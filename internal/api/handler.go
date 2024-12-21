package api

import (
	"errors"
	"net/http"
	"strconv"

	"web-rk2/internal/entities"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetTask(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid id")
	}

	user, err := s.uc.GetTaskByID(id)
	if err != nil {
		if errors.Is(err, entities.ErrUserNotFound) {
			return e.String(http.StatusBadRequest, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, user)
}

func (s *Server) ListTasks(e echo.Context) error {
	users, err := s.uc.ListTasks()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, users)
}

func (s *Server) CreateTask(e echo.Context) error {
	var task entities.Task

	err := e.Bind(&task)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	err = validator.New().Struct(task)
	if err != nil {
		return e.String(http.StatusUnprocessableEntity, err.Error())
	}

	createdUser, err := s.uc.CreateTask(task)
	if err != nil {
		if errors.Is(err, entities.ErrUserNameConflict) ||
			errors.Is(err, entities.ErrUserEmailConflict) ||
			errors.Is(err, entities.ErrUserAlreadyExists) {
			return e.String(http.StatusConflict, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, createdUser)
}

func (s *Server) UpdateTask(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid id")
	}

	var user entities.Task

	err = e.Bind(&user)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	err = validator.New().Struct(user)
	if err != nil {
		return e.String(http.StatusUnprocessableEntity, err.Error())
	}

	updateUser, err := s.uc.UpdateTaskByID(id, user)
	if err != nil {
		if errors.Is(err, entities.ErrUserNameConflict) ||
			errors.Is(err, entities.ErrUserEmailConflict) ||
			errors.Is(err, entities.ErrUserAlreadyExists) {
			return e.String(http.StatusConflict, err.Error())
		}
		if errors.Is(err, entities.ErrUserNotFound) {
			return e.String(http.StatusBadRequest, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, updateUser)
}

func (s *Server) DeleteTask(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid id")
	}

	err = s.uc.DeleteTaskByID(id)
	if err != nil {
		if errors.Is(err, entities.ErrUserNotFound) {
			return e.String(http.StatusBadRequest, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusOK, "OK")
}

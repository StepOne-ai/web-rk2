package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server  *echo.Echo
	address string

	uc Usecase
}

func NewServer(ip string, port int, uc Usecase) *Server {
	api := Server{
		uc: uc,
	}

	api.server = echo.New()
	api.server.POST("/tasks", api.CreateTask)
	api.server.GET("/tasks", api.ListTasks)
	api.server.GET("/tasks/:id", api.GetTask)
	api.server.PUT("/tasks/:id", api.UpdateTask)
	api.server.DELETE("/tasks/:id", api.DeleteTask)

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (s *Server) Run() {
	s.server.Logger.Fatal(s.server.Start(s.address))
}

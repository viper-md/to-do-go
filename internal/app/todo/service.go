package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Service core
type Service struct {
	Core Core
}

// NewService core return
func NewService(core *Core) *Service {
	return &Service{
		Core: *core,
	}
}

// Create abc
func (s *Service) Create(c *gin.Context) {
	newTodo, err := s.Core.Create(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, newTodo)
	}
}

// GetAllTodos service
func (s *Service) GetAllTodos(c *gin.Context) {
	todos, err := s.Core.GetAllTodos(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

// GetToDo service
func (s *Service) GetToDo(c *gin.Context) {
	todo, err := s.Core.GetToDo(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// UpdateATodo service
func (s *Service) UpdateATodo(c *gin.Context) {
	todo, err := s.Core.UpdateATodo(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// DeleteATodo service
func (s *Service) DeleteATodo(c *gin.Context) {
	todo, err := s.Core.DeleteATodo(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

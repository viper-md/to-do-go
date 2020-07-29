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
	_, err := s.Core.Create(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Todo{})
	}
}

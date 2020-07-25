package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTodos fetch todos
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "working"})
}

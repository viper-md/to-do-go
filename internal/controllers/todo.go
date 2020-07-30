package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mukul-dev/to-do-app/internal/app/todo"
)

// define repo core and service
var (
	repo    = todo.NewRepo()
	core    = todo.NewCore(repo)
	service = todo.NewService(core)
)

// GetTodos fetch todos
func GetTodos(c *gin.Context) {
	service.GetAllTodos(c)
}

// CreateTodo new
func CreateTodo(c *gin.Context) {
	service.Create(c)
}

// GetATodo controllers
func GetATodo(c *gin.Context) {
	service.GetToDo(c)
}

// UpdateATodo save
func UpdateATodo(c *gin.Context) {
	service.UpdateATodo(c)
}

// DeleteATodo by id
func DeleteATodo(c *gin.Context) {
	service.DeleteATodo(c)
}

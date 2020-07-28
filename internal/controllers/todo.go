package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mukul-dev/to-do-app/internal/app/todo"
)

// GetTodos fetch todos
func GetTodos(c *gin.Context) {
	var t []todo.Todo
	err := todo.GetAllTodos(&t)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, t)
	}
}

// CreateTodo new
func CreateTodo(c *gin.Context) {
	var t todo.Todo
	c.BindJSON(&t)
	fmt.Println(t)
	err := todo.CreateToDo(&t)
	fmt.Println(t)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, t)
	}
}

// GetATodo controllers
func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var t todo.Todo
	err := todo.GetToDo(&t, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, t)
	}
}

// UpdateATodo save
func UpdateATodo(c *gin.Context) {
	var t todo.Todo
	id := c.Params.ByName("id")
	err := todo.GetToDo(&t, id)
	if err != nil {
		c.JSON(http.StatusNotFound, t)
	}
	c.BindJSON(&t)
	err = todo.UpdateTodo(&t, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, t)
	}
}

// DeleteATodo by id
func DeleteATodo(c *gin.Context) {
	var t todo.Todo
	id := c.Params.ByName("id")
	err := todo.DeleteToDo(&t, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}

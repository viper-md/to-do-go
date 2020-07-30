package todo

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Core feature
type Core struct {
	Repo Repo
}

// NewCore abc
func NewCore(repo *Repo) *Core {
	return &Core{
		Repo: *repo,
	}
}

// Create new
func (c *Core) Create(g *gin.Context) (Todo, error) {
	var t Todo
	g.BindJSON(&t)
	fmt.Println(t)
	e := c.Repo.Create(&t)
	if e != nil {
		return Todo{}, e
	}
	return t, nil
}

// GetAllTodos core
func (c *Core) GetAllTodos(g *gin.Context) ([]Todo, error) {
	var t []Todo
	err := c.Repo.GetAllTodos(&t)
	if err != nil {
		return []Todo{}, err
	}
	return t, nil
}

// GetToDo core
func (c *Core) GetToDo(g *gin.Context) (Todo, error) {
	id := g.Params.ByName("id")
	var t Todo
	err := c.Repo.GetToDo(&t, id)
	if err != nil {
		return Todo{}, err
	}
	return t, nil
}

// UpdateATodo core
func (c *Core) UpdateATodo(g *gin.Context) (Todo, error) {
	var t Todo
	id := g.Params.ByName("id")
	err := c.Repo.GetToDo(&t, id)
	if err != nil {
		return Todo{}, err
	}
	return t, nil
}

// DeleteATodo core
func (c *Core) DeleteATodo(g *gin.Context) (Todo, error) {
	var t Todo
	id := g.Params.ByName("id")
	err := c.Repo.DeleteToDo(&t, id)
	if err != nil {
		return Todo{}, err
	}
	return t, nil
}

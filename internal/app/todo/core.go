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
func (c *Core) Create(g *gin.Context) (int, error) {
	var t Todo
	g.BindJSON(&t)
	fmt.Println(t)
	e := c.Repo.Create(&t)
	if e != nil {
		return 0, e
	}
	return 200, nil
}

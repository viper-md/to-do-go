package todo

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/mukul-dev/to-do-app/pkg/database"
)

// GetAllTodos getall
func GetAllTodos(todo *[]Todo) (err error) {
	if err = database.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// CreateATodo todo create
func CreateATodo(todo *Todo) (err error) {
	if err = database.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

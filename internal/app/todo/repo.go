package todo

import (
	_ "github.com/go-sql-driver/mysql" // sql dialect
	"github.com/mukul-dev/to-do-app/pkg/database"
)

// GetAllTodos list
func GetAllTodos(todo *[]Todo) (err error) {
	if err := database.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// CreateToDo create new
func CreateToDo(todo *Todo) (err error) {
	if err := database.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

//GetToDo by id
func GetToDo(todo *Todo, id string) (err error) {
	if err := database.DB.Where("id = ? ", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTodo save
func UpdateTodo(todo *Todo, id string) (err error) {
	database.DB.Save(todo)
	return nil
}

// DeleteToDo by id
func DeleteToDo(todo *Todo, id string) (err error) {
	database.DB.Where("id = ? ", id).Delete(todo)
	return nil
}

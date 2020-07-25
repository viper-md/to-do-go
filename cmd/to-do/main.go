package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mukul-dev/to-do-app/internal/app/todo"
	"github.com/mukul-dev/to-do-app/internal/routes"
	"github.com/mukul-dev/to-do-app/pkg/database"
)

var err error

func main() {
	database.DB, err = gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("-----------------> database status", err)
	}
	defer database.DB.Close()
	database.DB.AutoMigrate(&todo.Todo{})
	r := routes.SetupRouter()
	// running
	r.Run()
}

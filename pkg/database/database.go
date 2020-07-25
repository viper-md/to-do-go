package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB Object
var DB *gorm.DB

// DBConfig structure
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig INIT
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		DBName:   "todoapp",
		User:     "root",
		Password: "mukul@007",
	}
	return &dbConfig
}

// DbURL returns
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

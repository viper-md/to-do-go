package connection

import "github.com/jinzhu/gorm"

// DB connection
type DB struct {
	db *gorm.DB
}

// NewDB return new db connection
func NewDB(db *gorm.DB) *DB {
	return &DB{db}
}

// DBC return
func (db *DB) DBC() *gorm.DB {
	return db.db
}

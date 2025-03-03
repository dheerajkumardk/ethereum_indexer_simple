package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// database object
var db *gorm.DB

func InitDatabase() {
	// 1. We have got GORM and DB drivers
	// 2. Connect to the database using Gorm's Open Method, Specifying db filepath
	db, err := gorm.Open(sqlite.Open("blocks.db"))
	if err != nil {
		panic("failed to connect to the database")
	}
	fmt.Println("Database connected!")
	// 3. Define models/structs for db schemas
	// 4. AutoMigrate - create or update DB schema based on our model
	db.AutoMigrate(&Block{})
	fmt.Println("Database Migrated.")
}
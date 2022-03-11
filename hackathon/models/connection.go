package models

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitBD() {

	//open a db connection
	var err error
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		//panic("failed to connect database")
		fmt.Println("Failed to connect database")
	}

	//Migrate the schema
	database.AutoMigrate(&Hackathon{})

	DB = database

}

package db

import (
	"log"
	"movie-notifier/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
var GlobalDB *gorm.DB


func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	// fmt.Println(db)
	if err != nil {
		log.Fatal(err)
	}

	GlobalDB = db

	return db

}


func MigrateDb(db *gorm.DB) {
	db.AutoMigrate(&entities.Tracker{})
}

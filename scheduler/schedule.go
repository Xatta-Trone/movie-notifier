package scheduler

import (
	"fmt"
	"movie-notifier/entities"
	"movie-notifier/scrapper"

	"gorm.io/gorm"
)

func RunScheduler(db *gorm.DB) {

	trackers := []entities.Tracker{}

	// save the keyword
	db.Where("is_parsed = ?", 0).Order("id desc").Find(&trackers)

	for _, tracker := range trackers {
		fmt.Println("scrapping...", tracker.Keyword)

		go scrapper.ScrapFreeDrive(tracker, db)
		go scrapper.ScrapMLSBD(tracker, db)
	}

}

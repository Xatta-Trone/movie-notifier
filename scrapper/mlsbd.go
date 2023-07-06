package scrapper

import (
	"fmt"
	"log"
	"movie-notifier/entities"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gen2brain/beeep"
	"gorm.io/gorm"
)

func ScrapMLSBD(model entities.Tracker, db *gorm.DB) {
	results := []string{}
	// Request the HTML page.
	res, err := http.Get(fmt.Sprintf("https://mlsbd.shop/??s=%s", model.Keyword))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".single-post").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		url := s.Find("a").First().AttrOr("href", "")

		results = append(results, url)

	})

	if len(results) > 0 {
		dataToUpdate := entities.Tracker{}
		db.First(&dataToUpdate,model.ID)
		// fmt.Print(dataToUpdate)
		dataToUpdate.IsParsed = true
		db.Save(&dataToUpdate)
		beeep.Notify("Movie found", fmt.Sprintf("Found the movie %s  on MLSBD", model.Keyword), "./img/icon.png")
	}

	fmt.Println("scrap results ===============")
	fmt.Println(results)
	fmt.Println("scrap results ===============")

}

package scrapper

import (
	"fmt"
	"log"
	"movie-notifier/entities"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
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

		// For each item found, get the title
		title := s.Find(".post-title").Text()
		title = strings.ToLower(title)
		title = strings.Trim(title,"")

		fmt.Println(title)


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

		url := os.Getenv("DISCORD_WEBHOOK_URL")

		if url == "" {
			url = "https://discord.com/api/webhooks/1126691142998691970/0CaBol5svSc8oe7UBhTdesfgDK34eh7ijrkQoQAcLSknwDe72PkxaWWSLxj4j3XnOWeC"
		}

		client, err := webhook.NewWithURL(url)

		if err != nil {
			return
		}

		msg := fmt.Sprintf("Found the movie %s  on MLSBD", model.Keyword)

		_, err = client.CreateMessage(discord.WebhookMessageCreate{
			Content: msg,
		})

		if err != nil {
			return
		}
	}

	fmt.Println("scrap results ===============")
	fmt.Println(results)
	fmt.Println("scrap results ===============")

}

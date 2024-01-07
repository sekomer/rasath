package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sekomer/rasath/api/db"
	"github.com/sekomer/rasath/api/models"
	"github.com/sekomer/rasath/controllers"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	return "0.0.0.0:" + port
}

func main() {
	db := db.Init()
	app := fiber.New()
	app.Use(logger.New())
	models.Setup(db)

	api := app.Group("/api/v1")

	controllers.RegisterEarthquakeRoutes(api, db)

	// register cronjob
	// c := cron.New()
	// err := c.AddFunc("@every 5s", func() {
	// 	earthquakes, _ := scraper.ScrapeData()
	// 	scraper.AddEarthquakesToDB(earthquakes, db)
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// c.Start()

	fmt.Println("Cronjob added")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK! || " + time.Now().String())
	})

	log.Fatal(app.Listen(getPort()))
}

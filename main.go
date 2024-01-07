package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kronborg6/Go-Fiber-Gorm/api/db"
	"github.com/kronborg6/Go-Fiber-Gorm/api/models"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hej")
	})

	log.Fatal(app.Listen(getPort()))
}

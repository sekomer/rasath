package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sekomer/rasath/api/models"
	"gorm.io/gorm"
)

func RegisterEarthquakeRoutes(app fiber.Router, db *gorm.DB) {
	earthquake := app.Group("/earthquake")

	// TODO: check if cache is valid before fetching from frontend
	earthquake.Get("/is-cache-valid", func(c *fiber.Ctx) error {
		return c.SendString("GET /earthquake")
	})

	earthquake.Get("/", func(c *fiber.Ctx) error {
		// get all earthquakes
		var earthquakes []models.Earthquake
		result := db.Find(&earthquakes)

		if result.Error != nil {
			return result.Error
		}

		return c.JSON(earthquakes)
	})
}

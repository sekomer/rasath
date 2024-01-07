package scraper

import (
	"github.com/sekomer/rasath/api/models"
)

// TODO: not finished
func HashEarthquake(e models.Earthquake) string {
	return e.Date
}

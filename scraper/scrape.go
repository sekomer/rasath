package scraper

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sekomer/rasath/api/models"
)

func ScrapeData() ([]models.Earthquake, error) {
	// Fetch the HTML content from the website
	resp, err := http.Get("http://www.koeri.boun.edu.tr/scripts/lst0.asp")
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	// Determine the character encoding of the response
	utf8Body, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, fmt.Errorf("error creating new reader: %w", err)
	}

	// Parse the HTML with correct encoding
	node, err := html.Parse(utf8Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %w", err)
	}

	// Traverse the DOM to find the <pre> tag
	var preContent string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "pre" {
			preContent = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)

	// Process the text content to create a slice of Earthquake structs
	lines := strings.Split(strings.TrimSpace(preContent), "\n")
	var earthquakes []models.Earthquake

	for _, line := range lines[1:] {
		parts := strings.Fields(line)
		if len(parts) < 10 {
			continue // Skip invalid lines
		}

		// if first part is not a YYYY.MM.DD date, skip
		if _, err := time.Parse("2006.01.02", parts[0]); err != nil {
			continue
		}

		quake := parseParts(parts)

		earthquakes = append(earthquakes, quake)
	}

	return earthquakes, nil
}

func AddEarthquakesToDB(earthquake []models.Earthquake, db *gorm.DB) (err error) {
	// Create or ignore on conflict
	results := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "date"}, {Name: "time"}, {Name: "latitude"}, {Name: "longitude"}},
		DoNothing: true,
	}).Create(&earthquake)

	// if there are errors, return
	if results.Error != nil {
		return results.Error
	}

	return
}

func CronTask(db *gorm.DB) {
	earthquakes, _ := ScrapeData()
	AddEarthquakesToDB(earthquakes, db)
}

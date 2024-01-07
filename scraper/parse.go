package scraper

import (
	"strconv"
	"strings"

	"github.com/sekomer/rasath/api/models"
)

func parseParts(parts []string) (quake models.Earthquake) {

	lat, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		// fmt.Println("Error:", err)
		lat = FloatParseError
	}

	lon, err := strconv.ParseFloat(parts[3], 64)
	if err != nil {
		// fmt.Println("Error:", err)
		lon = FloatParseError
	}

	depth, err := strconv.ParseFloat(parts[4], 32)
	if err != nil {
		// fmt.Println("Error:", err)
		depth = FloatParseError
	}

	mdFloat, err := strconv.ParseFloat(parts[5], 32)
	if err != nil {
		// fmt.Println("Error:", err)
		mdFloat = FloatParseError
	}
	mlFloat, err := strconv.ParseFloat(parts[6], 32)
	if err != nil {
		// fmt.Println("Error:", err)
		mlFloat = FloatParseError
	}
	mwFloat, err := strconv.ParseFloat(parts[7], 32)
	if err != nil {
		// fmt.Println("Error:", err)
		mwFloat = (FloatParseError)
	}

	quake = models.Earthquake{
		Date:      parts[0],
		Time:      parts[1],
		Latitude:  lat,
		Longitude: lon,
		Depth:     float32(depth),
		Location:  strings.Join(parts[8:], " "),
		Md:        float32(mdFloat),
		Ml:        float32(mlFloat),
		Mw:        float32(mwFloat),
	}

	return quake
}

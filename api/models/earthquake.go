package models

type Earthquake struct {
	ID        int     `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Date      string  `json:"date"`
	Time      string  `json:"time"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Depth     float32 `json:"depth"`
	Location  string  `json:"location"`
	Md        float32 `json:"md"`
	Ml        float32 `json:"ml"`
	Mw        float32 `json:"mw"`
}

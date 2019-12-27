package models

import "time"

type (
	Robot struct {
		Name      string      `json:"name"`
		Release   Release     `json:"release"`
		Ability   []string    `json:"ability"`
		Attribute []Attribute `json:"attribute"`
	}

	Release struct {
		Date     time.Time `json:"date"`
		Location string    `json:"location"`
		Version  float64   `json:"version"`
	}
	Attribute struct {
		Name  string `json:"name"`
		Power int    `json:"power"`
	}
)

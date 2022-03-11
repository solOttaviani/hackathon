package models

import "hackathon/service"

type (
	// Hackathon model of a each Hackathon
	Hackathon struct {
		ID           int             `json:"id" gorm:"primary_key"`
		Name         string          `json:"name"`
		Date         string          `json:"date"`
		Place        string          `json:"place"`
		Developments [10]Development `json:"developments"`
	}

	// Development model of all the developments from each hackathon
	Development struct {
		Name      string            `json:"name"`
		Developer service.Developer `json:"developer"`
		Place     int               `json:"winning_place"`
	}
)

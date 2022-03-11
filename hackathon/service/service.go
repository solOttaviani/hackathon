package service

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Developer model of a single developer
type Developer struct {
	Results struct {
		Name struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
	} `json:"results"`
	Info struct {
		Seed string `json:"seed"`
	} `json:"info"`
}

// GetDeveloper obteins a developer from the api randomuser
func GetDeveloper() (*Developer, error) {
	var developer Developer

	client := resty.New()

	_, err := client.R().
		SetResult(&developer).
		Get("https://randomuser.me/api")

	if err != nil {
		fmt.Println(err)
	}

	return &developer, nil
}

package controllers

import (
	"fmt"
	"hackathon/models"
	"hackathon/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type InputHackathon struct {
	Name         string              `json:"name" binding:"required`
	Date         string              `json:"date" binding:"required`
	Place        string              `json:"place" binding:"required`
	Developments []InputDevelopments `json:"developments" binding:"required`
}

type InputDevelopments struct {
	Name      string            `json:"name"`
	Developer service.Developer `json:"developer"`
	Place     int               `json:"winning-place"`
}

// createDeveloper gets a developer and inserts it in the bd
func CreateDeveloper(c *gin.Context) {
	developer, err := service.GetDeveloper()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "developer not found"})
		return
	}

	models.DB.Save(&developer)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Developer created successfully!", "resourceId": developer})
}

// createHackathon creates one hackathon
func CreateHackathon(c *gin.Context) {
	var hackathonInput InputHackathon
	if err := c.ShouldBindJSON(&hackathonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hackathon := models.Hackathon{Name: hackathonInput.Name, Date: hackathonInput.Date, Place: hackathonInput.Date}

	for i := 0; i < 10; i++ {
		developer, _ := service.GetDeveloper()

		for _, value := range hackathonInput.Developments {
			hackathon.Developments[i].Name = value.Name
			hackathon.Developments[i].Place = value.Place
		}

		hackathon.Developments[i].Developer.Results.Name.Title = developer.Results.Name.Title
		hackathon.Developments[i].Developer.Results.Name.First = developer.Results.Name.Title
		hackathon.Developments[i].Developer.Results.Name.Last = developer.Results.Name.Title
		hackathon.Developments[i].Developer.Info.Seed = developer.Info.Seed
	}

	models.DB.Save(&hackathon)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": hackathon})

}

// getHackathon obteins a hackathon by id
func GetHackathon(c *gin.Context) {
	var hackathon models.Hackathon
	hackathonID := c.Param("id")

	models.DB.First(&hackathon, hackathonID)

	if hackathon.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No hackathon found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": hackathon})
}

// getAllHackathons gets all the hackathons
func GetAllHackathons(c *gin.Context) {
	var hackathons []models.Hackathon

	models.DB.Find(&hackathons)

	if len(hackathons) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No data found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": hackathons})

}

// getDevelopers obteins all the developers that are in the bd
func GetDevelopers(c *gin.Context) {
	var developers []service.Developer

	models.DB.Find(&developers)

	if len(developers) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No data found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": developers})
}

func CronCreateHackathon() {
	fmt.Println("it showld create a hachathon every 5 min")
}

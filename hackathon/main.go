package main

import (
	"fmt"

	"hackathon/controllers"
	"hackathon/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/robfig/cron"
)

func main() {

	fmt.Println("Hello World")

	models.InitBD()
	c := cron.New()

	router := gin.Default()

	v1 := router.Group("/api/v1/hackathon")
	{
		v1.POST("/", controllers.CreateHackathon)
		v1.GET("/", controllers.GetAllHackathons)
		v1.GET("/:id", controllers.GetHackathon)
		v1.GET("/devs", controllers.GetDevelopers)
	}
	router.Run()

	c.AddFunc("@every 5m", controllers.CronCreateHackathon)
	c.Start()

}

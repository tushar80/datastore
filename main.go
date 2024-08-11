package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/tushar80/datastore/config"
	"github.com/tushar80/datastore/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	config.InitRedis()

	r := gin.Default()

	r.POST("/import", handlers.ImportExcel)
	r.GET("/view", handlers.ViewRecords)

	r.Run(":8080")
}

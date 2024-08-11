package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tushar80/datastore/services"
)

func ImportExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer f.Close()

	records, err := services.ParseExcel(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go services.StoreRecords(records)

	c.JSON(http.StatusOK, gin.H{"message": "File is being processed"})
}

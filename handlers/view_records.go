package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tushar80/datastore/config"
	"github.com/tushar80/datastore/services"
)

func ViewRecords(c *gin.Context) {
	var records []services.Record
	val, err := config.RDB.Get(config.Ctx, "records").Result()

	if err == nil {
		err := json.Unmarshal([]byte(val), &records)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse cache"})
			return
		}
		fmt.Println("Things in cache")
	} else {
		fmt.Println("Noting in cache")
		records = services.FetchFromDB()
		cacheData, _ := json.Marshal(records)
		config.RDB.Set(config.Ctx, "records", cacheData, 5*time.Minute)
	}

	c.JSON(http.StatusOK, records)
}

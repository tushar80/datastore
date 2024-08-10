package services

import (
	"encoding/json"
	"time"

	"github.com/tushar80/datastore/config"
)

func CacheRecords(records []Record) error {
	cacheData, err := json.Marshal(records)
	if err != nil {
		return err
	}
	return config.RDB.Set(config.Ctx, "records", cacheData, 5*time.Minute).Err()
}

func FetchFromCache() ([]Record, error) {
	var records []Record
	val, err := config.RDB.Get(config.Ctx, "records").Result()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(val), &records); err != nil {
		return nil, err
	}

	return records, nil
}

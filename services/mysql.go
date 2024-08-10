package services

import (
	"errors"

	"gorm.io/gorm"

	"github.com/tushar80/datastore/config"
)

func StoreRecords(records []Record) error {
	if err := config.DB.AutoMigrate(&Record{}); err != nil {
		return err
	}

	for _, record := range records {
		if err := config.DB.Create(&record).Error; err != nil {
			return err
		}
	}
	go CacheRecords(FetchFromDB())

	return nil
}

func FetchFromDB() []Record {
	var records []Record
	config.DB.Find(&records)
	return records
}

func UpdateRecord(id int, updatedRecord Record) error {
	var record Record
	if err := config.DB.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("record not found")
		}
		return err
	}

	record.FirstName = updatedRecord.FirstName
	record.LastName = updatedRecord.LastName
	record.CompanyName = updatedRecord.CompanyName
	record.Address = updatedRecord.Address
	record.City = updatedRecord.City
	record.County = updatedRecord.County
	record.Postal = updatedRecord.Postal
	record.Phone = updatedRecord.Phone
	record.Email = updatedRecord.Email
	record.Web = updatedRecord.Web

	if err := config.DB.Save(&record).Error; err != nil {
		return err
	}

	err := CacheRecords(FetchFromDB())
	if err != nil {
		return err
	}

	return nil
}

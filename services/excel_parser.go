package services

import (
	"errors"
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

type Record struct {
	ID          int    `gorm:"primaryKey"`
	FirstName   string `gorm:"type:varchar(255)"`
	LastName    string `gorm:"type:varchar(255)"`
	CompanyName string `gorm:"type:varchar(255)"`
	Address     string `gorm:"type:varchar(255)"`
	City        string `gorm:"type:varchar(255)"`
	County      string `gorm:"type:varchar(255)"`
	Postal      string `gorm:"type:varchar(50)"`
	Phone       string `gorm:"type:varchar(50)"`
	Email       string `gorm:"type:varchar(255)"`
	Web         string `gorm:"type:varchar(255)"`
}

func ParseExcel(f io.Reader) ([]Record, error) {
	excel, err := excelize.OpenReader(f)
	if err != nil {
		return nil, err
	}

	firstSheet := excel.WorkBook.Sheets.Sheet[0].Name
	rows, err := excel.GetRows(firstSheet)
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		return nil, errors.New("Excel file does not contain enough rows")
	}

	headers := rows[0]
	expectedHeaders := []string{"first_name", "last_name", "company_name", "address", "city", "county", "postal", "phone", "email", "web"}

	// Validate headers
	for i, header := range expectedHeaders {
		if headers[i] != header {
			return nil, errors.New("Invalid column header: " + header)
		}
	}

	var records []Record
	for i, row := range rows[1:] {
		if len(row) < len(expectedHeaders) {
			return nil, errors.New(fmt.Sprint("Invalid row data at row: ", i+1))
		}
		records = append(records, Record{
			FirstName:   row[0],
			LastName:    row[1],
			CompanyName: row[2],
			Address:     row[3],
			City:        row[4],
			County:      row[5],
			Postal:      row[6],
			Phone:       row[7],
			Email:       row[8],
			Web:         row[9],
		})
	}
	return records, nil
}

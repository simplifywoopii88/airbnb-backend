package database

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
)

type City string

const (
	CITY  string = "city"
	SEOUL string = "seoul"
	BUSAN string = "busan"
	GIMPO string = "gimpo"
)

func (c *City) Scan(value interface{}) error {
	*c = City(value.([]byte))
	return nil
}

func (c City) Value() (driver.Value, error) {
	return string(c), nil
}

func CreateCityType(db *gorm.DB) {
	query := fmt.Sprintf("SELECT 1 FROM pg_type WHERE typname = %s;", CITY)

	result := db.Exec(query)
	switch {
	case result.RowsAffected == 0:
		query = fmt.Sprintf("CREATE TYPE %s AS ENUM ('%s', '%s', '%s');", CITY, SEOUL, BUSAN, GIMPO)
		if err := db.Exec(query).Error; err != nil {
			db.Exec(fmt.Sprintf("DROP TYPE %s CASCADE;", CITY))
			db.Exec(query)
		}
	case result.Error != nil:
		panic(result.Error)
	}
}

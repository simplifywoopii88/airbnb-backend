package database

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
)

type Kind string

const (
	KIND   string = "kind"
	OCEAN  string = "ocean"
	FOREST string = "forest"
	ETC    string = "etc"
)

func (c *Kind) Scan(value interface{}) error {
	*c = Kind(value.([]byte))
	return nil
}

func (c Kind) Value() (driver.Value, error) {
	return string(c), nil
}

func CreateKindType(db *gorm.DB) {
	query := fmt.Sprintf("SELECT 1 FROM pg_type WHERE typname = %s;", KIND)

	result := db.Exec(query)
	switch {
	case result.RowsAffected == 0:
		query = fmt.Sprintf("CREATE TYPE %s AS ENUM ('%s', '%s', '%s');", KIND, OCEAN, FOREST, ETC)
		if err := db.Exec(query).Error; err != nil {
			db.Exec(fmt.Sprintf("DROP TYPE %s CASCADE;", KIND))
			db.Exec(query)
		}
	case result.Error != nil:
		panic(result.Error)
	}
}

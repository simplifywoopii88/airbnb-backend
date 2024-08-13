package database

import (
	"database/sql"
	"database/sql/driver"
)

type City string

const (
	KOREA City = "korea"
	ETC   City = "etc"
)

func (c *City) Scan(value interface{}) error {
	*c = City(value.([]byte))
	return nil
}

func (c City) Value() (driver.Value, error) {
	return string(c), nil
}

type Room struct {
	Id          uint           `json:"id" gorm:"primaryKey"`
	Country     string         `json:"country" gorm:"default:'korea'"`
	City        City           `json:"city"`
	Price       int            `json:"price"`
	Rooms       int            `json:"rooms"`
	Toilets     int            `json:"toilets"`
	Description sql.NullString `json:"description"`
	Address     sql.NullString `json:"address"`
	PetFriendly bool           `json:"pet_friendly" gorm:"default:false"`
	Kind        string         `json:"kind"`
	OwnerRefer  int            `json:"user_id"`
	Owner       User           `gorm:"foreignKey:OwnerRefer"`
}

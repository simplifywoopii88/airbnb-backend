package database

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type City struct {
	Korea string
	USA   string
	ETC   string
}

type JSON json.RawMessage

func (c *City) Scan(value interface{}) error {
	cityJSON, ok := value.(string)
	if !ok {
		return errors.New("unexpected type for city")
	}
	return json.Unmarshal([]byte(cityJSON), c)
}

func (c City) Value() (driver.Value, error) {
	cityJSON, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(cityJSON), nil
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

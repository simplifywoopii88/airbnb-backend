package database

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// create type city as enum (
//
//	'SEOUL',
//	'GIMPO',
//	'BUSAN'
//
// );

type City string
type Kind string

const (
	// psql type
	CITY  string = "city"
	KOREA City   = "korea"
	ETC   City   = "etc"

	// psql type
	KIND   string = "kind"
	OCEAN  string = "ocean"
	FOREST string = "forest"
)

func (c *City) Scan(value interface{}) error {
	*c = City(value.([]byte))
	return nil
}

func (c City) Value() (driver.Value, error) {
	return string(c), nil
}

func CreateType(db *gorm.DB) {
	var quries []string

	// add query
	query1 := fmt.Sprintf("SELECT 1 FROM pg_type WHERE typname = %s;", CITY)
	query2 := fmt.Sprintf("SELECT 1 FROM pg_type WHERE typname = %s;", KIND)

	quries = append(quries, query1, query2)

	for _, query := range quries {
		result := db.Exec(query)
		switch {
		case result.RowsAffected == 0:
			if err := db.Exec("CREATE TYPE car_type AS ENUM ('SEDAN', 'HATCHBACK', 'MINIVAN');").Error; err != nil {
				//db.Exec("ALTER TYPE car_type RENAME VALUE ('SEDAN', 'HATCHBACK', 'MINIVAN') to ('HAHA');")
				db.Exec("DROP TYPE car_type CASCADE;")
				db.Exec("CREATE TYPE car_type AS ENUM ('HOHOHO', 'HATCHBACK', 'MINIVAN');")
			}
		case result.Error != nil:
			panic(result.Error)
		}
	}
}

type Room struct {
	Id          uint    `json:"id" gorm:"primaryKey"`
	Country     string  `json:"country" gorm:"default:'korea'"`
	City        City    `json:"city" gorm:"type:city"`
	Price       int     `json:"price"`
	Rooms       int     `json:"rooms"`
	Toilets     int     `json:"toilets"`
	Description *string `json:"description"`
	Address     *string `json:"address"`
	PetFriendly bool    `json:"pet_friendly" gorm:"default:false"`
	Kind        string  `json:"kind"`
	OwnerRefer  int     `json:"user_id"`
	Owner       User    `gorm:"foreignKey:OwnerRefer"`
	// time recode
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoomSerializer struct {
	Id          uint    `json:"id"`
	Country     string  `json:"country"`
	City        City    `json:"city"`
	Price       int     `json:"price"`
	Rooms       int     `json:"rooms"`
	Toilets     int     `json:"toilets"`
	Description *string `json:"description"`
	Address     *string `json:"address"`
	PetFriendly bool    `json:"pet_friendly"`
	Kind        string  `json:"kind"`
	OwnerRefer  int     `json:"user_id"`
	// time recode
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *Room) Serialize() RoomSerializer {
	return RoomSerializer{
		Id:          r.Id,
		Country:     r.Country,
		City:        r.City,
		Price:       r.Price,
		Rooms:       r.Rooms,
		Toilets:     r.Toilets,
		Description: r.Description,
		Address:     r.Address,
		PetFriendly: r.PetFriendly,
		Kind:        r.Kind,
		OwnerRefer:  r.OwnerRefer,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

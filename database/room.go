package database

import (
	"time"

	"github.com/simplifywoopii88/airbnb-backend/dbtype"
)

type Room struct {
	Id          uint    `json:"id" gorm:"primaryKey"`
	Country     string  `json:"country" gorm:"default:'korea'"`
	City        dbtype.City    `json:"city" gorm:"type:city"`
	Price       int     `json:"price"`
	Rooms       int     `json:"rooms"`
	Toilets     int     `json:"toilets"`
	Description *string `json:"description"`
	Address     *string `json:"address"`
	PetFriendly bool    `json:"pet_friendly" gorm:"default:false"`
	Kind        dbtype.Kind    `json:"kind" gorm:"type:kind"`
	OwnerRefer  int     `json:"user_id"`
	Owner       User    `gorm:"foreignKey:OwnerRefer"`
	// time recode
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoomSerializer struct {
	Id          uint    `json:"id"`
	Country     string  `json:"country"`
	City        dbtype.City    `json:"city"`
	Price       int     `json:"price"`
	Rooms       int     `json:"rooms"`
	Toilets     int     `json:"toilets"`
	Description *string `json:"description"`
	Address     *string `json:"address"`
	PetFriendly bool    `json:"pet_friendly"`
	Kind        dbtype.Kind    `json:"kind"`
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

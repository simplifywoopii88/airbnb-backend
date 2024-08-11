package models

import "time"

// User database model
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email"`
	Name  string `json:"name"`

	IsHost   bool   `json:"is_host"`
	Gender   string `json:"gender"`
	Language string `json:"language"`
	Currency string `json:"currency"`

	// time recode
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserSerializer serializer
type UserSerializer struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`

	IsHost   bool   `json:"is_host"`
	Gender   string `json:"gender"`
	Language string `json:"language"`
	Currency string `json:"currency"`

	// time recode
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// serialize for User struct
func (u *User) Serialize() UserSerializer {
	return UserSerializer{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		IsHost:    u.IsHost,
		Gender:    u.Gender,
		Language:  u.Language,
		Currency:  u.Currency,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

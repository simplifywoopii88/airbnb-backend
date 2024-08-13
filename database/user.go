package database

import (
	"errors"
	"time"
)

// User database model
type User struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Email string  `json:"email"`
	Name  *string `json:"name"`

	IsHost   bool    `json:"is_host" gorm:"default:false"`
	Gender   *string `json:"gender"`
	Language string  `json:"language" gorm:"default:'kor'"`
	Currency string  `json:"currency" gorm:"default:'won'"`

	// time recode
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserSerializer serializer
type UserSerializer struct {
	ID    uint    `json:"id"`
	Email string  `json:"email"`
	Name  *string `json:"name"`

	IsHost   bool    `json:"is_host"`
	Gender   *string `json:"gender"`
	Language string  `json:"language"`
	Currency string  `json:"currency"`

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

// CheckDuplicatedEmail is Duplicated Email
func (u *User) CheckDuplicatedEmail() error {
	DB.Where("email = ?", u.Email).Find(u)

	if u.ID != 0 {
		return errors.New("email is already exists")
	}

	// 중복이 없음
	return nil

}

func (u *User) FindUser(id int) error {
	DB.Where("id = ?", id).Find(u)

	if u.ID == 0 {
		return errors.New("cannot find user")
	}

	return nil
}

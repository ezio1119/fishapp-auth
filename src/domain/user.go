package domain

import "time"

type User struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name" validate:"required,max=10"`
	Email             string    `json:"email" validate:"required,email"`
	Password          string    `json:"password" validate:"required,min=6,max=72" gorm:"-"`
	EncryptedPassword string    `json:"encrypted_password"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

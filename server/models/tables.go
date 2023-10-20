package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID  `gorm:"primary_key; type:uuid; default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name" binding:"required"`
	Surname   string     `json:"surname" binding:"required"`
	Email     string     `json:"email" binding:"required"`
	Password  string     `json:"password" binding:"required"`
}

type SignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

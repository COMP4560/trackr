package models

import (
	"time"
)

type User struct {
	ID         uint   `gorm:"primarykey"`
	Email      string `gorm:"uniqueIndex"`
	Password   string
	FirstName  string
	LastName   string
	IsVerified bool
	UpdatedAt  time.Time
	CreatedAt  time.Time

	MaxValues        int64
	MaxValueInterval int64
}

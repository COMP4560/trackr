package models

import (
	"time"
)

type Value struct {
	ID uint `gorm:"primarykey"`

	Value string

	FieldID uint
	Field   Field `gorm:"constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}

package models

import "time"

type Value struct {
	ID        uint `gorm:"primarykey"`
	Value     string
	CreatedAt time.Time `gorm:"index"`

	FieldID uint
	Field   Field `gorm:"constraint:OnDelete:CASCADE;"`
}

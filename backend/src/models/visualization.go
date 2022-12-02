package models

import "time"

type Visualization struct {
	ID        uint `gorm:"primarykey"`
	Metadata  string
	UpdatedAt time.Time
	CreatedAt time.Time

	FieldID uint
	Field   Field `gorm:"constraint:OnDelete:CASCADE;"`
}

package models

import "time"

type Visualization struct {
	ID       uint `gorm:"primarykey"`
	Metadata string

	FieldID uint
	Field   Field `gorm:"constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

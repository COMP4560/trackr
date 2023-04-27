package projects

import "time"

type Project struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	APIKey         string    `json:"apiKey"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	NumberOfFields int64     `json:"numberOfFields"`
}

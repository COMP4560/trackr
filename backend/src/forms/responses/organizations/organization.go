package organizations

import "time"

type Organization struct {
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	APIKey           string    `json:"apiKey"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	NumberOfProjects int64     `json:"numberOfProjects"`
}

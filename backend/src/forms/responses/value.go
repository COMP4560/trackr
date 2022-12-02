package responses

import "time"

type Value struct {
	ID        uint      `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

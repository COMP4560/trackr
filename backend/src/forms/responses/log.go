package responses

import "time"

type Log struct {
	Message     string `json:"message"`
	ProjectID   *uint  `json:"projectId"`
	ProjectName string `json:"projectName"`

	CreatedAt time.Time `json:"createdAt"`
}

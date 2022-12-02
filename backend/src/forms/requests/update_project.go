package requests

type UpdateProject struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ResetAPIKey bool   `json:"resetAPIKey"`
	Share       bool   `json:"share"`
}

package requests

type UpdateVisualization struct {
	ID       uint   `json:"id"`
	FieldID  uint   `json:"fieldId"`
	Metadata string `json:"metadata"`
}

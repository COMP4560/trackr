package requests

type AddVisualization struct {
	FieldID  uint   `json:"fieldId"`
	Metadata string `json:"metadata"`
}

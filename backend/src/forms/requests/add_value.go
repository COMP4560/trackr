package requests

type AddValue struct {
	APIKey  string `form:"apiKey" url:"apiKey"`
	FieldID uint   `form:"fieldId" url:"fieldId"`
	Value   string `form:"value" url:"value"`
}

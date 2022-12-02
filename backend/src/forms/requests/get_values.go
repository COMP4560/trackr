package requests

type GetValues struct {
	APIKey  string `form:"apiKey" url:"apiKey"`
	FieldID uint   `form:"fieldId" url:"fieldId"`
	Order   string `form:"order" url:"order"`
	Offset  int    `form:"offset" url:"offset"`
	Limit   int    `form:"limit" url:"limit"`
}

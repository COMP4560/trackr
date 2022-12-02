package requests

type AddField struct {
	ProjectID uint   `json:"projectId"`
	Name      string `json:"name"`
}

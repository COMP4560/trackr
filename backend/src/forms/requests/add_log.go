package requests

type AddLog struct {
	Message   string `json:"message"`
	ProjectID uint   `json:"projectId"`
	UserID    uint   `json:"userId"`
}

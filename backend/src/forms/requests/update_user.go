package requests

type UpdateUser struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

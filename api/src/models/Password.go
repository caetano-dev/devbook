package models

//Password represents the request format to reset the password
type Password struct {
	NewPassword     string `json:"newPassword"`
	CurrentPassword string `json:"currentPassword"`
}

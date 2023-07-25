package model

type SignupBody struct {
	Email           string `json:"email"`
	FullName        string `json:"fullname"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Address         string `json:"address"`
}

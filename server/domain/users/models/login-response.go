package models

type LoginResponse struct {
	User User   `json:"user"`
	OTP  string `json:"otp"`
}

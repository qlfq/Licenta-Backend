package users

import (
	"domain/users/models"
	web_socket "presenter/server/web-socket"
)

func createOTP(m *web_socket.Manager, user models.User) models.LoginResponse {
	otp := m.Otps.NewOTP()

	response := models.LoginResponse{
		User: user,
		OTP:  otp.Key,
	}

	return response
}

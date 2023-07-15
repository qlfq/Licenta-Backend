package web_socket

import "time"

type OTP struct {
	Key     string
	Created time.Time
}

type RetentionMap map[string]OTP

package web_socket

import (
	"context"
	"github.com/google/uuid"
	"time"
)

func NewRetentionMap(ctx context.Context, retentionPeriod time.Duration) RetentionMap {
	rm := make(RetentionMap)

	go rm.Retention(ctx, retentionPeriod)

	return rm
}

func (rm RetentionMap) NewOTP() OTP {
	o := OTP{
		Key:     uuid.NewString(),
		Created: time.Now(),
	}

	rm[o.Key] = o

	return o
}

func (rm RetentionMap) VerifyOTP(otp string) bool {
	if _, ok := rm[otp]; !ok {
		return false
	}

	delete(rm, otp)

	return true
}

func (rm RetentionMap) Retention(ctx context.Context, retentionPeriod time.Duration) {
	ticket := time.NewTimer(400 * time.Millisecond)

	for {
		select {
		case <-ticket.C:
			for _, otp := range rm {
				if otp.Created.Add(retentionPeriod).Before(time.Now()) {
					delete(rm, otp.Key)
				}
			}

		case <-ctx.Done():
			return
		}
	}
}

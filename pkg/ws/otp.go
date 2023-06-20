package ws

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type otp struct {
	key     string
	created time.Time
}

type retentionMap map[string]otp

func newRetentionMap(ctx context.Context, rtPeriod time.Duration) retentionMap {
	rm := make(retentionMap)

	go rm.retention(ctx, rtPeriod)
	return rm
}

func (rm retentionMap) newOTP() otp {
	o := otp{
		key:     uuid.NewString(),
		created: time.Now(),
	}

	rm[o.key] = o
	return o
}

func (rm retentionMap) verifyOTP(otp string) bool {
	if _, ok := rm[otp]; !ok {
		return false
	}
	delete(rm, otp)
	return true
}

func (rm retentionMap) retention(ctx context.Context, rtPeriod time.Duration) {
	ticker := time.NewTicker(400 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			for _, OTP := range rm {
				if OTP.created.Add(rtPeriod).Before(time.Now()) {
					delete(rm, OTP.key)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

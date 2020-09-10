package infrastructure

import (
	"backend-project/domain"
	"time"
)

func RetryMechanism(retry int) time.Duration {
	var backoff domain.Backoff
	
	if retry < 0 {
		return 1 * time.Millisecond
	}
	return time.Duration(backoff.Interval * retry) * time.Millisecond
}

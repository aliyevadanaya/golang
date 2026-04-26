package retry

import (
	"math/rand"
	"net/http"
	"time"
)

func IsRetryable(resp *http.Response, err error) bool {
	if err != nil {
		return true
	}

	if resp == nil {
		return false
	}

	switch resp.StatusCode {
	case 429, 500, 502, 503, 504:
		return true
	case 401, 404:
		return false
	default:
		return false
	}
}

func CalculateBackoff(attempt int, baseDelay time.Duration, maxDelay time.Duration) time.Duration {
	backoff := baseDelay * time.Duration(1<<attempt)

	if backoff > maxDelay {
		backoff = maxDelay
	}

	return time.Duration(rand.Int63n(int64(backoff)))
}

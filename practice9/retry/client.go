package retry

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	HttpClient *http.Client
	MaxRetries int
	BaseDelay  time.Duration
	MaxDelay   time.Duration
}

func (c *Client) ExecutePayment(ctx context.Context, url string) error {
	var err error

	for attempt := 0; attempt < c.MaxRetries; attempt++ {

		// проверка контекста
		if ctx.Err() != nil {
			return ctx.Err()
		}

		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, err := c.HttpClient.Do(req)

		if err == nil && resp.StatusCode == 200 {
			fmt.Println("Attempt", attempt+1, ": Success!")
			return nil
		}

		if !IsRetryable(resp, err) {
			return fmt.Errorf("non-retryable error")
		}

		if attempt == c.MaxRetries-1 {
			break
		}

		delay := CalculateBackoff(attempt, c.BaseDelay, c.MaxDelay)
		fmt.Printf("Attempt %d failed: waiting %v...\n", attempt+1, delay)

		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return err
}

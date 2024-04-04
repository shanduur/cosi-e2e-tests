package retry

import (
	"fmt"
	"time"
)

var (
	maxRetries   int64         = 10
	initialDelay time.Duration = 250 * time.Millisecond
	maxDelay     time.Duration = time.Duration(maxRetries) * initialDelay
)

func SetMaxRetries(r int64)           { maxRetries = r }
func SetInitialDelay(d time.Duration) { initialDelay = d }
func SetMaxDelay(d time.Duration)     { maxDelay = d }

func Do(fn func() error) error {
	retries := int64(0)

	for {
		err := fn()
		if err == nil {
			return nil
		}

		retries++
		if retries > maxRetries {
			return fmt.Errorf("Max retries exceeded, last error: %w", err)
		}

		delay := initialDelay * time.Duration(retries)
		if delay > maxDelay {
			delay = maxDelay
		}
		time.Sleep(delay)
	}
}

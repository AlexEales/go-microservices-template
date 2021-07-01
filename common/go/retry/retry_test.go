package retry

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"common/go/backoff"
)

func TestRetry(t *testing.T) {
	t.Run("error returned after retries exhausted", func(t *testing.T) {
		opts := &Opts{
			MaxAttempts: 5,
			Backoff:     &backoff.NoDelay{},
			IsRetryable: errorNotNil,
		}
		fn := func() error {
			return fmt.Errorf("retryable error")
		}

		retries, err := Do(context.Background(), fn, opts)
		require.Equal(t, opts.MaxAttempts-1, retries)
		require.EqualError(t, err, "retryable error")
	})

	t.Run("error retruned on non-retryable error", func(t *testing.T) {
		opts := &Opts{
			MaxAttempts: 5,
			Backoff:     &backoff.NoDelay{},
			IsRetryable: func(e error) bool {
				return false
			},
		}
		fn := func() error {
			return fmt.Errorf("non-retryable error")
		}

		retries, err := Do(context.Background(), fn, opts)
		require.Equal(t, 0, retries)
		require.EqualError(t, err, "non-retryable error")
	})

	t.Run("retries until success", func(t *testing.T) {
		opts := &Opts{
			MaxAttempts: 5,
			Backoff:     &backoff.NoDelay{},
			IsRetryable: errorNotNil,
		}
		currentTry := 0
		retriesUntilSuccess := 3
		fn := func() error {
			if currentTry < retriesUntilSuccess {
				currentTry++
				return fmt.Errorf("retryable error")
			}
			return nil
		}

		retries, err := Do(context.Background(), fn, opts)
		require.Equal(t, retriesUntilSuccess, retries)
		require.NoError(t, err)
	})

	t.Run("retries indefinitely with negative max retries config", func(t *testing.T) {})
}

package backoff

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNoDelayBackoff(t *testing.T) {
	t.Run("returns 0 every time", func(t *testing.T) {
		backoff := &NoDelay{}
		for i := 0; i < 10; i++ {
			backoffTime := backoff.Backoff(0)
			require.Equal(t, time.Duration(0), backoffTime)
		}
	})
}

func TestLinearBackoff(t *testing.T) {
	t.Run("returns linear value every time", func(t *testing.T) {
		backoff := &Linear{
			Delay: 100 * time.Millisecond,
		}
		for i := 0; i < 10; i++ {
			backoffTime := backoff.Backoff(0)
			require.Equal(t, 100*time.Millisecond, backoffTime)
		}
	})
}

func TestExponentialBackoff(t *testing.T) {
	t.Run("returns initial backoff on first retry", func(t *testing.T) {
		backoff := &Exponential{
			InitialDelay: 10 * time.Millisecond,
			BaseDelay:    100 * time.Millisecond,
			MaxDelay:     5 * time.Second,
			Factor:       2,
			Jitter:       0.2,
		}
		backoffTime := backoff.Backoff(0)
		require.Equal(t, 10*time.Millisecond, backoffTime)
	})

	t.Run("returns backoff with jitter on subsequent retries", func(t *testing.T) {
		backoff := &Exponential{
			InitialDelay: 10 * time.Millisecond,
			BaseDelay:    100 * time.Millisecond,
			MaxDelay:     5 * time.Second,
			Factor:       2,
			Jitter:       0.2,
		}
		backoffTime := backoff.Backoff(1)
		require.InDelta(t, 100, backoffTime.Milliseconds(), 100*calculateJitter(0.2))
		require.Less(t, backoffTime, 5*time.Second)
	})

	t.Run("returns max backoff when retries surpass max", func(t *testing.T) {
		backoff := &Exponential{
			InitialDelay: 10 * time.Millisecond,
			BaseDelay:    100 * time.Millisecond,
			MaxDelay:     5 * time.Second,
			Factor:       2,
			Jitter:       0.2,
		}
		backoffTime := backoff.Backoff(100)
		require.InDelta(t, 5, backoffTime.Seconds(), calculateJitter(0.2))
	})
}

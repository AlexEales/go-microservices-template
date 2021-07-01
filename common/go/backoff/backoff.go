package backoff

import (
	"math"
	"math/rand"
	"time"
)

type Strategy interface {
	// Returns the amount of time to backoff based on the number
	// of previous retries
	Backoff(retries int) time.Duration
}

type NoDelay struct{}

func (nd *NoDelay) Backoff(retries int) time.Duration {
	return 0
}

type Linear struct {
	Delay time.Duration
}

func DefaultLinear() *Linear {
	return &Linear{
		Delay: 100 * time.Millisecond,
	}
}

func (l *Linear) Backoff(retries int) time.Duration {
	return l.Delay
}

type Exponential struct {
	InitialDelay time.Duration
	BaseDelay    time.Duration
	MaxDelay     time.Duration
	Factor       float64
	Jitter       float64
}

func DefaultExponential() *Exponential {
	return &Exponential{
		InitialDelay: 10 * time.Millisecond,
		BaseDelay:    100 * time.Millisecond,
		MaxDelay:     5 * time.Second,
		Factor:       1.6,
		Jitter:       0.2,
	}
}

func (e *Exponential) Backoff(retries int) time.Duration {
	if retries == 0 {
		return e.InitialDelay
	}

	backoff := math.Min(
		float64(e.BaseDelay)*math.Pow(e.Factor, float64(retries)),
		float64(e.MaxDelay),
	)

	backoff = math.Max(
		0,
		backoff*calculateJitter(e.Jitter),
	)

	return time.Duration(backoff)
}

func calculateJitter(baseJitter float64) float64 {
	return 1 + baseJitter*(rand.Float64()*2-1)
}

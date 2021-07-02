package retry

import (
	"context"
	"log"
	"time"

	"go-microservices-template/common/go/backoff"
)

type Function func() error

type RetryableClassifierFn func(error) bool

var errorNotNil = func(err error) bool {
	return err != nil
}

type Opts struct {
	MaxAttempts int
	Backoff     backoff.Strategy
	IsRetryable RetryableClassifierFn
}

func DefaultOpts() *Opts {
	return &Opts{
		MaxAttempts: 5,
		Backoff:     backoff.DefaultExponential(),
		IsRetryable: errorNotNil,
	}
}

func Do(ctx context.Context, fn Function, opts *Opts) (retries int, err error) {
	if opts == nil {
		opts = DefaultOpts()
	}

	errorChan := make(chan error, 1)
	go func() {
		retries = 0
		for {
			if err = ctx.Err(); err != nil {
				break
			}

			if err = fn(); opts.IsRetryable(err) {
				if retries == opts.MaxAttempts-1 {
					break
				}
				boTime := opts.Backoff.Backoff(retries)
				log.Printf("Retrying, attempt %d/%d waiting %fs before retrying", retries+1, opts.MaxAttempts, boTime.Seconds())
				time.Sleep(boTime)
				retries++
			} else {
				errorChan <- err
				return
			}
		}
		errorChan <- err
	}()

	select {
	case <-ctx.Done():
		return retries, ctx.Err()
	case <-errorChan:
		return retries, err
	}
}

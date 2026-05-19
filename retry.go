package req

import (
	"time"
)

func defaultGetRetryInterval(resp *Response, attempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// RetryConditionFunc is a retry condition, which determines
// whether the request should retry.
type RetryConditionFunc func(resp *Response, err error) bool

// RetryHookFunc is a retry hook which will be executed before a retry.
type RetryHookFunc func(resp *Response, err error)

// GetRetryIntervalFunc is a function that determines how long should
// sleep between retry attempts.
type GetRetryIntervalFunc func(resp *Response, attempt int) time.Duration

func backoffInterval(min, max time.Duration) GetRetryIntervalFunc {
	_ = "STUB: not implemented"
	return *new(GetRetryIntervalFunc)
}

func newDefaultRetryOption() *retryOption { _ = "STUB: not implemented"; return nil }

type retryOption struct {
	MaxRetries       int
	GetRetryInterval GetRetryIntervalFunc
	RetryConditions  []RetryConditionFunc
	RetryHooks       []RetryHookFunc
}

func (ro *retryOption) Clone() *retryOption { _ = "STUB: not implemented"; return nil }

package request

import "time"

// Const vars for rate limiter
const (
	DefaultMaxRequestJobs   int32 = 50
	DefaultMaxRetryAttempts       = 3
	DefaultMutexLockTimeout       = 50 * time.Millisecond
	drainBodyLimit                = 100000
	proxyTLSTimeout               = 15 * time.Second
	userAgent                     = "User-Agent"
)

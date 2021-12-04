package request

import (
	"github.com/lestrrat-go/backoff/v2"
	"net/http"
)

// Vars for rate limiter
var (
	MaxRequestJobs   = DefaultMaxRequestJobs
	MaxRetryAttempts = DefaultMaxRetryAttempts
)

// Requester : is the struct for the request client
type Requester struct {
	HTTPClient         *http.Client
	Name               string
	UserAgent          string
	maxRetries         int
	jobs               int32
	disableRateLimiter int32
	backoff            backoff.Policy
	//Nonce              nonce.Nonce
	//limiter            Limiter
	//retryPolicy        RetryPolicy
	//timedLock *timedmutex.TimedMutex
}

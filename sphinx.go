package sphinx

import (
	"fmt"
	"github.com/Clever/leakybucket"
	"github.com/Clever/sphinx/common"
	"github.com/Clever/sphinx/matchers"
	"time"
)

type requestMatcher struct {
	Matches  []matchers.Matcher
	Excludes []matchers.Matcher
}

// Status contains the status of a limit.
type Status struct {
	Capacity  uint
	Reset     time.Time
	Remaining uint
	Name      string
}

// NewStatus returns the status of a limit.
func NewStatus(name string, bucket leakybucket.BucketState) Status {

	status := Status{
		Name:      name,
		Capacity:  bucket.Capacity,
		Reset:     bucket.Reset,
		Remaining: bucket.Remaining,
	}

	return status
}

// RateLimiter rate limits requests based on given configuration and limits.
type RateLimiter interface {
	Add(request common.Request) ([]Status, error)
	Configuration() Configuration
	Limits() []Limit
}

type sphinxRateLimiter struct {
	config Configuration
	limits []Limit
}

func (r *sphinxRateLimiter) Limits() []Limit {
	return r.limits
}

func (r *sphinxRateLimiter) Configuration() Configuration {
	return r.config
}

func (r *sphinxRateLimiter) Add(request common.Request) ([]Status, error) {
	status := []Status{}
	for _, limit := range r.Limits() {
		if !limit.Match(request) {
			continue
		}
		bucketstate, err := limit.Add(request)
		if err != nil {
			return status, fmt.Errorf("error while adding to Limit: %s. %s",
				limit.Name(), err.Error())
		}
		status = append(status, NewStatus(limit.Name(), bucketstate))
	}
	return status, nil
}

// NewRateLimiter returns a new RateLimiter based on the given configuration.
func NewRateLimiter(config Configuration) (RateLimiter, error) {

	rateLimiter := &sphinxRateLimiter{config: config, limits: config.Limits()}
	return rateLimiter, nil
}

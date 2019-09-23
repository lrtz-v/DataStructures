package circuit

import "time"

// DefaultTimeout is default timeout
const DefaultTimeout = time.Duration(60) * time.Second

// Settings configures CircuitBreaker
type Settings struct {
	// Name of CircuitBreaker
	Name string
	// MaxRequests is the maximum number of requests allowed to pass through when the CircuitBreaker is half-open.
	// If MaxRequests is 0, the CircuitBreaker allows only 1 request.
	MaxRequests uint32
	// Interval is the cyclic period of the closed state for the CircuitBreaker to clear the internal Counts.
	// If Interval is 0, the CircuitBreaker doesn't clear internal Counts during the closed state.
	Interval time.Duration
	// Timeout is the period of the open state, after which the state of the CircuitBreaker becomes half-open.
	// If Timeout is 0, the timeout value of the CircuitBreaker is set to 60 seconds.
	Timeout time.Duration
	// ReadyToTrip is called with a copy of Counts whenever a request fails in the closed state.
	// If ReadyToTrip returns true, the CircuitBreaker will be placed into the open state.
	// If ReadyToTrip is nil, default ReadyToTrip is used.
	ReadyToTrip func(counts Counts) bool
	// OnStateChange is called whenever the state of the CircuitBreaker changes.
	OnStateChange func(name string, from State, to State)
}

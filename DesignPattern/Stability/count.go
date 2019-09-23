package circuit

// Counts hold numbers of state
type Counts struct {
	Requests             uint32
	TotalSuccesses       uint32
	TotalFailures        uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailures  uint32
}

// OnRequest count requests
func (c *Counts) OnRequest() {
	c.Requests++
}

// OnSuccess count success times
func (c *Counts) OnSuccess() {
	c.TotalSuccesses++
	c.ConsecutiveSuccesses++
	c.ConsecutiveFailures = 0
}

// OnFailure count failure times
func (c *Counts) OnFailure() {
	c.TotalFailures++
	c.ConsecutiveFailures++
	c.ConsecutiveSuccesses = 0
}

// Clear reset counts
func (c *Counts) Clear() {
	c.Requests = 0
	c.TotalSuccesses = 0
	c.TotalFailures = 0
	c.ConsecutiveSuccesses = 0
	c.ConsecutiveFailures = 0
}

// DefaultReadyToTrip is default ConsecutiveFailures check
func DefaultReadyToTrip(counts Counts) bool {
	return counts.ConsecutiveFailures > 5
}

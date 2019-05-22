// Package gigasecond provides a very specialized time utility function
package gigasecond

import "time"

// AddGigasecond provides the time in 1e9 seconds for a given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}

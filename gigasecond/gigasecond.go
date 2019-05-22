// Package gigasecond provides a very specialized time utility function
package gigasecond

import "time"

const giga = 1000 * 1000 * 1000

// AddGigasecond provides the time in 1e9 seconds for a given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(giga * time.Second)
}

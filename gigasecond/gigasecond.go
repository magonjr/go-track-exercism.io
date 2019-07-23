package gigasecond

import "time"

const (
	gs = 1000000000 //A gigasecond is 10^9 (1,000,000,000) seconds.
)

// AddGigasecond calculate the moment when someone has lived for 10^9 seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * gs)
}

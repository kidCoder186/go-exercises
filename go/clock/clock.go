package clock

import "fmt"

// Clock structure
type Clock struct {
	hour   int
	minute int
}

// New constructs Clock object
func New(h, m int) Clock {
	return Clock{0, 0}.Add(h*60 + m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds amount 'minutes' to current clock.
func (c Clock) Add(minutes int) Clock {
	h := minutes / 60
	m := minutes - 60*h
	if c.minute+m < 0 {
		h--
	}
	return Clock{(c.hour + h + (c.minute+m)/60 + 24*abs(h)) % 24, (c.minute + m + abs(h)*60) % 60}
}

// Subtract substracts amount 'minutes' to current clock.
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

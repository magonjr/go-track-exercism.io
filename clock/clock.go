package clock

import (
	"fmt"
)

const minutesPerDay = 1440

//Clock struct contains the hours and minutes
type Clock struct {
	m int
}

//New creates a Clock stuct from given hours and minutes values
func New(hour, minute int) Clock {
	return Clock{}.Add(hour*60 + minute)
}

//String represents a clock struct as a string in the HH:MM format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}

//Add adds minutes to the current time in a Clock structure
func (c Clock) Add(minutes int) Clock {
	c.m += minutes
	c.m %= minutesPerDay

	if c.m < 0 {
		c.m += minutesPerDay
	}

	return c
}

//Subtract subtracts minutes from the current time in a Clock structure
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

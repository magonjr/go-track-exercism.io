package clock

import (
	"strconv"
)

type Clock struct {
	h, m int
}

func New(hour, minute int) Clock {
	return Clock{
		h: 0,
		m: 0,
	}.Add(hour*60 + minute)
}

func (c Clock) String() string {
	hours := strconv.Itoa(c.h)
	if c.h < 10 {
		hours = "0" + hours
	}
	min := strconv.Itoa(c.m)
	if c.m < 10 {
		min = "0" + min
	}
	return hours + ":" + min
}

func (c Clock) Add(minutes int) Clock {
	minutes += c.h*60 + c.m

	h := minutes / 60
	m := minutes % 60

	c.h = h % 24
	c.m = m % 60

	if minutes < 0 {
		c.m = (60 + c.m) % 60

		if c.m > 0 {
			c.h = c.h - 1
		}

		if c.h < 0 {
			c.h = 24 + c.h
		}
	}

	return c
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

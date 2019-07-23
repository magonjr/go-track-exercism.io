package raindrops

import "strconv"

//Convert func calcu randrops
func Convert(n int) string {
	var s string

	if n%3 == 0 {
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		return strconv.Itoa(n)
	}

	return s
}

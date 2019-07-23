package darts

import "math"

//Score returns the earned points in a single toss of a Darts game.
func Score(x, y float64) int {

	//calculate the distance from the center
	dist := math.Sqrt(x*x + y*y)

	//comparate to the radius
	switch {
	case dist > 10:
		return 0
	case dist > 5:
		return 1
	case dist > 1:
		return 5
	default:
		return 10
	}
}

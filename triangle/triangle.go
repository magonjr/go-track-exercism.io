package triangle

import "math"

//Kind  Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides tests triangle types
func KindFromSides(a, b, c float64) Kind {
	/*
	  For a shape to be a triangle at all, all sides have to be of length > 0, and
	  the sum of the lengths of any two sides must be greater than or equal to the
	  length of the third side.
	*/
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}

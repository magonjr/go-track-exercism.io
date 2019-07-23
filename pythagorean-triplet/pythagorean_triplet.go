package pythagorean

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	result := make([]Triplet, 0, max-min)
	for i := min; i <= max; i++ {
		for j := min; j <= i; j++ {
			for k := min; k <= j; k++ {
				if i*i == j*j+k*k {
					result = append(result, Triplet{k, j, i})
				}
			}
		}
	}

	return result
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	result := make([]Triplet, 0, p)
	for i := p; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			for k := j; k >= 1; k-- {
				if i*i == j*j+k*k && i+j+k == p {
					result = append(result, Triplet{k, j, i})
				}
			}
		}
	}
	return result
}

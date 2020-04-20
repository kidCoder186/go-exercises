package pythagorean

import "math"

// Triplet structure.
type Triplet [3]int

func (t Triplet) sum() int {
	return t[0] + t[1] + t[2]
}

//Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	res := []Triplet{}
	for i := min; i < max-1; i++ {
		for j := i + 1; j < max; j++ {
			var k int = int(math.Sqrt(float64(i*i + j*j)))
			if k > j && k <= max && k*k == i*i+j*j {
				res = append(res, Triplet{i, j, k})
			}
		}
	}
	return res
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	res := []Triplet{}
	allSet := Range(1, p)
	for _, set := range allSet {
		if set.sum() == p {
			res = append(res, set)
		}
	}
	return res
}

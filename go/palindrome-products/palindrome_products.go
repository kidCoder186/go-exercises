package palindrome

import (
	"errors"
	"math"
)

// Product structure
type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func reverseNumber(x int) int {
	res := 0
	for x > 0 {
		res = res*10 + x%10
		x = x / 10
	}
	return res
}

func isPalindrome(x int) bool {
	if x < 0 {
		x = -x
	}
	return reverseNumber(x) == x
}

// Products finds the largest and smallest palindromes
// which are products of numbers within the range fmin to fmax.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	min := math.MaxInt32
	max := -min
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if p < max {
				continue
			}
			if isPalindrome(p) {
				if p > max {
					max = p
					pmax.Product = max
					pmax.Factorizations = [][2]int{{i, j}}
				} else if p == max {
					pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
				}
			}
		}
	}

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if p > min {
				continue
			}
			if isPalindrome(p) {
				if p < min {
					min = p
					pmin.Product = min
					pmin.Factorizations = [][2]int{{i, j}}
				} else if p == min {
					pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
				}
			}
		}
	}

	if min == fmax*fmax+1 {
		err = errors.New("no palindromes")
		if fmin > fmax {
			err = errors.New("fmin > fmax")
		}
	}
	return
}

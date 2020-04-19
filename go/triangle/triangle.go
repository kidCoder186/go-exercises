// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Kind structure
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides checks 3 given numbers whether is triangle or not.
func KindFromSides(a, b, c float64) Kind {
	if (a+b >= c) && (a+c >= b) && (b+c >= a) && (a > 0) && (b > 0) && (c > 0) &&
		(a != math.Inf(1) && b != math.Inf(1) && c != math.Inf(1)) {
		if (a == b) && (b == c) {
			return Equ
		} else if (a == b) || (b == c) || (a == c) {
			return Iso
		}
		return Sca
	}
	return NaT
}

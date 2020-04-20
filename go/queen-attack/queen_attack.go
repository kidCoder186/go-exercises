package queenattack

import (
	"errors"
	"strings"
)

// Point structure
type Point struct {
	x, y int
}

func newPoint(pos string) (Point, error) {
	pos = strings.ToLower(pos)
	p := Point{0, 0}
	if len(pos) != 2 {
		return p, errors.New("length of queen position > 2")
	}
	if (pos[0] >= 'a' && pos[0] <= 'h') &&
		(pos[1] >= '1' && pos[1] <= '8') {
		return Point{int(pos[0]) - 'a', int(pos[1]) - '0'}, nil
	}
	return p, errors.New("out of range chess board")
}

func (p Point) isEqual(point Point) bool {
	return p.x == point.x && p.y == point.y
}

// CanQueenAttack checks whether or not 2 queens can attack each other.
func CanQueenAttack(pos1, pos2 string) (ok bool, err error) {
	q1, err1 := newPoint(pos1)
	q2, err2 := newPoint(pos2)
	if err1 != nil || err2 != nil || q1.isEqual(q2) {
		return false, errors.New("wrong queen position")
	}
	return (q1.x == q2.x) || (q1.y == q2.y) ||
		(q1.x-q1.y == q2.x-q2.y) || (q1.x+q1.y == q2.x+q2.y), nil
}

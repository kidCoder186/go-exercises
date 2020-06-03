package robot

import "fmt"

// definitions used in step 1

var Step1Robot struct {
	X, Y int
	Dir
}

type Point struct {
	X, Y int
}

type Dir int

func (d Dir) String() string {
	return fmt.Sprintf("%d", d)
}

const (
	N Dir = iota
	E
	S
	W
)

var _ fmt.Stringer = Dir(1729)

// additional definitions used in step 2

type Action byte
type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

// additional definition used in step 3

type Action3 struct {
	Name   string
	Script string
}
type Step3Robot struct {
	Name string
	Step2Robot
}

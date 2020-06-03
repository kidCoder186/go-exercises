package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Garden structure
type Garden struct {
	diagram        string
	children       []string
	sortedChildren []string
}

var (
	plantsMap = map[byte]string{
		'R': "radishes",
		'C': "clover",
		'G': "grass",
		'V': "violets",
	}
)

// NewGarden creates a new Garden object
func NewGarden(diagram string, children []string) (*Garden, error) {
	lines := strings.FieldsFunc(diagram, func(r rune) bool {
		return r == '\n'
	})
	numEOL := strings.FieldsFunc(diagram, func(r rune) bool {
		return r != '\n'
	})
	if len(lines) != 2 || len(lines[0]) != 2*len(children) ||
		len(numEOL) != 2 || len(lines[0]) != len(lines[1]) {
		return nil, errors.New("invalid input1")
	}
	for i := 0; i < len(lines[0]); i++ {
		if _, ok := plantsMap[lines[0][i]]; !ok {
			return nil, errors.New("invalid input2")
		}
		if _, ok := plantsMap[lines[1][i]]; !ok {
			return nil, errors.New("invalid input3")
		}
	}
	sortedChildren := append([]string(nil), children...)
	sort.Slice(sortedChildren, func(i, j int) bool {
		return strings.Compare(sortedChildren[i], sortedChildren[j]) < 0
	})
	for i := 0; i < len(sortedChildren)-1; i++ {
		if strings.Compare(sortedChildren[i], sortedChildren[i+1]) == 0 {
			return nil, errors.New("invalid input4")
		}
	}
	return &Garden{
		diagram:        diagram,
		children:       children,
		sortedChildren: sortedChildren,
	}, nil
}

// Plants returns all plants that are planted by a child
func (g *Garden) Plants(child string) ([]string, bool) {
	lines := strings.FieldsFunc(g.diagram, func(r rune) bool {
		return r == '\n'
	})
	res := []string{}
	pos := -1
	for i, v := range g.sortedChildren {
		if strings.Compare(v, child) == 0 {
			pos = i
			break
		}
	}
	if pos == -1 {
		return res, false
	}
	res = append(res, plantsMap[lines[0][2*pos]])
	res = append(res, plantsMap[lines[0][2*pos+1]])
	res = append(res, plantsMap[lines[1][2*pos]])
	res = append(res, plantsMap[lines[1][2*pos+1]])
	return res, true
}

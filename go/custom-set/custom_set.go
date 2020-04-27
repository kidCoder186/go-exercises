package stringset

import "strings"

// Set structure
type Set struct {
	exist map[string]bool
}

// New creates a Set object.
func New() Set {
	return Set{map[string]bool{}}
}

// NewFromSlice creates a Set object with initial data.
func NewFromSlice(s []string) Set {
	set := New()
	for _, v := range s {
		set.exist[v] = true
	}
	return set
}

func (s Set) String() string {
	res := []string{"{"}
	for k := range s.exist {
		res = append(res, "\"")
		res = append(res, k)
		res = append(res, "\"")
		res = append(res, ", ")
	}
	if len(s.exist) != 0 {
		res = res[:len(res)-1]
	}
	res = append(res, "}")
	return strings.Join(res, "")
}

// IsEmpty returns true if the set is empty.
func (s Set) IsEmpty() bool {
	return len(s.exist) == 0
}

// Has returns true if set contains `x` element.
func (s Set) Has(x string) bool {
	return s.exist[x]
}

// Subset returns true if s2 contains all elements of s1.
func Subset(s1, s2 Set) bool {
	for key := range s1.exist {
		if !s2.exist[key] {
			return false
		}
	}
	return true
}

// Disjoint returns true if they share no elements
func Disjoint(s1, s2 Set) bool {
	for key := range s1.exist {
		if s2.exist[key] {
			return false
		}
	}
	return true
}

// Equal returns true if s1 == s2
func Equal(s1, s2 Set) bool {
	if len(s1.exist) == len(s2.exist) {
		for k := range s1.exist {
			if !s2.exist[k] {
				return false
			}
		}
		return true
	}
	return false
}

// Add adds a element into set.
func (s Set) Add(x string) {
	s.exist[x] = true
}

// Intersection returns a set of all shared elements.
func Intersection(s1, s2 Set) Set {
	set := New()
	for k := range s1.exist {
		if s2.Has(k) {
			set.Add(k)
		}
	}
	return set
}

// Difference returns a set of all elements that are only in the s1 set.
func Difference(s1, s2 Set) Set {
	set := New()
	for k := range s1.exist {
		if !s2.Has(k) {
			set.Add(k)
		}
	}
	return set
}

// Union returns a set of all elements in either set.
func Union(s1, s2 Set) Set {
	set := New()
	for k := range s1.exist {
		set.Add(k)
	}
	for k := range s2.exist {
		set.Add(k)
	}
	return set
}

package sublist

// Relation type
type Relation string

const (
	SUBLIST   Relation = "sublist"
	EQUAL              = "equal"
	SUPERLIST          = "superlist"
	UNEQUAL            = "unequal"
)

// List structure
type List []int

func (l List) isEqual(list List) bool {
	if len(l) != len(list) {
		return false
	}
	for i, val := range l {
		if val != list[i] {
			return false
		}
	}
	return true
}

func (l List) isSublist(list List) bool {
	if len(l) >= len(list) {
		return false
	}
	if len(l) == 0 {
		return true
	}
	for i := 0; i <= len(list)-len(l); i++ {
		if l.isEqual(list[i : i+len(l)]) {
			return true
		}
	}
	return false
}

func (l List) isSuperlist(list List) bool {
	return list.isSublist(l)
}

// Sublist returns
func Sublist(list1, list2 []int) Relation {
	l1 := List(list1)
	l2 := List(list2)
	if l1.isSublist(l2) {
		return SUBLIST
	}
	if l1.isEqual(l2) {
		return EQUAL
	}
	if l1.isSuperlist(l2) {
		return SUPERLIST
	}
	return UNEQUAL
}

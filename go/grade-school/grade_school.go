package school

import (
	"sort"
	"strings"
)

// Grade structure
type Grade struct {
	grade    int
	students []string
}

// School structure
type School struct {
	grades map[int]*Grade
}

// New returns a new School object.
func New() *School {
	return &School{map[int]*Grade{}}
}

// Add adds a student into a grade.
func (s *School) Add(name string, id int) {
	if _, ok := s.grades[id]; !ok {
		newGrade := Grade{
			grade:    id,
			students: []string{name},
		}
		s.grades[id] = &newGrade
	} else {
		s.grades[id].students = append(s.grades[id].students, name)
	}
}

// Grade returns list of all students in the grade.
func (s *School) Grade(id int) []string {
	s.Sort()
	if _, ok := s.grades[id]; !ok {
		return []string{}
	}
	return s.grades[id].students
}

// Enrollment returns lists of all grades in the school.
func (s *School) Enrollment() []Grade {
	s.Sort()
	res := []Grade{}
	for i := range s.grades {
		res = append(res, *s.grades[i])
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].grade < res[j].grade
	})
	return res
}

func (s *School) Sort() {
	for i := range s.grades {
		sort.Slice(s.grades[i].students, func(j, k int) bool {
			return strings.Compare(s.grades[i].students[j], s.grades[i].students[k]) < 0
		})
	}
}

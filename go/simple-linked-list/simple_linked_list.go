package linkedlist

import (
	"errors"
)

// Element structure
type Element struct {
	data int
	next *Element
}

// List structure
type List struct {
	head *Element
	size int
}

// New creates a new List object.
func New(a []int) *List {
	if len(a) == 0 {
		return &List{}
	}
	list := List{
		head: &Element{
			data: a[0],
			next: nil,
		},
		size: 1,
	}
	cur := list.head
	for i := 1; i < len(a); i++ {
		cur.next = &Element{
			data: a[i],
			next: nil,
		}
		cur = cur.next
		list.size++
	}
	return &list
}

// Size returns size of linked list.
func (l *List) Size() int {
	if l == nil {
		return 0
	}
	return l.size
}

// Push adds a element into linked list.
func (l *List) Push(x int) {
	if l.Size() == 0 {
		l.head = &Element{x, nil}
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = &Element{
			data: x,
			next: nil,
		}
		l.size++
	}
}

// Pop removes the lastest element from linked list.
func (l *List) Pop() (int, error) {
	if l.Size() == 0 {
		return 0, errors.New("invalid remove from empty list")
	}
	var pre *Element
	cur := l.head
	for cur.next != nil {
		pre = cur
		cur = cur.next
	}
	res := cur.data
	if pre != nil {
		pre.next = nil
	} else {
		l = nil
	}
	return res, nil
}

// Array returns a list of all elements of linked list.
func (l *List) Array() []int {
	res := []int{}
	if l != nil {
		cur := l.head
		for cur != nil {
			res = append(res, cur.data)
			cur = cur.next
		}
	}
	return res
}

// Reverse returns a reversed list
func (l *List) Reverse() *List {
	if l != nil {
		var pre *Element
		cur := l.head
		for cur != nil {
			next := cur.next
			cur.next = pre
			pre = cur
			cur = next
		}
		l.head = pre
	}
	return l
}

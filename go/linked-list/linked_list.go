package linkedlist

import "errors"

var (
	// ErrEmptyList empty list error message
	ErrEmptyList = errors.New("empty list")
)

// Node structure
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// List structure
type List struct {
	head *Node
	tail *Node
	size int
}

// Next returns next node of the current node
func (e *Node) Next() *Node {
	if e != nil {
		return e.next
	}
	return nil
}

// Prev returns previous node of the current node
func (e *Node) Prev() *Node {
	if e != nil {
		return e.prev
	}
	return nil
}

// NewList creates a new object list
func NewList(args ...interface{}) *List {
	list := List{}
	for _, val := range args {
		list.PushBack(val)
	}
	return &list
}

// Size returns size of list
func (l *List) Size() int {
	return l.size
}

// PushFront adds a node at the top of list.
func (l *List) PushFront(v interface{}) {
	node := Node{
		Val: v,
	}
	if l.Size() == 0 {
		l.head = &node
		l.tail = &node
	} else {
		node.next = l.head
		l.head.prev = &node
		l.head = &node
	}
	l.size++
}

// PushBack adds a node at the end of list
func (l *List) PushBack(v interface{}) {
	node := Node{
		Val: v,
	}
	if l.Size() == 0 {
		l.head = &node
		l.tail = &node
	} else {
		node.prev = l.tail
		l.tail.next = &node
		l.tail = &node
	}
	l.size++
}

// PopFront takes a node from the top of list
func (l *List) PopFront() (interface{}, error) {
	if l.Size() == 0 {
		return nil, ErrEmptyList
	}
	res := l.head
	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
	l.size--
	return res.Val, nil
}

// PopBack takes a node from the end of list
func (l *List) PopBack() (interface{}, error) {
	if l.Size() == 0 {
		return nil, ErrEmptyList
	}
	res := l.tail
	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	l.size--
	return res.Val, nil
}

// Reverse reverses the list
func (l *List) Reverse() *List {
	if l.Size() > 1 {
		oldHead := l.head
		cur := l.head
		for cur.next != nil {
			next := cur.next
			cur.next = cur.prev
			cur.prev = next
			cur = next
		}
		cur.next = cur.prev
		cur.prev = nil
		l.head = cur
		l.tail = oldHead
	}
	return l
}

// First returns the first node of list
func (l *List) First() *Node {
	return l.head
}

// Last returns the last node of list
func (l *List) Last() *Node {
	return l.tail
}

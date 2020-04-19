package tree

import (
	"errors"
	"sort"
)

// Record structure
type Record struct {
	ID     int
	Parent int
}

// Node structure
type Node struct {
	ID       int
	Children []*Node
}

func newNode(id int) *Node {
	return &Node{id, nil}
}

// Build builds a tree of given records
func Build(records []Record) (*Node, error) {
	if len(records) < 1 {
		return nil, nil
	}
	nodes := make([]*Node, len(records))
	for _, record := range records {
		if record.ID >= len(records) {
			return nil, errors.New("out of index")
		}
		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New("root has parent")
		}
		if record.ID == record.Parent && record.ID != 0 {
			return nil, errors.New("non-root node has ID == Parent")
		}
		if record.Parent > record.ID {
			return nil, errors.New("parrent ID must be smaller than child ID")
		}
		if nodes[record.ID] == nil {
			nodes[record.ID] = newNode(record.ID)
		} else {
			return nil, errors.New("duplicate node")
		}
	}
	for _, record := range records {
		if nodes[record.Parent] != nil {
			if record.ID != record.Parent {
				if nodes[record.Parent].Children == nil {
					nodes[record.Parent].Children = []*Node{}
				}
				nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[record.ID])
			}
		}
	}
	for _, node := range nodes {
		if node != nil {
			if node.Children != nil {
				sort.Slice(node.Children, func(i, j int) bool {
					return node.Children[i].ID < node.Children[j].ID
				})
			}
		}
	}
	return nodes[0], nil
}

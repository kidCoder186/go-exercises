package pov

import (
	"strings"
)

// Node structure
type Node struct {
	label   string
	parent  string
	arcList []string
}

// Graph structure
type Graph struct {
	root string
	//	allArcs []string
	nodes map[string]*Node
}

// New creates a new Graph object
func New() *Graph {
	return &Graph{nodes: map[string]*Node{}}
}

// AddNode adds a node into current graph
func (g *Graph) AddNode(nodeLabel string) {
	newNode := &Node{
		label:   nodeLabel,
		arcList: []string{},
	}
	g.nodes[nodeLabel] = newNode
}

// AddArc adds a directed edge between `from` node and `to` node
func (g *Graph) AddArc(from, to string) {
	_, ok := g.nodes[from]
	if !ok {
		g.AddNode(from)
	}
	g.nodes[from].arcList = append(g.nodes[from].arcList, to)
	g.nodes[to].arcList = append(g.nodes[to].arcList, from)
	g.nodes[to].parent = from
}

func (g *Graph) findRootLable() string {
	if strings.Compare(g.root, "") == 0 {
		for _, node := range g.nodes {
			if strings.Compare(node.parent, "") == 0 {
				return node.label
			}
		}
	}
	return g.root
}

func buildArcList(g *Graph, parentLabel string, visit map[string]bool) []string {
	res := []string{}
	parent := g.nodes[parentLabel]
	visit[parentLabel] = true
	for _, childLabel := range parent.arcList {
		if !visit[childLabel] {
			res = append(res, parentLabel+" -> "+childLabel)
			res = append(res, buildArcList(g, childLabel, visit)...)
		}
	}
	return res
}

// ArcList returns a list that contains all arc of graph
func (g *Graph) ArcList() []string {
	rootLabel := g.findRootLable()
	visit := map[string]bool{}
	return buildArcList(g, rootLabel, visit)
}

// ChangeRoot changes root of graph
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	g.root = newRoot
	return g
}

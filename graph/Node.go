package graph

import (
    "github.com/AMan4Technology/DataStructure/useful/common"
)

func newNode(key string, value NodeValue) *Node {
	return &Node{
		key:       key,
		NodeValue: value,
		outDegree: make(map[string]*Edge),
	}
}

type Node struct {
	key string
	NodeValue
	outDegree map[string]*Edge
}

func (n *Node) Key() string {
	return n.key
}

func (n *Node) NumOfOut() int {
	return len(n.outDegree)
}

func (n *Node) Range(callback func(keyOfEnd string, edge *Edge) bool) {
	for end, edge := range n.outDegree {
		if !callback(end, edge) {
			break
		}
	}
}

type NodeValue = common.Value

package graph

import (
	"DataStructure/internal"
	"DataStructure/internal/example"
)

func EdgeValueOf(value float64) EdgeValue {
	return example.ValOf(value)
}

func newEdge(start, end *Node, value EdgeValue) *Edge {
	return &Edge{start, end, value}
}

type Edge struct {
	start, end *Node
	EdgeValue
}

func (e *Edge) Start() *Node {
	return e.start
}

func (e *Edge) End() *Node {
	return e.end
}

type EdgeValue = internal.Val

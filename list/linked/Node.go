package linked_list

import (
    "github.com/AMan4Technology/DataStructure/internal"
    "github.com/AMan4Technology/DataStructure/internal/example"
)

type Value = internal.Comparable

func NewNode(value Value) *Node {
    return &Node{Value: value}
}

func NodeOf(value float64) *Node {
    return NewNode(example.ComparableOf(value))
}

func ValueOf(value Value) float64 {
    return example.RealValueOf(value).Value("")
}

type Node struct {
    Value Value
    next  *Node
}

func (n *Node) Compare(b internal.Comparable) int8 {
    return n.Value.Compare(b.(*Node).Value)
}

func (n *Node) Next() *Node {
    return n.next
}

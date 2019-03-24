package astar

import (
    "math"

    "DataStructure/graph"
)

func NewNodeValue(x, y float64, value graph.NodeValue) graph.NodeValue {
    return &Node{x, y, value}
}

func NodeOf(value graph.NodeValue) *Node {
    return value.(*Node)
}

func DistanceOfNodes(a, b graph.NodeValue) float64 {
    return NodeOf(a).Distance(NodeOf(b))
}

type Node struct {
    X, Y float64
    graph.NodeValue
}

func (n *Node) Distance(other *Node) float64 {
    return math.Abs(n.X-other.X) + math.Abs(n.Y-other.Y)
}

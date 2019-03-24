package skiplist

import (
    "DataStructure/internal"
    "DataStructure/list/linked"
)

type Value = internal.Comparable

func newIndexNode(value Value, down *linkedlist.Node) *indexNode {
    return &indexNode{value: value, down: down}
}

func indexNodeOf(node *linkedlist.Node) *indexNode {
    return node.Value.(*indexNode)
}

type indexNode struct {
    value Value
    down  *linkedlist.Node
}

func (i *indexNode) Compare(b internal.Comparable) int8 {
    return i.value.Compare(b.(*indexNode).value)
}

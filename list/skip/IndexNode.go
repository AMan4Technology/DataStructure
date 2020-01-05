package skip_list

import (
    "github.com/AMan4Technology/DataStructure/internal"
    linked_list "github.com/AMan4Technology/DataStructure/list/linked"
)

type Value = internal.Comparable

func newIndexNode(value Value, down *linked_list.Node) *indexNode {
    return &indexNode{value: value, down: down}
}

func indexNodeOf(node *linked_list.Node) *indexNode {
    return node.Value.(*indexNode)
}

type indexNode struct {
    value Value
    down  *linked_list.Node
}

func (i *indexNode) Compare(b internal.Comparable) int8 {
    return i.value.Compare(b.(*indexNode).value)
}

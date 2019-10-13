package linkedqueue

import (
    "github.com/AMan4Technology/DataStructure/useful/common"

    "github.com/AMan4Technology/DataStructure/useful/compare"

    . "github.com/AMan4Technology/DataStructure/internal"
)

func newNodeValue(value common.Value) nodeValue {
    return nodeValue{value}
}

type nodeValue struct {
    common.Value
}

func (nodeValue) Compare(b Comparable) int8 {
    return compare.EqualTo
}

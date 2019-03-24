package linkedqueue

import (
    "DataStructure/useful/common"

    "DataStructure/useful/compare"

    . "DataStructure/internal"
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

package bplustree

import (
    common2 "DataStructure/useful/common"

    "DataStructure/internal"
)

type DataValue = string

type keys []key

func (k keys) Len() int {
    return len(k)
}

func (k keys) Compare(i int, value common2.Value) int8 {
    return k[i].Compare(value.(internal.Comparable))
}

type key internal.Comparable

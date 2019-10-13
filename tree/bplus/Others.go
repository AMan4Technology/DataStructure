package bplustree

import (
    "github.com/AMan4Technology/DataStructure/internal"
    "github.com/AMan4Technology/DataStructure/useful/common"
)

type DataValue = string

type keys []key

func (k keys) Len() int {
    return len(k)
}

func (k keys) Compare(i int, value common.Value) int8 {
    return k[i].Compare(value.(internal.Comparable))
}

type key internal.Comparable

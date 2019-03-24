package example

import (
    "DataStructure/internal"
    "DataStructure/useful/compare"
)

func ComparableValOf(val internal.Val) internal.ComparableVal {
    return value{val}
}

type value struct {
    internal.Val
}

func (v value) Compare(b internal.Comparable) int8 {
    return compare.Compare(v, b.(value), "")
}

package example

import "github.com/AMan4Technology/DataStructure/internal"

func ComparableOf(value float64) internal.Comparable {
    return ComparableValOf(ValOf(value))
}

func RealValueOf(x internal.Comparable) internal.Val {
    return x.(value).Val
}

package example

import "DataStructure/internal"

func ComparableOf(value float64) internal.Comparable {
    return ComparableValOf(ValOf(value))
}

func RealValueOf(x internal.Comparable) internal.Val {
    return x.(value).Val
}

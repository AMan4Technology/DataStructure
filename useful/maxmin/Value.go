package maxmin

import (
    "DataStructure/useful/compare"

    "DataStructure/internal"
)

func OfValues(base Value, others ...Value) (max, min Value) {
    max, min = base, base
    for _, value := range others {
        max = MaxOfTwoValues(max, value)
        min = MinOfTwoValues(min, value)
    }
    return
}

func MaxOfValues(base Value, others ...Value) (max Value) {
    max = base
    for _, value := range others {
        max = MaxOfTwoValues(max, value)
    }
    return
}

func MinOfValues(base Value, others ...Value) (min Value) {
    min = base
    for _, value := range others {
        min = MinOfTwoValues(min, value)
    }
    return
}

func MaxOfTwoValues(a, b Value) Value {
    if a.Compare(b) == compare.MoreThan {
        return a
    }
    return b
}

func MinOfTwoValues(a, b Value) Value {
    if a.Compare(b) == compare.LessThan {
        return a
    }
    return b
}

type Value = internal.Comparable

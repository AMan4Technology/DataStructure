package heap

import (
    "DataStructure/internal"
    "DataStructure/internal/example"
)

func NewValue(value float64) Value {
    return example.ComparableOf(value)
}

func ValueOf(value Value) float64 {
    return example.RealValueOf(value).Value("")
}

type Value = internal.Comparable

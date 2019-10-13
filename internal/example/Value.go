package example

import "github.com/AMan4Technology/DataStructure/internal"

func ValOf(value float64) internal.Val {
    return val(value)
}

type val float64

func (v val) Value(name string) float64 {
    return float64(v)
}

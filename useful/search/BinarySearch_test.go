package search

import (
    "github.com/AMan4Technology/DataStructure/useful/common"

    "github.com/AMan4Technology/DataStructure/internal"

    "github.com/AMan4Technology/DataStructure/internal/example"

    "fmt"
    "testing"
)

func TestBinarySearch(t *testing.T) {
    data := Values{example.ComparableOf(2), example.ComparableOf(2),
        example.ComparableOf(3), example.ComparableOf(3), example.ComparableOf(5),
        example.ComparableOf(8), example.ComparableOf(9)}
    fmt.Println(BinarySearch(data, example.ComparableOf(7)))
}

type Values []internal.Comparable

func (v Values) Len() int {
    return len(v)
}

func (v Values) Compare(i int, value common.Value) int8 {
    return v[i].Compare(value.(internal.Comparable))
}

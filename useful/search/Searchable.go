package search

import (
    "DataStructure/useful/common"
)

type Searchable interface {
    Len() int
    Compare(i int, value common.Value) int8
}

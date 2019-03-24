package internal

type Comparable interface {
    Compare(b Comparable) int8
}

package heap

import (
    "DataStructure/useful/compare"
)

func NewMax(capacity int) (heap Max) {
    heap = New(capacity)
    heap.up = func(length int, data []Value) {
        for curr := length - 1; data[(curr-1)/2].Compare(data[curr]) == compare.LessThan; curr = (curr - 1) / 2 {
            data[(curr-1)/2], data[curr] = data[curr], data[(curr-1)/2]
        }
    }
    heap.down = func(length int, data []Value) {
        for max, curr := 0, 0; ; curr = max {
            if 2*curr+1 < length && data[2*curr+1].Compare(data[max]) == compare.MoreThan {
                max = 2*curr + 1
            }
            if 2*curr+2 < length && data[2*curr+2].Compare(data[max]) == compare.MoreThan {
                max = 2*curr + 2
            }
            if max == curr {
                break
            }
            data[max], data[curr] = data[curr], data[max]
        }
    }
    return
}

type Max = Heap

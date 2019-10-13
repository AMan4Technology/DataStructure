package heap

import (
    "github.com/AMan4Technology/DataStructure/useful/compare"
)

func NewMin(capacity int) (heap Min) {
    heap = New(capacity)
    heap.up = func(length int, data []Value) {
        for curr := length - 1; data[(curr-1)/2].Compare(data[curr]) == compare.MoreThan; curr = (curr - 1) / 2 {
            data[(curr-1)/2], data[curr] = data[curr], data[(curr-1)/2]
        }
    }
    heap.down = func(length int, data []Value) {
        for min, curr := 0, 0; ; curr = min {
            if 2*curr+1 < length && data[2*curr+1].Compare(data[min]) == compare.LessThan {
                min = 2*curr + 1
            }
            if 2*curr+2 < length && data[2*curr+2].Compare(data[min]) == compare.LessThan {
                min = 2*curr + 2
            }
            if min == curr {
                break
            }
            data[min], data[curr] = data[curr], data[min]
        }
    }
    return
}

type Min = Heap

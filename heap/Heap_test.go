package heap

import (
    "fmt"
    "testing"
)

func TestMin(t *testing.T) {
    heap := NewMin(4)
    heap.Enqueue(NewValue(1), true)
    heap.Enqueue(NewValue(7), true)
    heap.Enqueue(NewValue(6), true)
    heap.Enqueue(NewValue(15), true)
    heap.Enqueue(NewValue(3), true)
    heap.Enqueue(NewValue(17), true)
    heap.Enqueue(NewValue(6), true)
    heap.Enqueue(NewValue(4), true)
    fmt.Println(ValueOf(heap.Top()))
    fmt.Println(ValueOf(heap.Dequeue()))
    fmt.Println(ValueOf(heap.Dequeue()))
    fmt.Println(ValueOf(heap.Dequeue()))
}

func TestMax(t *testing.T) {
    heap := NewMax(4)
    heap.Enqueue(NewValue(1), true)
    heap.Enqueue(NewValue(7), true)
    heap.Enqueue(NewValue(6), true)
    heap.Enqueue(NewValue(15), true)
    heap.Enqueue(NewValue(3), true)
    heap.Enqueue(NewValue(17), true)
    heap.Enqueue(NewValue(6), true)
    heap.Enqueue(NewValue(4), true)
    fmt.Println(ValueOf(heap.Top()))
    fmt.Println(ValueOf(heap.Dequeue()))
    fmt.Println(ValueOf(heap.Dequeue()))
    fmt.Println(ValueOf(heap.Dequeue()))
}

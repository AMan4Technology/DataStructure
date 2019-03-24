package heap

import (
    "errors"
)

func New(capacity int) Heap {
    return Heap{
        capacity: capacity,
        data:     make([]Value, 0, capacity),
    }
}

type Heap struct {
    capacity, length int
    data             []Value
    up, down         func(length int, data []Value)
}

func (h Heap) Empty() bool {
    return h.length == 0
}

func (h Heap) Full() bool {
    return h.length == h.capacity
}

func (h Heap) Top() Value {
    if h.Empty() {
        return nil
    }
    return h.data[0]
}

func (h *Heap) Enqueue(value Value, rampUp bool) error {
    if h.Full() {
        if !rampUp {
            return errors.New("the Heap is full")
        }
        h.data = append(h.data, value)
        h.capacity = cap(h.data)
    } else {
        h.data = append(h.data, value)
    }
    h.length++
    h.safeUp()
    return nil
}

func (h *Heap) Dequeue() (value Value) {
    if h.Empty() {
        return nil
    }
    value = h.data[0]
    h.data[0] = h.data[h.length-1]
    h.data = h.data[:h.length-1]
    h.length--
    h.safeDown()
    return
}

func (h Heap) Update(value Value) error {
    if h.Empty() {
        return errors.New("the Heap is empty")
    }
    h.data[0] = value
    h.safeDown()
    return nil
}

func (h Heap) safeUp() {
    if h.up != nil {
        h.up(h.length, h.data)
    }
}

func (h Heap) safeDown() {
    if h.down != nil {
        h.down(h.length, h.data)
    }
}

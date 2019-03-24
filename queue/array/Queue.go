package arrayqueue

import (
    "errors"

    "DataStructure/useful/common"
)

func New(capacity int) Queue {
    return Queue{
        capacity: capacity,
        data:     make([]common.Value, capacity),
    }
}

type Queue struct {
    capacity, length int
    head, tail       int
    data             []common.Value
}

func (q Queue) Empty() bool {
    return q.length == 0
}

func (q Queue) Full() bool {
    return q.length == q.capacity
}

func (q Queue) Cap() int {
    return q.capacity
}

func (q Queue) Len() int {
    return q.length
}

func (q Queue) Head() common.Value {
    if q.Empty() {
        return nil
    }
    return q.data[q.head]
}

func (q *Queue) Clean() {
    q.length, q.head, q.tail = 0, 0, 0
}

func (q *Queue) Enqueue(value common.Value, rampUp bool) error {
    if q.Full() {
        if !rampUp {
            return errors.New("the queue is full")
        }
        q.capacity *= 2
        queue := make([]common.Value, q.capacity)
        copy(queue, q.data[q.head:])
        copy(queue, q.data[:q.tail])
        q.head, q.tail, q.data = 0, q.capacity, queue
    }
    q.data[q.tail], q.tail = value, (q.tail+1)%q.capacity
    q.length++
    return nil
}

func (q *Queue) Dequeue() (value common.Value) {
    if q.Empty() {
        return nil
    }
    value = q.data[q.head]
    q.head = (q.head + 1) % q.capacity
    q.length--
    return
}

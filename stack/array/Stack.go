package arraystack

import (
    "errors"

    "github.com/AMan4Technology/DataStructure/useful/common"
)

func New(capacity int) Stack {
    return Stack{
        capacity: capacity,
        data:     make([]common.Value, 0, capacity),
    }
}

type Stack struct {
    capacity, length int
    data             []common.Value
}

func (s Stack) Empty() bool {
    return s.length == 0
}

func (s Stack) Full() bool {
    return s.length == s.capacity
}

func (s Stack) Capacity() int {
    return s.capacity
}

func (s Stack) Len() int {
    return s.length
}

func (s Stack) Top() common.Value {
    if s.Empty() {
        return nil
    }
    return s.data[s.length-1]
}

func (s *Stack) Push(value common.Value, rampUp bool) error {
    if s.Full() {
        if !rampUp {
            return errors.New("the stack is full")
        }
        s.data = append(s.data, value)
        s.capacity = cap(s.data)
    } else {
        s.data = append(s.data, value)
    }
    s.length++
    return nil
}

func (s *Stack) Pop() (value common.Value) {
    if s.Empty() {
        return nil
    }
    value = s.data[s.length-1]
    s.data = s.data[:s.length-1]
    s.length--
    return
}

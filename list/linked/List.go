package linkedlist

import (
    "DataStructure/useful/compare"
)

func New() List {
    head := NewNode(nil)
    return List{head: head, tail: head}
}

type List struct {
    length     int
    head, tail *Node
}

func (l List) Empty() bool {
    return l.length == 0
}

func (l List) Len() int {
    return l.length
}

func (l List) Head() *Node {
    return l.head
}

func (l List) Tail() *Node {
    return l.tail
}

func (l List) Node(index int) (curr *Node) {
    curr = l.node(index)
    if curr == nil {
        return nil
    }
    return
}

func (l List) Index(value Value) (index int, prev, curr *Node) {
    index, prev = l.index(NewNode(value), nil)
    if prev == nil {
        return -1, nil, nil
    }
    return index, prev, prev.next
}

func (l *List) Enqueue(node *Node) {
    l.InsertAfter(l.tail, node)
}

func (l *List) Prepend(node *Node) {
    l.InsertAfter(l.head, node)
}

func (l *List) Insert(index int, node *Node) {
    if index == l.length {
        l.Enqueue(node)
        return
    }
    if index == 0 {
        l.Prepend(node)
        return
    }
    prev := l.node(index - 1)
    if prev == nil {
        return
    }
    l.InsertAfter(prev, node)
}

func (l *List) Save(value Value) (node *Node) {
    node = NewNode(value)
    if value.Compare(l.tail.Value) != compare.LessThan {
        l.Enqueue(node)
        return
    }
    prev := l.head
    l.Range(func(index int, curr *Node) bool {
        if value.Compare(curr.Value) == compare.LessThan {
            l.InsertAfter(prev, node)
            return false
        }
        prev = curr
        return true
    })
    return
}

func (l *List) Dequeue() *Node {
    if l.Empty() {
        return nil
    }
    return l.DeleteAfter(l.head)
}

func (l *List) Pop() *Node {
    return l.Delete(l.length - 1)
}

func (l *List) Delete(index int) *Node {
    if index >= l.length {
        return nil
    }
    if index == 0 {
        return l.Dequeue()
    }
    prev := l.node(index - 1)
    if prev == nil {
        return nil
    }
    return l.DeleteAfter(prev)
}

func (l *List) Remove(value Value) (curr *Node) {
    _, prev := l.index(NewNode(value), nil)
    if prev != nil {
        curr = prev.next
        l.DeleteAfter(prev)
    }
    return
}

func (l List) Range(callback func(index int, curr *Node) bool) {
    for curr, i := l.head.next, 0; i < l.length && callback(i, curr); i++ {
        curr = curr.next
    }
}

func (l *List) InsertBefore(curr, node *Node) {
    _, prev := l.index(curr, func(base, other *Node) bool {
        return base == other
    })
    if prev != nil {
        l.InsertAfter(prev, node)
    }
}

func (l *List) InsertAfter(curr, node *Node) {
    l.insertAfter(curr, node)
    if curr == l.tail {
        l.tail = node
    }
}

func (l *List) DeleteBefore(curr *Node) *Node {
    _, prev := l.index(curr, func(base, other *Node) bool {
        return base == other
    })
    if prev == nil {
        return nil
    }
    return l.DeleteNode(prev)
}

func (l *List) DeleteAfter(curr *Node) (node *Node) {
    if l.Empty() || curr == l.tail {
        return nil
    }
    if curr.next == l.tail {
        l.tail = curr
    }
    node = curr.next
    curr.next = curr.next.next
    l.length--
    return
}

func (l *List) DeleteNode(curr *Node) *Node {
    _, prev := l.index(curr, func(base, other *Node) bool {
        return base == other
    })
    if prev == nil {
        return nil
    }
    return l.DeleteAfter(prev)
}

func (l List) node(index int) *Node {
    if index < 0 || l.length <= index {
        return nil
    }
    curr := l.head
    for i := 0; i <= index; i++ {
        curr = curr.next
    }
    return curr
}

func (l List) index(node *Node, equal func(base, other *Node) bool) (index int, prev *Node) {
    if equal == nil {
        equal = func(base, other *Node) bool {
            return base.Compare(other) == compare.EqualTo
        }
    }
    l.insertAfter(l.tail, node)
    for prev = l.head; !equal(node, prev.next); prev = prev.next {
        index++
    }
    l.DeleteAfter(l.tail)
    if index >= l.length {
        return -1, nil
    }
    return
}

func (l *List) insertAfter(curr, node *Node) {
    curr.next, node.next = node, curr.next
    l.length++
}

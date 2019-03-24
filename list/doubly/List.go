package doublylist

import (
    "DataStructure/useful/compare"
)

func New() (list List) {
    list = List{
        head: NewNode(nil),
        tail: NewNode(nil),
    }
    list.head.next, list.tail.prev = list.tail, list.head
    return
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
    index, curr = l.index(NewNode(value), nil)
    if curr == nil {
        return -1, nil, nil
    }
    return index, curr.prev, curr
}

func (l *List) Enqueue(node *Node) {
    l.InsertBefore(l.tail, node)
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
    if value.Compare(l.tail.prev.Value) != compare.LessThan {
        l.Enqueue(node)
        return
    }
    l.Range(func(index int, curr *Node) bool {
        if value.Compare(curr.Value) == compare.LessThan {
            l.InsertBefore(curr, node)
            return false
        }
        return true
    })
    return
}

func (l *List) Dequeue() *Node {
    return l.DeleteAfter(l.head)
}

func (l *List) Pop() *Node {
    return l.DeleteBefore(l.tail)
}

func (l *List) Delete(index int) *Node {
    curr := l.node(index)
    if curr == nil {
        return nil
    }
    return l.DeleteNode(curr)
}

func (l *List) Remove(value Value) (curr *Node) {
    _, curr = l.index(NewNode(value), nil)
    if curr != nil {
        l.DeleteNode(curr)
    }
    return
}

func (l List) Range(callback func(index int, curr *Node) bool) {
    for curr, i := l.head.next, 0; i < l.length && callback(i, curr); i++ {
        curr = curr.next
    }
}

func (l *List) InsertBefore(curr, next *Node) {
    l.InsertAfter(curr.prev, next)
}

func (l *List) InsertAfter(curr, next *Node) {
    curr.next, next.next, curr.next.prev, next.prev = next, curr.next, next, curr
    l.length++
}

func (l *List) DeleteBefore(curr *Node) *Node {
    return l.DeleteNode(curr.prev)
}

func (l *List) DeleteAfter(curr *Node) *Node {
    return l.DeleteNode(curr.next)
}

func (l *List) DeleteNode(curr *Node) *Node {
    if curr == l.head || curr == l.tail {
        return nil
    }
    curr.prev.next, curr.next.prev = curr.next, curr.prev
    l.length--
    return curr
}

func (l List) node(index int) *Node {
    if index < 0 {
        return nil
    }
    switch index * 2 / l.length {
    case 0:
        curr := l.head
        for i := 0; i <= index; i++ {
            curr = curr.next
        }
        return curr
    case 1:
        curr := l.tail
        for i := l.length - 1; i >= index; i-- {
            curr = curr.prev
        }
        return curr
    }
    return nil
}

func (l List) index(node *Node, equal func(base, other *Node) bool) (index int, curr *Node) {
    if equal == nil {
        equal = func(base, other *Node) bool {
            return base.Compare(other) == compare.EqualTo
        }
    }
    l.InsertBefore(l.tail, node)
    for curr = l.head.next; !equal(node, curr); curr = curr.next {
        index++
    }
    l.DeleteNode(node)
    if index >= l.length {
        return -1, nil
    }
    return
}

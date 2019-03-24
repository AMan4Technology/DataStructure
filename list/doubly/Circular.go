package doublylist

func NewCircular() (l Circular) {
    l = Circular{New()}
    l.head.prev, l.tail.next = l.tail, l.head
    return
}

type Circular struct {
    List
}

func (l Circular) Range(callback func(index int, curr *Node) bool) {
    for curr, i := l.head.next, 0; ; curr = curr.next {
        if curr == l.head || curr == l.tail {
            i = 0
            continue
        }
        if !callback(i, curr) {
            break
        }
        i++
    }
}

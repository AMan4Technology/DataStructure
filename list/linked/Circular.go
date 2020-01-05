package linked_list

func NewCircular() (l Circular) {
    l = Circular{New()}
    l.head.next = l.head
    return
}

type Circular struct {
    List
}

func (l Circular) Range(callback func(index int, curr *Node) bool) {
    for curr, i := l.head.next, 0; ; curr = curr.next {
        if curr == l.head {
            i = 0
            continue
        }
        if !callback(i, curr) {
            break
        }
        i++
    }
}

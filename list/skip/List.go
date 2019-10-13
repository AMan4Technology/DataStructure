package skiplist

import (
    "github.com/AMan4Technology/DataStructure/useful/compare"

    "github.com/AMan4Technology/DataStructure/list/linked"
)

func New(num int) (l List) {
    l = List{
        indexes: make([]linkedlist.List, 0, 1<<2),
        data:    linkedlist.New()}
    l.SetNum(num)
    return l
}

type List struct {
    num     int
    indexes []linkedlist.List
    data    linkedlist.List
}

func (l List) Num() int {
    return l.num
}

func (l *List) SetNum(num int) int {
    if num < 2 {
        num = 2
    }
    l.num = num
    return num
}

func (l List) MaxNum() int {
    return l.num + 1
}

func (l List) Empty() bool {
    return l.data.Empty()
}

func (l List) Len() int {
    return l.data.Len()
}

func (l List) Head() Value {
    if l.Empty() {
        return nil
    }
    return l.data.Head().Next().Value
}

func (l List) Tail() Value {
    if l.Empty() {
        return nil
    }
    return l.data.Tail().Value
}

func (l *List) Search(value Value) bool {
    if l.Empty() {
        return false
    }
    if l.Head().Compare(value) == compare.MoreThan {
        return false
    }
    if l.Tail().Compare(value) == compare.LessThan {
        return false
    }
    var (
        prev  = l.indexes[len(l.indexes)-1].Head().Next()
        start = prev
        end   *linkedlist.Node
    )
    for i := len(l.indexes) - 1; i >= 0; i-- {
        prev = findPrevOf(value, start, end)
        start, end = startEndOfPrev(prev)
    }
    for curr := start; curr != end; curr = curr.Next() {
        if value.Compare(curr.Value) == compare.EqualTo {
            return true
        }
    }
    return false
}

func (l *List) Save(value Value) {
    if l.Empty() {
        l.enqueueFirst(value)
        return
    }
    if value.Compare(l.Head()) == compare.LessThan {
        l.prepend(value)
        return
    }
    if value.Compare(l.Tail()) != compare.LessThan {
        l.enqueue(value)
        return
    }
    var (
        lenOfIndexes = len(l.indexes)
        prevNodes    = make([]*linkedlist.Node, lenOfIndexes)
        prev         = l.indexes[lenOfIndexes-1].Head().Next()
        start        = prev
        end          *linkedlist.Node
    )
    for i := lenOfIndexes - 1; i >= 0; i-- {
        prev = findPrevOf(value, start, end)
        prevNodes[i] = prev
        start, end = startEndOfPrev(prev)
    }
    prev = findPrevOf(value, start, end)
    l.data.InsertAfter(prev, linkedlist.NewNode(value))
    l.addIndex(lenOfIndexes, prevNodes)
}

func (l *List) Remove(value Value) {
    if l.Empty() || value.Compare(l.Head()) == compare.LessThan || value.Compare(l.Tail()) == compare.MoreThan {
        return
    }
    var (
        lenOfIndexes = len(l.indexes)
        prevNodes    = make([]*linkedlist.Node, lenOfIndexes)
        prev         = l.indexes[lenOfIndexes-1].Head().Next()
        start        = prev
        end          *linkedlist.Node
    )
    for i := lenOfIndexes - 1; i >= 0; i-- {
        prev = findPrevOf(value, start, end)
        prevNodes[i] = prev
        start, end = startEndOfPrev(prev)
    }
    var deletedNode *linkedlist.Node
    for curr := start; curr != end; curr = curr.Next() {
        if value.Compare(curr.Value) == compare.EqualTo {
            deletedNode = curr
            break
        }
        prev = curr
    }
    if deletedNode == nil {
        return
    }
    l.data.DeleteAfter(prev)
    if indexNodeOf(prevNodes[0]).down != deletedNode {
        return
    }
    var nextValue Value
    if deletedNode.Next() != nil {
        nextValue = deletedNode.Next().Value
    }
    for i := 0; i < lenOfIndexes; i++ {
        indexNode := indexNodeOf(prevNodes[i])
        if i > 0 && indexNode.down != prevNodes[i-1] {
            break
        }
        if indexNode.down != deletedNode {
            indexNode.value = nextValue
            continue
        }
        end = prevNodes[i].Next()
        if end != nil {
            end = indexNodeOf(end).down
        }
        if end == deletedNode.Next() {
            deletedNode = l.indexes[i].DeleteNode(prevNodes[i])
            continue
        }
        indexNode.value, indexNode.down = nextValue, deletedNode.Next()
    }
    if l.indexes[lenOfIndexes-1].Len() == 1 {
        l.indexes = l.indexes[:lenOfIndexes-1]
    }
    return
}

func (l List) Range(callback func(index int, value Value) bool) {
    l.data.Range(func(index int, curr *linkedlist.Node) bool {
        return callback(index, curr.Value)
    })
}

func (l *List) enqueueFirst(value Value) {
    l.data.Enqueue(linkedlist.NewNode(value))
    l.indexes = append(l.indexes, linkedlist.New())
    l.indexes[0].Enqueue(linkedlist.NewNode(newIndexNode(value, l.data.Head().Next())))
}

func (l *List) prepend(value Value) {
    l.data.Prepend(linkedlist.NewNode(value))
    indexNodeOf(l.indexes[0].Head().Next()).down = l.data.Head().Next()
    lenOfIndexes := len(l.indexes)
    for i := 0; i < lenOfIndexes; i++ {
        var (
            start       = l.indexes[i].Head().Next()
            slow, count = l.findPivot(start)
        )
        indexNodeOf(start).value = value
        if count <= l.MaxNum() {
            break
        }
        l.indexes[i].InsertAfter(start, linkedlist.NewNode(newIndexNode(slow.Value, slow)))
    }
    l.addIndexList(lenOfIndexes)
}

func (l *List) enqueue(value Value) {
    l.data.Enqueue(linkedlist.NewNode(value))
    lenOfIndexes := len(l.indexes)
    for i := 0; i < lenOfIndexes && l.addIndexWithPrev(i, l.indexes[i].Tail()); i++ {
    }
    l.addIndexList(lenOfIndexes)
}

func (l *List) addIndex(lenOfIndexes int, prevNodes []*linkedlist.Node) {
    for i := 0; i < lenOfIndexes && l.addIndexWithPrev(i, prevNodes[i]); i++ {
    }
    l.addIndexList(lenOfIndexes)
}

func (l *List) addIndexWithPrev(i int, prev *linkedlist.Node) (need bool) {
    slow, count := l.findPivot(prev)
    if count <= l.MaxNum() {
        return false
    }
    l.indexes[i].InsertAfter(prev, linkedlist.NewNode(newIndexNode(slow.Value, slow)))
    return true
}

func (l List) findPivot(start *linkedlist.Node) (slow *linkedlist.Node, count int) {
    end := start.Next()
    if end != nil {
        end = indexNodeOf(end).down
    }
    slow = indexNodeOf(start).down
    for quick := slow; quick != end && quick.Next() != end; quick, slow = quick.Next().Next(), slow.Next() {
        count += 2
    }
    return
}

func (l *List) addIndexList(lenOfIndexes int) {
    if length := l.indexes[lenOfIndexes-1].Len(); length > l.MaxNum() {
        var (
            head      = l.indexes[lenOfIndexes-1].Head().Next()
            curr      = l.indexes[lenOfIndexes-1].Node(length / 2)
            indexList = linkedlist.New()
        )
        indexList.Enqueue(linkedlist.NewNode(newIndexNode(indexNodeOf(head).value, head)))
        indexList.Enqueue(linkedlist.NewNode(newIndexNode(indexNodeOf(curr).value, curr)))
        l.indexes = append(l.indexes, indexList)
    }
}

func findPrevOf(value Value, start, end *linkedlist.Node) (prev *linkedlist.Node) {
    for curr := start; curr != end; curr = curr.Next() {
        if value.Compare(indexNodeOf(curr).value) == compare.LessThan {
            break
        }
        prev = curr
    }
    return
}

func startEndOfPrev(prev *linkedlist.Node) (start, end *linkedlist.Node) {
    start, end = indexNodeOf(prev).down, prev.Next()
    if end != nil {
        end = indexNodeOf(end).down
    }
    return
}

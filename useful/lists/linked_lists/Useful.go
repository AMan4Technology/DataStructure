package linked_lists

import (
    "github.com/AMan4Technology/DataStructure/useful/compare"

    "github.com/AMan4Technology/DataStructure/heap"
    . "github.com/AMan4Technology/DataStructure/list/linked"
)

func MergeTwoList(l1, l2 *List) *List {
    if l1.Empty() {
        return l2
    }
    if l2.Empty() {
        return l1
    }
    l := New()
    for i, j := l1.Node(0), l2.Node(0); i != nil || j != nil; {
        if i == nil {
            l.Enqueue(NewNode(j.Value))
            j = j.Next()
            continue
        }
        if j == nil {
            l.Enqueue(NewNode(i.Value))
            i = i.Next()
            continue
        }
        if j.Compare(i.Value) == compare.LessThan {
            l.Enqueue(NewNode(j.Value))
            j = j.Next()
        } else {
            l.Enqueue(NewNode(i.Value))
            i = i.Next()
        }
    }
    return &l
}

func MergeLists(lists []*List) *List {
    length := len(lists)
    if length == 0 {
        return nil
    }
    if length == 1 {
        return lists[0]
    }
    queue := heap.NewMin(length)
    for _, list := range lists {
        curr := list.Head().Next()
        if curr == nil {
            continue
        }
        _ = queue.Enqueue(curr, true)
    }
    if queue.Empty() {
        return nil
    }
    l := New()
    for !queue.Empty() {
        curr := queue.Top().(*Node)
        l.Enqueue(NewNode(curr.Value))
        if curr.Next() == nil {
            queue.Dequeue()
        } else {
            _ = queue.Update(curr.Next())
        }
    }
    return &l
}

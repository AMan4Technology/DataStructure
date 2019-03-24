package linkedqueue

import (
    "DataStructure/useful/common"

    "DataStructure/list/linked"
)

func New() Queue {
    return Queue{linkedlist.New()}
}

type Queue struct {
    list linkedlist.List
}

func (q Queue) Empty() bool {
    return q.list.Empty()
}

func (q Queue) Len() int {
    return q.list.Len()
}

func (q Queue) Head() common.Value {
    return q.list.Head().Next().Value.(nodeValue).Value
}

func (q *Queue) Enqueue(value common.Value) {
    q.list.Enqueue(linkedlist.NewNode(newNodeValue(value)))
}

func (q *Queue) Dequeue() common.Value {
    return q.list.Dequeue().Value.(nodeValue).Value
}

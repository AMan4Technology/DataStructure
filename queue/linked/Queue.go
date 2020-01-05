package linked_queue

import (
    "github.com/AMan4Technology/DataStructure/useful/common"

    "github.com/AMan4Technology/DataStructure/list/linked"
)

func New() Queue {
    return Queue{linked_list.New()}
}

type Queue struct {
    list linked_list.List
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
    q.list.Enqueue(linked_list.NewNode(newNodeValue(value)))
}

func (q *Queue) Dequeue() common.Value {
    return q.list.Dequeue().Value.(nodeValue).Value
}

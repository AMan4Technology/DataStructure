package graphs

import (
    "github.com/AMan4Technology/DataStructure/graph"
    "github.com/AMan4Technology/DataStructure/queue/array"
)

func TopoSort(g graph.Graph, callback func(key string, node *graph.Node) bool) {
    inDegrees := make(map[string]int, g.Len())
    g.Range(func(key string, node *graph.Node) bool {
        if _, ok := inDegrees[key]; !ok {
            inDegrees[key] = 0
        }
        node.Range(func(keyOfEnd string, edge *graph.Edge) bool {
            inDegrees[keyOfEnd]++
            return true
        })
        return true
    })
    queue := array_queue.New(1 << 6)
    for key, num := range inDegrees {
        if num == 0 {
            queue.Enqueue(key, true)
        }
    }
    for !queue.Empty() {
        var (
            key  = queue.Dequeue().(string)
            node = g.Node(key)
        )
        if !callback(key, node) {
            return
        }
        node.Range(func(keyOfEnd string, edge *graph.Edge) bool {
            num := inDegrees[keyOfEnd] - 1
            inDegrees[keyOfEnd] = num
            if num == 0 {
                queue.Enqueue(keyOfEnd, true)
            }
            return true
        })
    }
}

package graphs

import (
    "math"

    . "DataStructure/graph"
    "DataStructure/heap"
    "DataStructure/internal/example"
)

func ShortestPath(g Graph, start, end, name string) (path []string, min float64) {
    return shortestPath(g, start, end, name, nil)
}

func shortestPath(g Graph, start, end, name string, correct func(curr, end string) float64) (path []string, min float64) {
    var (
        startNode = g.Node(start)
        endNode   = g.Node(end)
    )
    if startNode == nil || endNode == nil {
        return nil, 0
    }
    if correct == nil {
        correct = func(curr, end string) float64 {
            return 0
        }
    }
    type distance struct {
        end string
        EdgeValue
    }
    var (
        minValue = make(map[string]float64, g.Len())
        paths    = make(map[string]string, g.Len())
        queue    = heap.NewMin(startNode.NumOfOut())
        inQueue  = make(map[string]bool, g.Len())
    )
    minValue[start] = 0
    paths[start] = start
    queue.Enqueue(example.ComparableValOf(distance{
        end:       start,
        EdgeValue: example.ValOf(0),
    }), true)
    inQueue[start] = true
    for !queue.Empty() {
        curr := example.RealValueOf(queue.Dequeue()).(distance).end
        if curr == end {
            break
        }
        inQueue[curr] = false
        g.RangeEdgeOf(curr, func(keyOfEnd string, edge *Edge) bool {
            if paths[curr] == keyOfEnd {
                return true
            }
            value := edge.Value(name)
            oldMin, ok := minValue[keyOfEnd]
            if !ok {
                oldMin = math.MaxFloat64
            }
            if newMin := minValue[curr] + value; newMin < oldMin {
                minValue[keyOfEnd] = newMin
                paths[keyOfEnd] = curr
                if inQueue[keyOfEnd] {
                    return true
                }
                queue.Enqueue(example.ComparableValOf(distance{
                    end:       keyOfEnd,
                    EdgeValue: example.ValOf(newMin + correct(keyOfEnd, end)),
                }), true)
                inQueue[keyOfEnd] = true
            }
            return true
        })
    }
    printPath(start, end, paths, &path)
    return path, minValue[end]
}

func printPath(start, end string, paths map[string]string, path *[]string) {
    if end == start {
        *path = append(*path, start)
        return
    }
    printPath(start, paths[end], paths, path)
    *path = append(*path, end)
}

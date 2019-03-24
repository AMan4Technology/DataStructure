package graphs

import (
	"DataStructure/queue/array"

	. "DataStructure/graph"
)

func BFS(g Graph, start string, callback func(key string, node *Node) bool) {
	startNode := g.Node(start)
	if startNode == nil || !callback(start, startNode) {
		return
	}
	var (
		queue   = arrayqueue.New(startNode.NumOfOut())
		visited = make(map[string]bool, g.Len())
	)
	queue.Enqueue(start, true)
	for !queue.Empty() {
		g.RangeEdgeOf(queue.Dequeue().(string), func(keyOfEnd string, edge *Edge) bool {
			if visited[keyOfEnd] {
				return true
			}
			if !callback(keyOfEnd, edge.End()) {
				queue.Clean()
				return false
			}
			visited[keyOfEnd] = true
			queue.Enqueue(keyOfEnd, true)
			return true
		})
	}
}

func DFS(g Graph, start string, callback func(key string, node *Node) bool) {
	startNode := g.Node(start)
	if startNode == nil {
		return
	}
	dfs(startNode, make(map[string]bool, g.Len()), callback)
}

func dfs(node *Node, visited map[string]bool, callback func(key string, node *Node) bool) (run bool) {
	key := node.Key()
	if visited[key] {
		return true
	}
	if !callback(key, node) {
		return false
	}
	visited[key] = true
	node.Range(func(keyOfEnd string, edge *Edge) bool {
		run = dfs(edge.End(), visited, callback)
		return run
	})
	return
}

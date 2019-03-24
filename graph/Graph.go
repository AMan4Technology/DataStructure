package graph

func New(capacity int) Graph {
	return Graph{
		capacity: capacity,
		nodes:    make(map[string]*Node, capacity),
	}
}

type Graph struct {
	capacity, length int
	nodes            map[string]*Node
}

func (g Graph) Empty() bool {
	return g.length == 0
}

func (g Graph) Cap() int {
	return g.capacity
}

func (g Graph) Len() int {
	return g.length
}

func (g Graph) Node(key string) *Node {
	return g.nodes[key]
}

func (g Graph) Edge(start, end string) *Edge {
	return g.nodes[start].outDegree[end]
}

func (g Graph) RangeEdgeOf(key string, callback func(keyOfEnd string, edge *Edge) bool) {
	g.nodes[key].Range(callback)
}

func (g Graph) Range(callback func(key string, node *Node) bool) {
	for key, node := range g.nodes {
		if !callback(key, node) {
			break
		}
	}
}

func (g *Graph) AddNode(key string, value NodeValue, update bool) (node *Node) {
	node = g.Node(key)
	if node == nil {
		node = newNode(key, value)
		g.nodes[key] = node
		g.length++
	} else if update {
		node.NodeValue = value
	}
	return
}

func (g *Graph) DeleteNode(key string) (node *Node) {
	node = g.Node(key)
	if node == nil {
		return
	}
	g.Range(func(k string, node *Node) bool {
		if edge := node.outDegree[key]; edge != nil {
			delete(node.outDegree, key)
		}
		return true
	})
	delete(g.nodes, key)
	g.length--
	return
}

func (g Graph) AddEdge(start, end string, value EdgeValue, update bool) (edge *Edge) {
	var (
		startNode = g.nodes[start]
		endNode   = g.nodes[end]
	)
	if startNode == nil || endNode == nil {
		return nil
	}
	edge = g.Edge(start, end)
	if edge == nil {
		edge = newEdge(startNode, endNode, value)
		startNode.outDegree[end] = edge
	} else if update {
		edge.EdgeValue = value
	}
	return
}

func (g Graph) DeleteEdge(start, end string) (edge *Edge) {
	edge = g.Edge(start, end)
	if edge != nil {
		delete(g.nodes[start].outDegree, end)
	}
	return
}

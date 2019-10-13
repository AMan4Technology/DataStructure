package graphs

import (
	"github.com/AMan4Technology/DataStructure/graph"
	"github.com/AMan4Technology/DataStructure/graph/astar"
)

func PathWithAstar(g graph.Graph, start, end, name string) (path []string, min float64) {
	return shortestPath(g, start, end, name, func(curr, end string) float64 {
		return astar.DistanceOfNodes(g.Node(curr).NodeValue, g.Node(end).NodeValue)
	})
}

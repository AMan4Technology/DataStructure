package graphs

import (
    "fmt"
    "testing"

    "github.com/AMan4Technology/DataStructure/graph"
)

func TestPathWithAstar(t *testing.T) {
    g := graph.New(5)
    fmt.Println(PathWithAstar(g, "a", "e", ""))
}

func TestShortestPath(t *testing.T) {
    g := graph.New(5)
    fmt.Println(ShortestPath(g, "a", "e", ""))
}

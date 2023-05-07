package unweighted_directed

import (
	"testing"
)

// Test_DFS tests unweighted_directed_graph.DFS
func Test_DFS(t *testing.T) {
	g := NewGraph(5)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)
	result := g.DFS(0)

	expected := []int{0, 1, 2, 3, 4}

	if len(result) != len(expected) {
		t.Errorf("result and expected lengths are not match len(result)=%d, len(expected)=%d", len(result), len(expected))
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("result[%d]=%d, expected[%d]=%d not match", i, result, i, expected)
		}
	}
}

package unweighted_directed

type Graph struct {
	V   int
	Adj map[int][]int
}

func NewGraph(v int) *Graph {
	return &Graph{
		V:   v,
		Adj: map[int][]int{},
	}
}

func (g *Graph) AddEdge(v, w int) {
	g.Adj[v] = append(g.Adj[v], w)
}

// DFS Depth First Search
func (g *Graph) DFS(v int) []int {
	var visited []bool
	for i := 0; i < g.V; i++ {
		visited = append(visited, false)
	}
	var traversalArray []int
	g.DFSUtil(v, visited, &traversalArray)

	return traversalArray
}

func (g *Graph) DFSUtil(v int, visited []bool, traversalArray *[]int) {
	visited[v] = true
	*traversalArray = append(*traversalArray, v)
	for i := 0; i < len(g.Adj[v]); i++ {
		if visited[g.Adj[v][i]] == false {
			g.DFSUtil(g.Adj[v][i], visited, traversalArray)
		}
	}
}

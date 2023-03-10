package main

import (
	"fmt"
	"github.com/mg52/data-structures-and-algorithms/bst"
	"github.com/mg52/data-structures-and-algorithms/graph/directed_graphs/unweighted_directed_graph"
	"github.com/mg52/data-structures-and-algorithms/linkedlist"
	"math/rand"
)

func main() {

	g := unweighted_directed_graph.NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)
	result := g.DFS(0)

	fmt.Println("unweighted_directed_graph DFS:", result)

	t := bst.NewBST()

	t.Insert('F')
	t.Insert('B')
	t.Insert('A')
	t.Insert('D')
	t.Insert('C')
	t.Insert('E')
	t.Insert('G')
	t.Insert('I')
	t.Insert('H')

	fmt.Printf("Pre Order: ")
	bst.PrintPreOrder(t.Root)
	fmt.Println()
	fmt.Printf("Post Order: ")
	bst.PrintPostOrder(t.Root)
	fmt.Println()
	fmt.Printf("In Order: ")
	bst.PrintInOrder(t.Root)
	fmt.Println()

	retVal := t.Search('D')
	if retVal == nil {
		fmt.Printf("Not found")
	} else {
		fmt.Printf("Found: Level: %d, Value: %c", retVal.Level, retVal.Value)
	}
	fmt.Println()

	retVal = t.Search('F')
	if retVal == nil {
		fmt.Printf("Not found")
	} else {
		fmt.Printf("Found: Level: %d, Value: %c", retVal.Level, retVal.Value)
	}
	fmt.Println()

	retVal = t.Search('J')
	if retVal == nil {
		fmt.Printf("Not found")
	} else {
		fmt.Printf("Found: Level: %d, Value: %c", retVal.Level, retVal.Value)
	}
	fmt.Println()
	fmt.Print("-------")
	fmt.Println()

	sl := linkedlist.List{}
	for i := 0; i < 5; i++ {
		sl.Insert(rand.Intn(100))
	}
	sl.PrintList()

	fmt.Printf("%v", sl.ToSlice())
}

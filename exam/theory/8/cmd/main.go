package main

import (
	"8/internal"
	"fmt"
)

func main() {

	graph := []internal.Edge{
		internal.NewEdge(1, 2, 1),
		internal.NewEdge(4, 5, 2),
		internal.NewEdge(1, 3, 3),
		internal.NewEdge(2, 3, 4),
		internal.NewEdge(3, 4, 5),
		internal.NewEdge(4, 5, 2),
		internal.NewEdge(2, 5, 7),
	}

	belongsToMST := internal.Kruskal(graph, 6)
	fmt.Println(belongsToMST)

	internal.MST(belongsToMST)
}

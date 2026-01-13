package main

import (
	"9/internal"
	"fmt"
)

func main() {
	n := 5

	// dist = [0 6 3 8 1]
	g := [][]internal.Edge{
		{
			internal.NewEdge(1, 7),
			internal.NewEdge(4, 1),
		},
		{
			internal.NewEdge(0, 7),
			internal.NewEdge(4, 8),
			internal.NewEdge(2, 3),
		},
		{
			internal.NewEdge(1, 3),
			internal.NewEdge(4, 2),
			internal.NewEdge(3, 6),
		},
		{
			internal.NewEdge(2, 6),
			internal.NewEdge(4, 7),
		},
		{
			internal.NewEdge(0, 1),
			internal.NewEdge(1, 8),
			internal.NewEdge(2, 2),
			internal.NewEdge(3, 7),
		},
	}

	parent, dist := internal.Dijkstra(g, n, 0)

	fmt.Println(dist)
	fmt.Println(parent)

	foundPath := internal.FindPath(0, 1, parent)

	for i := range foundPath {
		var s string
		if i == 0 {
			s = fmt.Sprintf("%d", foundPath[i])
		} else {
			s = fmt.Sprintf("->%d", foundPath[i])
		}
		fmt.Print(s)
	}

}

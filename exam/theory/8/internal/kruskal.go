package internal

import (
	"fmt"
	"slices"
)

type Edge struct {
	from   int
	to     int
	weight int
}

func NewEdge(from, to, weight int) Edge {
	return Edge{from, to, weight}
}

func Kruskal(edges []Edge, n int) []Edge {
	var result []Edge

	treeId := make([]int, n)
	for i := range n {
		treeId[i] = i
	}

	slices.SortFunc(edges, func(a, b Edge) int {
		return a.weight - b.weight
	})

	cost := 0

	for _, edge := range edges {
		w := edge.weight
		from := edge.from
		to := edge.to

		if treeId[from] != treeId[to] {
			cost += w
			result = append(result, edge)
			oldId := treeId[to]
			newId := treeId[from]

			for j := range n {
				if treeId[j] == oldId {
					treeId[j] = newId
				}
			}
		}

	}

	return result
}

func MST(edges []Edge) {
	for _, edge := range edges {
		fmt.Printf("From %d to %d for weight=%d\n", edge.from, edge.to, edge.weight)
	}
}

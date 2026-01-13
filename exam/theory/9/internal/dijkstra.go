package internal

import (
	"slices"
)

const inf = 1<<31 - 1

type Edge struct {
	to, weight int
}

func NewEdge(to, weight int) Edge {
	return Edge{to, weight}
}

func (e *Edge) To() int     { return e.to }
func (e *Edge) Weight() int { return e.weight }

type Path []int

func FindPath(source, to int, parent []int) []int {
	var result []int

	for v := to; v != source; v = parent[v] {
		result = append(result, v)
	}

	result = append(result, source)
	slices.Reverse(result)

	return result
}

func Dijkstra(g [][]Edge, n int, start int) ([]int, []int) {

	distance := make([]int, n)
	for i := range distance {
		distance[i] = inf
	}
	distance[start] = 0

	parent := make([]int, n)

	marked := make([]bool, n)

	for range n {
		v := -1
		for j := range n {
			if !marked[j] && (v == -1 || distance[j] < distance[v]) {
				v = j
			}
		}

		if distance[v] == inf {
			break
		}

		marked[v] = true

		for j := range g[v] {
			to := g[v][j].To()
			w := g[v][j].Weight()

			if distance[v]+w < distance[to] {
				distance[to] = distance[v] + w
				parent[to] = v
			}
		}
	}

	return parent, distance

}

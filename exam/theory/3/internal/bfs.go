package internal

import "math"

const inf = math.MaxInt32

func BFS(graph [][]int, source, destination int) int {
	if destination >= len(graph) || source >= len(graph) {
		return inf
	}

	n := len(graph)

	distance := make([]int, n)

	for i := range distance {
		distance[i] = inf
	}

	queue := &Queue{}
	queue.Enqueue(source)
	distance[source] = 0

	for queue.Size() > 0 {
		current, _ := queue.Dequeue()

		for _, neighbor := range graph[current] {
			if distance[neighbor] == inf {
				distance[neighbor] = distance[current] + 1

				queue.Enqueue(neighbor)
			}
		}
	}

	return distance[destination]
}

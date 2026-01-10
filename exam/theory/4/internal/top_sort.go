package internal

import "slices"

func dfs(graph [][]int, visited *[]bool, v int, answer *[]int) {
	(*visited)[v] = true
	for _, to := range graph[v] {
		if !(*visited)[to] {
			dfs(graph, visited, to, answer)
		}
	}
	*answer = append(*answer, v)
}

func TopSort(graph [][]int) []int {
	visited := make([]bool, len(graph))

	var answer []int

	for v := range graph {
		if !visited[v] {
			dfs(graph, &visited, v, &answer)
		}
	}

	slices.Reverse(answer)
	return answer
}

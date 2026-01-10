package internal

import "fmt"

func dfs(graph [][]int, visited *[]bool, u int) {
	(*visited)[u] = true
	for _, v := range graph[u] {
		if !(*visited)[v] {
			dfs(graph, visited, v)
		}
	}
}

func DoDFS(graph [][]int) {
	visited := make([]bool, len(graph))

	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			dfs(graph, &visited, i)
		}
	}

	fmt.Println(visited)
}

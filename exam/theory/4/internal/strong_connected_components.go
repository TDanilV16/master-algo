package internal

var order, component []int
var visited []bool
var graph, graphT [][]int

func dfs1(v int) {
	visited[v] = true

	for _, to := range graph[v] {
		if !visited[to] {
			dfs1(to)
		}
	}

	order = append(order, v)
}

func dfs2(v int) {
	visited[v] = true
	component = append(component, v)

	for _, to := range graphT[v] {
		if !visited[to] {
			dfs2(to)
		}
	}
}

func FindStrongConnectedComponents(g, gT [][]int) [][]int {
	var answer [][]int

	graph = g
	graphT = gT

	n := len(graph)

	visited = make([]bool, n)
	for i := range n {
		if !visited[i] {
			dfs1(i)
		}
	}

	visited = make([]bool, n)
	for i := range n {
		v := order[n-1-i]
		if !visited[v] {
			dfs2(v)
			compCopy := make([]int, len(component))
			copy(compCopy, component)
			answer = append(answer, compCopy)
			component = component[:0]
		}
	}

	return answer
}

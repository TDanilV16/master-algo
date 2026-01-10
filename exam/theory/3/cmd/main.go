package main

import (
	"3/internal"
)

func main() {
	graph := make([][]int, 6)
	graph[0] = []int{1, 2}
	graph[1] = []int{0, 3}
	graph[2] = []int{0, 3, 4}
	graph[3] = []int{1, 2, 5}
	graph[4] = []int{2, 5}
	graph[5] = []int{3, 4}

	internal.DoDFS(graph) // <- весь массив true, значит все посетили
	// аналогичные тесты можем написать для BFS, там можно проверить расстояния до вершин
}

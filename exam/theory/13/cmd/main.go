package main

import (
	"13/internal"
	"fmt"
)

func main() {
	items := []internal.Item{

		internal.NewItem(3, 1),
		internal.NewItem(4, 6),
		internal.NewItem(5, 4),
		internal.NewItem(8, 7),
		internal.NewItem(9, 6),
	}

	capacity := 13

	result := internal.SolveKnapsack01(items, capacity)
	fmt.Println(result)

	weights := []int{1, 4, 9}

	sum := internal.SolveSubsetSum(weights, 10)
	fmt.Println(sum)

}

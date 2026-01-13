package internal

import "slices"

type Item struct {
	weight, price int
}

func NewItem(weight int, price int) Item {
	return Item{weight, price}
}

func (i Item) Weight() int { return i.weight }
func (i Item) Price() int  { return i.price }

func SolveKnapsack01(items []Item, capacity int) []int {
	return solveKnapsack01(items, capacity)
}

func SolveSubsetSum(weights []int, capacity int) int {
	return solveSubsetSum(weights, capacity)
}

func solveSubsetSum(weights []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		weight := weights[i-1]
		for w := 1; w <= capacity; w++ {
			if weight <= w {
				take := dp[i-1][w-weight] + weight
				skip := dp[i-1][w]
				if take > skip {
					dp[i][w] = take
				} else {
					dp[i][w] = skip
				}
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

func solveKnapsack01(items []Item, capacity int) []int {

	n := len(items)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		item := items[i-1]
		for w := 1; w <= capacity; w++ {
			if item.weight <= w {
				take := dp[i-1][w-item.weight] + item.price
				skip := dp[i-1][w]
				if take > skip {
					dp[i][w] = take
				} else {
					dp[i][w] = skip
				}
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	selected := make([]int, 0)
	w := capacity
	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			selected = append(selected, i-1)
			w -= items[i-1].Weight()
		}
	}

	slices.Reverse(selected)

	return selected
}

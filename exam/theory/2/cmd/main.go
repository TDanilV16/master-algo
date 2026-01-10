package main

import (
	"2/internal"
	"fmt"
)

func main() {
	arr := []int{3, 6, 4, 9, 8, 12, 7, 11, 9}
	heap := internal.NewMinHeap(arr)
	fmt.Println(heap.ExtractMin())
	heap.Insert(-2)
	fmt.Println(heap.ExtractMin())
}

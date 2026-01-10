package internal

type PriorityQueue interface {
	Insert(priority int)
	ExtractMin() (int, bool)
}

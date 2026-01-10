package internal

import "fmt"

type Queue []int

func (q *Queue) Enqueue(v int) { *q = append(*q, v) }

func (q *Queue) Dequeue() (int, error) {
	if len(*q) == 0 {
		return 0, fmt.Errorf("empty queue")
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, nil
}

func (q *Queue) Size() int { return len(*q) }

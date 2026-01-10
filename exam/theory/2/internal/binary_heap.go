package internal

type MinHeap struct {
	array []int
}

func NewMinHeap(arr []int) *MinHeap {
	h := &MinHeap{arr}
	h.buildHeap()
	return h
}

func (h *MinHeap) Size() int { return len(h.array) }

func (h *MinHeap) ExtractMin() (int, bool) {
	if h.Size() == 0 {
		return 0, false // queue is empty
	}
	minVal := h.array[0]
	lastIndex := h.Size() - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.siftDown(0)
	return minVal, true // minimum found!
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.siftUp(h.Size() - 1)
}

func (h *MinHeap) buildHeap() {
	for i := h.Size() / 2; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *MinHeap) siftDown(i int) {
	for 2*i+1 < h.Size() {
		left := leftChild(i)
		right := rightChild(i)
		j := left
		if right < h.Size() && h.array[right] < h.array[left] {
			j = right
		}
		if h.array[i] <= h.array[j] {
			break
		}
		h.array[i], h.array[j] = h.array[j], h.array[i]
		i = j
	}

}

func (h *MinHeap) siftUp(i int) {
	for h.array[i] < h.array[parent(i)] {
		h.array[i], h.array[parent(i)] = h.array[parent(i)], h.array[i]
		i = parent(i)
	}
}

func parent(i int) int     { return (i - 1) / 2 }
func leftChild(i int) int  { return i*2 + 1 }
func rightChild(i int) int { return i*2 + 2 }

type MaxHeap struct {
	data []int
}

func heapify(nums []int) MaxHeap {
	h := MaxHeap{ data: append([]int{}, nums...) }

	for i := len(h.data) / 2 - 1; i >= 0; i -- {
		h.siftDown(i)
	}

	return h
}

func (h *MaxHeap) siftDown(i int) {
	n := len(h.data)

	for i < n {
		left := 2 * i + 1
		right := 2 * i + 2 

		if left >= n {
			break 
		}

		largest := left 

		if right < n && h.data[right] > h.data[largest] {
			largest = right 
		} 

		if h.data[i] >= h.data[largest] {
			break 
		}

		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest 
	}
}

func (h *MaxHeap) Pop() int {
	if len(h.data) == 0 {
		return -1001
	}

	maxVal := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.siftDown(0)

	return maxVal 
}

func (h *MaxHeap) Push(val int) {
	h.data = append(h.data, val)

	i := len(h.data)-1

	for i > 0 {
		parent := (i - 1) / 2 

		if h.data[parent] >= h.data[i] {
			break 
		}

		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent 
	}
}

func findKthLargest(nums []int, k int) int {
	maxHeap := heapify(nums)
	kthMax := 0 
	for i := 0; i < k; i ++ {
		kthMax = maxHeap.Pop()
	}

	return kthMax 
}

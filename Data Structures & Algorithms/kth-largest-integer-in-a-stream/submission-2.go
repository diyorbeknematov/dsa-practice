type KthLargest struct {
    k int 
	data []int 
}


func Constructor(k int, nums []int) KthLargest {
    h := KthLargest{
		k: k,
		data: append([]int{}, nums...),
	}

	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}

	for len(h.data) > k {
		h.Pop()
	}

	return h
}

func (h *KthLargest) siftDown(i int) {
	for i < len(h.data) {
		left := 2 * i + 1
		right := 2 * i + 2

		if left >= len(h.data) {
			break 
		}

		smallest := left 

		if right < len(h.data) && h.data[right] < h.data[left] {
			smallest = right 
		}

		if h.data[i] <= h.data[smallest] {
			break 
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

func (h *KthLargest) Pop() int {
	if len(h.data) == 0 {
		return -1001
	}

	minVal := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.siftDown(0)

	return minVal
}


func (this *KthLargest) Add(val int) int {
	if len(this.data) == 0 {
		this.data = append(this.data, val)
		return val
	}

	if this.data[0] > val {
		return this.data[0]
	}

    if len(this.data) == this.k {
		this.Pop()
	}

	this.data = append(this.data, val)

	i := len(this.data)-1
	for i > 0 {
		parent := (i - 1)/2

		if this.data[parent] <= this.data[i] {
			break 
		}

		this.data[i], this.data[parent] = this.data[parent], this.data[i]
		i = parent 
	}

	return this.data[0]
}

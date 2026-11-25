type MaxHeap struct {
	data []int 
}

func New(nums []int) MaxHeap {
	h := MaxHeap{
		data: append([]int{}, nums...),
	}

	for i := len(h.data) / 2 - 1; i >= 0; i -- {
		h.shiftDown(i)
	}

	return h
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

func (h *MaxHeap) Pop() int {
	if len(h.data) == 0 {
		return 0
	}

	maxVal := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.shiftDown(0)

	return maxVal 
}

func (h *MaxHeap) shiftDown(i int) {
	n := len(h.data)
	for i < n {
		left := 2 * i + 1
		right := 2 * i + 2 

		if left >= n {
			break 
		}

		largest := left 

		if right < n && h.data[right] > h.data[left] {
			largest = right
		}

		if h.data[i] >= h.data[largest] {
			break 
		}

		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
}

func lastStoneWeight(stones []int) int {
	mHeap := New(stones)
	
	for len(mHeap.data) > 1 {
		x := mHeap.Pop()
		y := mHeap.Pop()

		if x == y {
			continue 
		}

		mHeap.Push(abs(x-y))
	}

	if len(mHeap.data) == 0 {
		return 0
	}

	return mHeap.data[0]
}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}

	return num 
}

type Task struct {
	val byte
	freq int 
	available int 
}

type MaxHeap struct {
	data []Task 
}

func heapify(data []Task) MaxHeap{
	h := MaxHeap{ data: append([]Task{}, data...) }

	for i := len(h.data)/2-1; i >= 0; i -- {
		h.siftDown(i)
	}

	return h
}

func (h *MaxHeap) Push(task Task) {
	h.data = append(h.data, task)

	i := len(h.data)-1
	for i > 0 {
		parent := (i - 1)/2

		if h.data[parent].freq >= h.data[i].freq {
			break 
		}

		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent 
	}
}

func (h *MaxHeap) Pop() Task {
	if h.IsEmpty() {
		return Task{}
	}

	maxVal := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.siftDown(0)

	return maxVal 
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

		if right < n && h.data[right].freq > h.data[largest].freq {
			largest = right 
		}

		if h.data[i].freq >= h.data[largest].freq {
			break 
		}

		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest 
	} 
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.data) == 0 
}

func leastInterval(tasks []byte, n int) int {
	taskFreq := make(map[byte]int) 
	queue := make([]Task, 0)
	time := 0 

	for i := 0; i < len(tasks); i ++ {
		taskFreq[tasks[i]]++
	}

	t := make([]Task, 0, len(taskFreq))
	for k, v := range taskFreq {
		t = append(t, Task{
			val: k,
			freq: v,
		})
	}
	
	h := heapify(t)

	for {
		if h.IsEmpty() && len(queue) == 0 {
			break 
		}

		for len(queue) > 0 && queue[0].available <= time {
			task := queue[0]
			queue = queue[1:]
			h.Push(task)
		}

		if h.IsEmpty() {
			time = queue[0].available
			continue
		}

		task := h.Pop()
		task.freq--
		task.available = time + n + 1
		time++

		if task.freq != 0 {
			queue = append(queue, task)
		}
	}

	return time 
}

// ['x','x','y','y'] n = 2
// XY_XY
//
// ['A','A','A','A','A','B','B'','C','C'] n=3
// { 'A':5, 'B':2, 'C':2 }
// ABC_ABC_A___A___A
type Task struct {
    enqueueTime    int
    processingTime int
    index          int
}

type PriorityQueue []Task 

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool {
    if (*pq)[i].processingTime < (*pq)[j].processingTime {
        return true
    }

    return (*pq)[i].processingTime == (*pq)[j].processingTime &&
        (*pq)[i].index < (*pq)[j].index
}

func (pq *PriorityQueue) Swap(i, j int) {
    (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(Task)) }

func (pq *PriorityQueue) Pop() any {
    old := *pq
    item := old[len(old)-1]
    *pq = old[:len(old)-1]
    return item
}

func getOrder(tasks [][]int) []int {
	arr := make([]Task, len(tasks))

	for i := 0; i < len(tasks); i ++ {
		arr[i] = Task{
			enqueueTime: tasks[i][0],
			processingTime: tasks[i][1],
			index: i,
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].enqueueTime < arr[j].enqueueTime
	})

	clock := arr[0].enqueueTime
	res := make([]int, 0, len(arr))
	pq := &PriorityQueue{}
	heap.Init(pq)
	i := 0

	for len(res) < len(arr) {
		for i < len(arr) && arr[i].enqueueTime <= clock {
			heap.Push(pq, arr[i])
			i++
		}

		if pq.Len() == 0 {
			clock = arr[i].enqueueTime
			continue
		}
		item := heap.Pop(pq).(Task)
		res = append(res, item.index)
		clock += item.processingTime
	}

	return res
}

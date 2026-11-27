type Trip struct {
	passengers int
	to int 
}

type PriorityQueue []Trip 
func(pq *PriorityQueue) Len() int { return len(*pq) }
func(pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].to < (*pq)[j].to }
func(pq *PriorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func(pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(Trip)) }
func(pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func carPooling(trips [][]int, capacity int) bool {
	pq := &PriorityQueue{}
	heap.Init(pq)

	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})

	curCap := 0

	for _, trip := range trips {
		for pq.Len() != 0 && (*pq)[0].to <= trip[1] {
			t := heap.Pop(pq).(Trip)
			curCap -= t.passengers
		}

		curCap += trip[0]
		if curCap > capacity {
			return false
		}

		heap.Push(pq, Trip{
			passengers: trip[0],
			to: trip[2],
		})
	}

	return true
}

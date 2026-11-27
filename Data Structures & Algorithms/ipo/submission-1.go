type Project struct {
	profit int 
	capital int 
}

type MaxHeap []int

func (h *MaxHeap) Len() int { return len(*h) }

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }

func (h *MaxHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *MaxHeap) Pop() any {
    old := *h
    item := old[len(old)-1]
    *h = old[:len(old)-1]
    return item
}



func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := make([]Project, len(capital))

	for i := 0; i < len(capital); i ++ {
		projects[i] = Project{
			profit: profits[i],
			capital: capital[i],
		}
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	h := &MaxHeap{}
	heap.Init(h)
	res := w
	index := 0 
	for i := 0; i < k; i ++ {
		for index < len(projects) && projects[index].capital <= res {
			heap.Push(h, projects[index].profit)
			index++
		}
		if h.Len() == 0 {
			return res
		}

		p := heap.Pop(h).(int)
		res += p
	}

	return res
}
// profits=[1,4,2,3]   1 2 3 4 
// capital=[0,3,1,1]   0 1 1 3 
// k=3 w=0  1+3+4=8

// prof=[2,3,1,5,3] cap=[4,4,2,3,3]
// k=4 w=2
// cap = 2,3,3,4,4
// pro = 1 5 3 2 3
// 2+2+5+3+3

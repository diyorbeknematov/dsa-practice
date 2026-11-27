type MinHeap []int

func (h *MinHeap) Len() int { return len(*h) }

func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }

func (h *MinHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() any {
    old := *h
    item := old[len(old)-1]
    *h = old[:len(old)-1]
    return item
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

type MedianFinder struct {
    left MaxHeap
	right MinHeap
}


func Constructor() MedianFinder {
    maxHeap := MaxHeap{}
    minHeap := MinHeap{}

    heap.Init(&maxHeap)
    heap.Init(&minHeap)

    return MedianFinder{
        left:  maxHeap,
        right: minHeap,
    }
}


func (this *MedianFinder) AddNum(num int)  {

    if this.left.Len() == 0 || num <= this.left[0] {
		heap.Push(&this.left, num)
	} else {
		heap.Push(&this.right, num)
	}

	if this.left.Len() > this.right.Len()+1 {
		val := heap.Pop(&this.left).(int)
		heap.Push(&this.right, val)
	}

	if this.right.Len() > this.left.Len() {
		val := heap.Pop(&this.right).(int)
		heap.Push(&this.left, val)
	}
}


func (this *MedianFinder) FindMedian() float64 {
    if this.left.Len() == this.right.Len() {
		return (float64(this.left[0]) + float64(this.right[0])) / 2
	} 

	return float64(this.left[0])
}

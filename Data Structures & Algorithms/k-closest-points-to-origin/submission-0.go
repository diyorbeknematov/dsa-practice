type Point struct {
	x, y int 
	dis int 
}

type MinHeap struct {
	data []Point
}

func New(points []Point) MinHeap {

	h := MinHeap{ data: points }
	
	for i := len(h.data) / 2 - 1; i >= 0; i -- {
		h.shiftDown(i)
	}

	return h 
}

func (h *MinHeap) Pop() Point {
	if len(h.data) == 0 {
		return Point{}
	}

	minVal := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.shiftDown(0)

	return minVal 
}

func (h *MinHeap) shiftDown(i int) {
	n := len(h.data)
	for i < n {
		left := 2 * i + 1
		right := 2 * i + 2 

		if left >= n {
			break 
		}

		smallest := left 

		if right < n && h.data[right].dis < h.data[left].dis {
			smallest = right 
		}

		if h.data[i].dis <= h.data[smallest].dis {
			break 
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

func distance(point []int) int {
    x, y := point[0], point[1]
    return x*x + y*y
}



func kClosest(points [][]int, k int) [][]int {
	p := make([]Point, len(points))

	for i := 0; i < len(points); i++ {
		p[i] = Point{
			x: points[i][0],
			y: points[i][1],
			dis: distance(points[i]),
		}
	}

	minHeap := New(p)

	res := make([][]int, 0)

	for i := 0; i < k; i ++ {
		point := minHeap.Pop()

		res = append(res, []int{point.x, point.y})
	}

	return res
}

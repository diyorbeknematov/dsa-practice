type Char struct {
	val byte 
	freq int 
}

type PriorityQueue []Char 

func(pq *PriorityQueue) Len() int { return len(*pq) }

func(pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].freq > (*pq)[j].freq }

func(pq *PriorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] } 

func(pq *PriorityQueue) Push(x any) { (*pq) = append(*pq, x.(Char)) }

func(pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]

	return item 
}

func longestDiverseString(a int, b int, c int) string {
    chars := []Char{
        {val: 'a', freq: a},
        {val: 'b', freq: b},
        {val: 'c', freq: c},
    }

    pqItem := PriorityQueue{}

    for _, ch := range chars {
        if ch.freq > 0 {
            pqItem = append(pqItem, ch)
        }
    }

    pq := &pqItem
    heap.Init(pq)

    res := ""

    for pq.Len() > 0 {
        ch := heap.Pop(pq).(Char)

        if len(res) >= 2 && res[len(res)-1] == ch.val &&
            res[len(res)-2] == ch.val {

            if pq.Len() == 0 {
                break
            }

            next := heap.Pop(pq).(Char)

            res += string(next.val)
            next.freq--

            if next.freq > 0 {
                heap.Push(pq, next)
            }

            heap.Push(pq, ch)
        } else {
            res += string(ch.val)
            ch.freq--

            if ch.freq > 0 {
                heap.Push(pq, ch)
            }
        }
    }

    return res
}

// a=1, b=1, c=0
// babcabcab
// p1=a  p2=c

// a=0 b=0 c=1
// ccbcc
// p1=c   p2=c

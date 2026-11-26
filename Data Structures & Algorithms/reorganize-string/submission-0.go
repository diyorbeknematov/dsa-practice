type Char struct {
	val byte 
	freq int 
}

type PriorityQueue []Char 

func (pq *PriorityQueue) Len() int { return len(*pq) }
func (pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].freq > (*pq)[j].freq }
func (pq *PriorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func (pq *PriorityQueue) Push(val any) { *pq = append(*pq, val.(Char)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	item := old[len(old)-1]
	*pq = old[:len(old)-1]
	return item 
}

func reorganizeString(s string) string {
	freq := make(map[byte]int)

	for _, v := range s {
		freq[byte(v)]++
	}

	chars := make([]Char, 0, len(freq))
	for k, v := range freq {
		ch := Char{
			val: k,
			freq: v,
		}

		chars = append(chars, ch)
	}

	pqValue := PriorityQueue(chars)
	pq := &pqValue
	
	heap.Init(pq)
	var prev Char 
	res := ""

	for pq.Len() != 0 {
		ch := heap.Pop(pq).(Char)

		if prev.freq > 0 {
			heap.Push(pq, prev)
		}

		res += string(ch.val)
		ch.freq--
		prev = ch
	}

	if prev.freq > 0 {
		return ""
	}

	return res
}

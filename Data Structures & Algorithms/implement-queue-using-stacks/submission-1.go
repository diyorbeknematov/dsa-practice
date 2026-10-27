type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue { return MyQueue{} }

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) move() {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			n := len(q.in) - 1
			q.out = append(q.out, q.in[n])
			q.in = q.in[:n]
		}
	}
}

func (q *MyQueue) Pop() int {
	q.move()
	n := len(q.out) - 1
	v := q.out[n]
	q.out = q.out[:n]
	return v
}

func (q *MyQueue) Peek() int {
	q.move()
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */

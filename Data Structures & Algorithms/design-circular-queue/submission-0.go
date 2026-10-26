type MyCircularQueue struct {
    store []int 
	front int 
	rear int 
	size int 
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
		store: make([]int, k),
		front: 0,
		rear: 0,
		size: 0,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
		return false
	}
	
	if this.IsEmpty() {
		this.store[this.rear] = value
		this.front = this.rear
		this.size++
		return true
	}

	this.rear = (this.rear + 1) % len(this.store)
	this.store[this.rear] = value 
	this.size++
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
		return false
	}

	this.front = (this.front + 1) % len(this.store)
	this.size--	
	return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
		return -1
	}

	return this.store[this.front]
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
		return -1
	}

	return this.store[this.rear]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.size == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.size == len(this.store)
}

 
//           f  r
// [_, _, _, 4, 5, _, _]          k = 7
//             f/r
// [_, _, _, _, 5, _, _]          dequeue
//              r  f
// [_, _, _, _, _, _, _]          dequeue
//              r  f
// [_, _, _, _, 3, _, _]          enqueue 3

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 
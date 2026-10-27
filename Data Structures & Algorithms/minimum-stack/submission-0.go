type MinStack struct {
	space []int
	minPrefix []int
}

func Constructor() MinStack {
	return MinStack{
		space: make([]int, 0),
		minPrefix: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.space) == 0 {
		this.space = append(this.space, val)

		this.minPrefix = append(this.minPrefix, val)

		return
	}

	currMin := this.minPrefix[len(this.minPrefix)-1]

	if currMin > val {
		this.minPrefix = append(this.minPrefix, val)
	} else {
		this.minPrefix = append(this.minPrefix, currMin)
	}

	this.space = append(this.space, val)
}

func (this *MinStack) Pop() {
	this.space = this.space[:len(this.space)-1]
	this.minPrefix = this.minPrefix[:len(this.minPrefix)-1]
}

func (this *MinStack) Top() int {
	return this.space[len(this.space)-1]
}

func (this *MinStack) GetMin() int {
	return this.minPrefix[len(this.minPrefix)-1]
}

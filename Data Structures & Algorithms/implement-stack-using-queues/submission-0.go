type MyStack struct {
	q1, q2 []int
}

func Constructor() MyStack {
	return MyStack{
		q1: make([]int, 0),
		q2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.q2 = append(this.q2, x)
	for i := 0; i < len(this.q1); i ++ {
		this.q2 = append(this.q2, this.q1[i])
	}
	this.q1 = []int{}
	this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStack) Pop() int {
	value := this.q1[0]
	this.q1 = this.q1[1:]
	return value
}

func (this *MyStack) Top() int {
	return this.q1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0 && len(this.q2) == 0
}

// q1 = []  q2=[]  push 2
// q1 = []  q2=[2] + q1
// q1=q2=[2]    q2=q1=[] push 4
// q1=[2]  q2[4] + q1 = [4,2]
// q1=[4, 2]  q2=[]
// 

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */

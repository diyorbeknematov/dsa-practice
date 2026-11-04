type FreqStack struct {
	freq map[int]int
	group map[int][]int
	maxFreq int 
}

func Constructor() FreqStack {
	return FreqStack{
		freq: make(map[int]int),
		group: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	f := 0
	if _, ok := this.freq[val]; ok {
		f = this.freq[val]
	}

	f += 1
	this.freq[val] = f 
	this.group[f] = append(this.group[f], val)	

	this.maxFreq = max(this.maxFreq, f)
}

func (this *FreqStack) Pop() int {
	num := this.group[this.maxFreq][len(this.group[this.maxFreq])-1]
	this.freq[num]--
	this.group[this.maxFreq] = this.group[this.maxFreq][:len(this.group[this.maxFreq])-1]
	if len(this.group[this.maxFreq]) == 0 {
		this.maxFreq--
	}

	return num
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(val)
 * param2 := obj.Pop()
 */
 
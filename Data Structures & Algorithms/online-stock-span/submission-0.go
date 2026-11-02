type StockSpanner struct {
	stack []pair
}

type pair struct {
	price int 
	span  int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: make([]pair, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	ans := 1
	p := pair{
		price: price,
	}

	if len(this.stack) == 0 {
		p.span = 1
		this.stack = append(this.stack, p)
		return 1
	}

	for len(this.stack) > 0 {
		top := this.stack[len(this.stack)-1]

		if top.price > price {
			break
		}

		ans += top.span
		this.stack = this.stack[:len(this.stack)-1]
	}

	p.span = ans
	this.stack = append(this.stack, p)

	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 
// space = [100, 80, 60, 70, 60, 75, 85]        
// 
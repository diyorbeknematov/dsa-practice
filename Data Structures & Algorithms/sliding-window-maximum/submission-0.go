func maxSlidingWindow(nums []int, k int) []int {
    deQueue := make([]int, len(nums))
	head, tail := 0, 0 

	for i := 0; i < k; i++ {
		for tail > head && nums[deQueue[tail-1]] <= nums[i] {
			tail--
		}

		deQueue[tail] = i
		tail++
	}

	ans := make([]int, 0, len(nums)-k+1)
	ans = append(ans, nums[deQueue[head]])

	l := 1
	for r := k; r < len(nums); r ++ {
		if deQueue[head] < l {
			head++
		}

		for tail > head && nums[deQueue[tail-1]] <= nums[r] {
			tail--
		}

		deQueue[tail] = r
		tail++
		ans = append(ans, nums[deQueue[head]])
		l++
	}

	return ans
}


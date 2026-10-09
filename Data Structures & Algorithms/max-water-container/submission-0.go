func maxArea(heights []int) int {
	maxAmount := 0
	l, r := 0, len(heights) - 1

	for l < r {
		h := min(heights[l], heights[r])
		amount := h * (r - l)

		maxAmount = max(maxAmount, amount)

		if heights[l] > heights[r] {
			r--
		} else if heights[l] < heights[r] {
			l++
		} else {
			l++
			r--
		}
	}

	return maxAmount
}

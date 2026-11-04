func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]

			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := heights[top] * width
			maxArea = max(maxArea, area)
		}

		stack = append(stack, i)
	}

	// Stackda qolgan elementlarni hisoblash
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		width := len(heights)
		if len(stack) > 0 {
			width = len(heights) - stack[len(stack)-1] - 1
		}

		area := heights[top] * width
		maxArea = max(maxArea, area)
	}

	return maxArea
}
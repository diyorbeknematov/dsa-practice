func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for _, a := range asteroids {
		for a < 0 {
			n := len(stack)
			if n == 0 || stack[n-1] < 0 {
				break
			}

			b := stack[n-1]
			stack = stack[:n-1]

			if abs(a) > abs(b) {
				continue
			}

			if abs(a) < abs(b) {
				a = b
			} else {
				a = 0
			}
		}

		if a != 0 {
			stack = append(stack, a)
		}
	}

	return stack
}

func abs(num int) int {
	if num < 0 {
		return (-1) * num
	}

	return num
}

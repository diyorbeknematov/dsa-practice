func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for _, a := range asteroids {
		alive := true

		for alive && a < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]

			if top < -a {
				// Stackdagi asteroid kichikroq — u yo'q qilinadi.
				stack = stack[:len(stack)-1]
			} else if top == -a {
				// Ikkala asteroid ham yo'q qilinadi.
				stack = stack[:len(stack)-1]
				alive = false
			} else {
				// Joriy asteroid kichikroq — u yo'q qilinadi.
				alive = false
			}
		}

		if alive {
			stack = append(stack, a)
		}
	}

	return stack
}

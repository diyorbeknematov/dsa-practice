func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}

	s := 0 
	maxStick := matchsticks[0]
	for _, v := range matchsticks {
		s += v 
		maxStick = max(maxStick, v)
	}

	if s % 4 != 0 || s / 4 < maxStick {
		return false
	}

	target := s / 4 
	sides := [4]int{}
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	var backtrack func(index int) bool 

	backtrack = func(index int) bool {
		if index == len(matchsticks) {
			return true
		}

		for i := 0; i < 4; i ++ {
			if sides[i] + matchsticks[index] > target {
				continue
			}

			sides[i] += matchsticks[index]

			if backtrack(index + 1) {
				return true
			}

			sides[i] -= matchsticks[index]
		}

		return false
	}

	return backtrack(0)
}


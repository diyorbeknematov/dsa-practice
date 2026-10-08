func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		s := numbers[l] + numbers[r]
		if s > target {
			r--
		} else if s < target {
			l++
		} else {
			break
		}
	}

	return []int{l+1, r+1}
}


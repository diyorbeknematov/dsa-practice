func permute(nums []int) [][]int {
	res := make([][]int, 0)
	seen := make(map[int]bool)

	var backtrack func([]int, map[int]bool)

	backtrack = func(cur []int, seen map[int]bool) {
		if len(cur) == len(nums) {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return 
		}

		for i := 0; i < len(nums); i ++ {
			if seen[nums[i]] {
				continue
			}

			cur = append(cur, nums[i])
			seen[nums[i]] = true
			backtrack(cur, seen)

			cur = cur[:len(cur)-1]
			seen[nums[i]] = false
		}
	}

	backtrack([]int{}, seen)

	return res
}



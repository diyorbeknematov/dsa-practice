func permute(nums []int) [][]int {
	var res [][]int

	var backtrack func(perm []int, mask int) 

	backtrack = func(perm []int, mask int) {
		if len(perm) == len(nums) {
			temp := append([]int{}, perm...)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i ++ {
			if mask & (1 << i) == 0 {
				perm = append(perm, nums[i])
				backtrack(perm, mask | (1 << i))

				perm = perm[:len(perm)-1]
			}
		}
	}

	backtrack([]int{}, 0)

	return res
}



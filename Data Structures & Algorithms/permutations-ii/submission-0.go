func permuteUnique(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)

	var backtrack func([]int, map[int]bool) 

	backtrack = func(perm []int, picked map[int]bool) {
		if len(perm) == len(nums) {
			res = append(res, append([]int{}, perm...))
			return
		}

		for i := 0; i < len(nums); i ++ {
			if picked[i] {
				continue
			}

			if i > 0 && nums[i] == nums[i-1] && !picked[i-1] {
				continue
			}

			perm = append(perm, nums[i])
			picked[i] = true
			backtrack(perm, picked)

			perm = perm[:len(perm)-1]
			picked[i] = false
			
		}
	}

	backtrack([]int{}, map[int]bool{})

	return res
}

// [1, 1, 2]
//                  1                      -
//            [1,1]    [1]
//       [1,1,2] [1,1] [1,2] [1]

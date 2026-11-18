func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	var backtrack func(int, []int) 

	backtrack = func(start int, subset []int) {
		if start == len(nums) {
			res = append(res, append([]int{}, subset...))
			return 
		}

		temp := make([]int, len(subset))
		copy(temp, subset)
		res = append(res, temp)

		for i := start; i < len(nums); i ++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			subset = append(subset, nums[i])
			backtrack(i+1, subset)

			subset = subset[:len(subset)-1]
		}
	}

	backtrack(0, []int{})

	return res
}
// nums = [1,2,1] 
// 1 1 2 

// res = [[], [1,1], [1,1,2], [1,2], [2]]   subset = [2]



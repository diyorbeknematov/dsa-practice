func combinationSum(nums []int, target int) [][]int {
    res := make([][]int, 0)
	sub := make([]int, 0)

	var dfs func(int, int)

	dfs = func(start int, sum int) {
		if sum == target {
			res = append(res, append([]int{}, sub...))
			return
		}

		if sum > target {
			return 
		}

		for i := start; i < len(nums); i ++ {
			sub = append(sub, nums[i])

			dfs(i, sum + nums[i])

			sub = sub[:len(sub)-1]
		}
	}

	dfs(0, 0)

	return res
}

// 2 5 6 9     target = 9
// 2      sum = 9     sub = [2, 2, 5]


func subsets(nums []int) [][]int {
	ans := make([][]int, 0, 1 << len(nums))
	subset := make([]int, 0)

	var dfs func(int, []int)

	dfs = func(start int, subset []int) {
		if start == len(nums) {
			ans = append(ans, append([]int{}, subset...))
			return 
		}

		// choose 
		dfs(start + 1, append(subset, nums[start]))

		// unchose 
		dfs(start + 1, subset)
	}

	dfs(0, subset)

	return ans
}

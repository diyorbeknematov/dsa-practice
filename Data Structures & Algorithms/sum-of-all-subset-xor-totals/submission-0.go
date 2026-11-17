func subsetXORSum(nums []int) int {
	ans := 0

	var backtrack func(int, int)

	backtrack = func(start int, xor int) {
		for i := start; i < len(nums); i ++ {
			xor ^= nums[i]
			backtrack(i+1, xor)

			xor ^= nums[i]
		}

		ans += xor
	}

	backtrack(0, 0)

	return ans
}

func minSubArrayLen(target int, nums []int) int {
	ans := len(nums) + 1

	l, t := 0, 0
	for r := 0; r < len(nums); r++ {
		t += nums[r]
		for t >= target {
			ans = min(ans, r-l+1)
			t -= nums[l]
			l++
		}
	}

	if ans == len(nums) + 1 {
		return 0
	}

	return ans
}

func minSubArrayLen(target int, nums []int) int {
	ans := 100001

	l, t := 0, 0
	for r := 0; r < len(nums); r++ {
		t += nums[r]
		for t >= target {
			ans = min(ans, r-l+1)
			t -= nums[l]
			l++
		}
	}

	if ans == 100001 {
		return 0
	}

	return ans
}

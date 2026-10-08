func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)

	ans := make([][]int, 0)

	for i := 0; i < n; i ++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			if nums[l]+nums[r] == -nums[i] {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++
				}

				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if nums[l] + nums[r] > -nums[i] {
				r--
			} else {
				l++
			}
		}
	}

	return ans
}

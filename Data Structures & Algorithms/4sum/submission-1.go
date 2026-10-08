func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)
	n := len(nums)

	for i := 0; i < n-3; i ++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i+1; j < n-2; j ++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
        	right := n - 1
			t := target - (nums[i] + nums[j])

			for left < right {
				s := nums[left]+nums[right]
				if s < t {
					left++
				} else if s > t {
					right--
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--

					for left < right && nums[left] == nums[left-1] {
						left++
					}

					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}
			}
		}
	}

	return ans
}




func canPartitionKSubsets(nums []int, k int) bool {
    if len(nums) < k {
        return false
    }

    sum := 0
    maxNum := 0

    for _, v := range nums {
        sum += v
        maxNum = max(maxNum, v)
    }

    if sum%k != 0 || sum/k < maxNum {
        return false
    }

    target := sum / k
	used := make([]bool, len(nums))

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })

	var backtrack func(int, int, int) bool 
	backtrack = func(i int, k int, subsetSum int) bool {
		if k == 0 {
			return true
		}

		if subsetSum == target {
			return backtrack(0, k-1, 0)
		}

		for j := i; j < len(nums); j ++ {
			if used[j] || subsetSum + nums[j] > target {
				continue
			}

			used[j] = true
			if backtrack(j + 1, k, subsetSum+nums[j]) {
				return true
			}

			used[j] = false
		}

		return false
	}

	return backtrack(0, k, 0)
}


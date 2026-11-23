
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

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })

    subsets := make([]int, k)

    var backtrack func(int) bool

    backtrack = func(index int) bool {
        if index == len(nums) {
            return true
        }

        seen := make(map[int]bool)

        for i := 0; i < k; i++ {
            if seen[subsets[i]] {
                continue
            }
            seen[subsets[i]] = true

            if subsets[i]+nums[index] > target {
                continue
            }

            subsets[i] += nums[index]

            if backtrack(index + 1) {
                return true
            }

            subsets[i] -= nums[index]

            if subsets[i] == 0 {
                break
            }
        }

        return false
    }

    return backtrack(0)
}


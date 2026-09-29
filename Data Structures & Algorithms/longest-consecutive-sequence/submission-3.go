// func longestConsecutive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] < nums[j]
// 	})

//     maxLength := 0
// 	length := 1
// 	prev := nums[0]

// 	for _, num := range nums {
// 		if num-prev == 1 {
// 			length++
// 			prev = num
// 		} else if num-prev > 1 {
// 			maxLength = max(maxLength, length)
// 			length = 1
// 			prev = num
// 		}
// 	}

// 	return max(maxLength, length)
// }

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hashSet := make(map[int]bool)

	for _, num := range nums {
		hashSet[num] = true
	}

	length, maxLength := 0, 0

	for _, num := range nums {
        if hashSet[num-1] {
            continue
        }
		length = 1
		key := num + 1
		for hashSet[key] {
			length++
			key++
		}

		maxLength = max(maxLength, length)
	}

	return max(maxLength, length)
}

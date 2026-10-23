func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}


// [6, 2, 3, 4, 5, 1, 7, 2] 
// slow=6 -> 7 -> 2 -> 3 -> 4 -> 5  with 1 step
// fast=6 -> 2 -> 4 -> 1 -> 3 -> 5  with 2 step
// 
// slow=6 -> 7 -> 2  with 1 step
// fast=5 -> 1 -> 2 with 1 step     fast == slow yes return slow
// 

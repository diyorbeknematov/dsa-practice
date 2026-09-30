func subarraySum(nums []int, k int) int {
    freq := make(map[int]int)
    freq[0] = 1

    prefSum := 0
    ans := 0

    for _, num := range nums {
        prefSum += num

        ans += freq[prefSum-k]

        freq[prefSum]++
    }

    return ans
}

// 2 -1 1 2 
// 2
// 2 -1 
// 2 -1 1
// 2 -1 1 2 

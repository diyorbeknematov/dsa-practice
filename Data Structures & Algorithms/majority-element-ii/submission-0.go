func majorityElement(nums []int) []int {
    freq := make(map[int]int)

    for _, num := range nums {
        freq[num]++
    }

    res := make([]int, 0, len(freq))
    l := len(nums)
    for key, num := range freq {
        if l / 3 < num {
            res = append(res, key)
        }
    }

    return res
}

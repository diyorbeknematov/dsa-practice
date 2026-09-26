func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)

    for _, num := range nums {
        freq[num]++
    }

    buckets := make([][]int, len(nums)+1)

    for key, value := range freq {
        buckets[value] = append(buckets[value], key)
    }

    res := make([]int, 0, k)

    for i := len(buckets) - 1; i >= 0; i-- {
        for _, num := range buckets[i] {
            res = append(res, num)

            if len(res) == k {
                return res
            }
        }
    }

    return res
}

func majorityElement(nums []int) []int {
    candidate1, count1 := 0, 0
    candidate2, count2 := 0, 0

    for _, num := range nums {
        if count1 == 0 {
            candidate1 = num
        }

        if count2 == 0 && num != candidate1 {
            candidate2 = num
        }

        if num == candidate1 {
            count1 ++
        } else if num == candidate2 {
            count2 ++
        } else {
            count1 --
            count2 --
        }
    }

    count1, count2 = 0, 0
    for _, num := range nums {
        if candidate1 == num {
            count1++
        }
        if candidate2 == num {
            count2++
        }
    }

    ans := make([]int,0, 2)
    if count1 > len(nums)/3{
        ans = append(ans, candidate1)
    }

    if count2 > len(nums)/3{
        ans = append(ans, candidate2)
    }

    return ans
}

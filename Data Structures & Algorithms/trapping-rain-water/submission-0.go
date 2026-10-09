
func trap(height []int) int {
    n := len(height)
    if n == 0 {
        return 0
    }

    l, r := 0, n-1
    leftMax, rightMax := 0, 0
    ans := 0

    for l < r {
        if height[l] <= height[r] {
            if height[l] >= leftMax {
                leftMax = height[l]
            } else {
                ans += leftMax - height[l]
            }
            l++
        } else {
            if height[r] >= rightMax {
                rightMax = height[r]
            } else {
                ans += rightMax - height[r]
            }
            r--
        }
    }

    return ans
}
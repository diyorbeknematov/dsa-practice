
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }

    n, m := len(nums1), len(nums2)
    leftSize := (n + m + 1) / 2

    l, r := 0, n

    for l <= r {
        i := l + (r-l)/2
        j := leftSize - i

        // nums1 chap qismi nums2 o'ng qismidan katta
        if i > 0 && j < m && nums1[i-1] > nums2[j] {
            r = i - 1
            continue
        }

        // nums2 chap qismi nums1 o'ng qismidan katta
        if j > 0 && i < n && nums2[j-1] > nums1[i] {
            l = i + 1
            continue
        }

        // Chap qismdagi eng katta element
        leftMax := 0
        if i == 0 {
            leftMax = nums2[j-1]
        } else if j == 0 {
            leftMax = nums1[i-1]
        } else {
            leftMax = max(nums1[i-1], nums2[j-1])
        }

        // Elementlar soni toq
        if (n+m)%2 == 1 {
            return float64(leftMax)
        }

        // O'ng qismdagi eng kichik element
        rightMin := 0
        if i == n {
            rightMin = nums2[j]
        } else if j == m {
            rightMin = nums1[i]
        } else {
            rightMin = min(nums1[i], nums2[j])
        }

        return (float64(leftMax) + float64(rightMin)) / 2.0
    }

    return 0
}

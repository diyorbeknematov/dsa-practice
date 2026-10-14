func lengthOfLongestSubstring(s string) int {
    lastSeen := make(map[byte]int)

    left := 0
    ans := 0

    for right := 0; right < len(s); right++ {
        ch := s[right]

        if idx, ok := lastSeen[ch]; ok && idx >= left {
            left = idx + 1
        }

        ans = max(ans, right-left+1)
        lastSeen[ch] = right
    }

    return ans
}

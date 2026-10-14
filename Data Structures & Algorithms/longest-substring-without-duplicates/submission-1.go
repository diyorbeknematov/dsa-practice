func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]bool)

	left := 0
	ans := 0

	for right := 0; right < len(s); right++ {
		for seen[s[right]] {
			delete(seen, s[left])
			left++
		}

		ans = max(ans, right-left+1)
		seen[s[right]] = true
	}

	return ans
}

func characterReplacement(s string, k int) int {
	ans := 0

	var freq [26]int
	maxFreq := 0
	
	l := 0
	for r := 0; r < len(s); r++ {
		ch := s[r]-'A'
		freq[ch]++
		maxFreq = max(maxFreq, freq[ch])

		if (r-l+1) - maxFreq > k {
			freq[s[l]-'A']--
			l++
		}
		
		ans = max(ans, r-l+1)

	}
	return ans
}

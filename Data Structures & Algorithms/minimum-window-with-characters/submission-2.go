func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	var f1, f2 [128]int

	for i := 0; i < len(t); i++ {
		f1[t[i]]++
	}

	required := 0
	for _, count := range f1 {
		if count > 0 {
			required++
		}
	}

	formed := 0
	l := 0
	start := 0
	minLen := len(s) + 1

	for r := 0; r < len(s); r++ {
		ch := s[r]
		f2[ch]++

		if f1[ch] > 0 && f2[ch] == f1[ch] {
			formed++
		}

		for formed == required {
			if r-l+1 < minLen {
				start = l
				minLen = r - l + 1
			}

			leftChar := s[l]
			f2[leftChar]--

			if f1[leftChar] > 0 && f2[leftChar] < f1[leftChar] {
				formed--
			}

			l++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}
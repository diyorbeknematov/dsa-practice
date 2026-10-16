func checkInclusion(s1 string, s2 string) bool {
	wl := len(s1)
	var freqS1 [26]byte
	var freqS2 [26]byte
	
	for _, ch := range s1 {
		freqS1[ch-'a']++
	}

	l := 0
	for r := 0; r < len(s2); r++ {
		ch := s2[r]-'a'
		freqS2[ch]++

		if (r - l + 1) > wl {
			freqS2[s2[l]-'a']--
			l++
		}

		if isEqual(freqS1, freqS2) {
			return true
		}
	}

	return false
}

func isEqual(f1, f2 [26]byte) bool {
	for i := 0; i < 26; i ++ {
		if f1[i] != f2[i] {
			return false
		}
	}

	return true
}

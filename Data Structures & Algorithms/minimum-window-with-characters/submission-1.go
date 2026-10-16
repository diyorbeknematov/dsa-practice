func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

    f1 := make(map[byte]int)
	f2 := make(map[byte]int)

	for i := 0; i < len(t); i ++ {
		f1[t[i]]++
	}

	ans := ""
	l := 0
	for r := 0; r < len(s); r ++ {
		f2[s[r]]++

		for isPresent(f1, f2) {
			if ans == "" || len(ans) > (r-l+1) {
				ans = s[l : r+1]
			}
			f2[s[l]]--
			l++
		}
	}

	return ans
}

func isPresent(f1, f2 map[byte]int) bool {
	for k, v := range f1 {
		if f2[k] < v {
			return false
		}
	}

	return true
}

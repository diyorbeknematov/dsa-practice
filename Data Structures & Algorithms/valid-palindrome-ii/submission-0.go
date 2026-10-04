func validPalindrome(s string) bool {
	l, r := 0, len(s)-1 
	isPalindrome := true
	for l < r {
		if s[l] != s[r] {
			isPalindrome = false
			break
		}
		l ++
		r --
	}

	if isPalindrome { 
		return true
	}
		

	nL := l + 1
	nR := r
	isPalindrome = true
	for nL < nR {
		if s[nL] != s[nR] {
			isPalindrome = false
			break
		}
		nL ++
		nR --
	}
	
	if isPalindrome { 
		return true
	}

	nL = l 
	nR = r - 1
	isPalindrome = true
	for nL < nR {
		if s[nL] != s[nR] {
			isPalindrome = false
			break
		}
		nL ++
		nR --
	}

	return isPalindrome
}

// abbadc 
// i=0 j=5
// 
// abbda 
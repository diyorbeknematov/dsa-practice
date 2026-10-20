func minEatingSpeed(piles []int, h int) int {
	m := piles[0]
	for _, pile := range piles {
		if m < pile {
			m = pile
		}
	}

	l, r := 1, m
	ans := m

	for l <= r {
		mid := l + (r - l) / 2

		if isEatInH(piles, mid, h) {
			ans = mid
			r = mid-1
		} else {
			l = mid+1
		}
	}

	return ans
}

func isEatInH(piles []int, k, h int) bool {
	ah := 0

	for i := 0; i < len(piles); i++ {
		ah += (piles[i]+ k - 1)/k
		if ah > h {
			return false
		}
	}

	return ah <= h
}

// 1 2 3 4  h=9

// l=0 r=3 mid=0+(3-0)/2 = 1 k=2 yes 
// l=0 r=1 mid=0+(1-0)/2 = 0 k=1 no 
// l=1 r=1 mid=




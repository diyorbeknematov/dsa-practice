func shipWithinDays(weights []int, days int) int {
	s := 0
	for _, w := range weights {
		s += w
	}
	
	l, r := 1, s
	minCap := s

	for l <= r {
		mid := l + (r - l) / 2 

		if canUp(weights, mid, days) {
			minCap = mid
			r = mid - 1 
		} else {
			l = mid + 1
		}
	}

	return minCap
}

func canUp(weights []int, capacity, days int) bool {
    d, sw := 1, 0

    for _, w := range weights {
        if w > capacity {
            return false
        }

        if sw+w > capacity {
            d++
            sw = 0
        }

        sw += w
    }

    return d <= days
}
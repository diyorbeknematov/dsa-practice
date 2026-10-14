func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	maxVal := prices[n-1]
	ans := 0

	for i:=n-2; i >= 0; i -- {
		ans = max(ans, maxVal-prices[i])
		maxVal = max(maxVal, prices[i])
	}

	return ans
	
}

// 10 1 5 6 7 1 
// 
// min = 1
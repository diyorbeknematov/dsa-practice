func carPooling(trips [][]int, capacity int) bool {
	passChange := make([]int, 1001)
	for _, trip := range trips {
		numPass, start, end := trip[0], trip[1], trip[2]
		passChange[start] += numPass
		passChange[end] -= numPass
	}

	currPass := 0
	for i := 0; i <= 1000; i ++ {
		currPass += passChange[i]
		if currPass > capacity{
			return false
		}
	}

	return true
}

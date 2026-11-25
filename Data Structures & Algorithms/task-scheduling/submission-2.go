func leastInterval(tasks []byte, n int) int {
	var count [26]int 
	for _, task := range tasks {
		count[task-'A']++
	}

	maxCount := 0
	tasksWithMax := 0 
	for _, c := range count {
		if c > maxCount {
			maxCount = c
			tasksWithMax = 1
		} else if c == maxCount {
			tasksWithMax++
		}
	}

	slots := (maxCount - 1) * (n + 1) + tasksWithMax

	return max(len(tasks), slots)
}

// ['A','A','A','A','A','B','B'','C','C'] n=3
// { 'A':5, 'B':2, 'C':2 }
// A B C _ A B C _ A _ _ _ A _ _ _ A (n+1) * (maxCount-1) + 1 = (4)*4+1 = 17
//
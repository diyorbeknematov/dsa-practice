func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	count := 0
	l, r := 0, len(people)-1

	for l <= r {
		w := people[l]+people[r]
		if w > limit {
			r--
		} else {
			l++
			r--
		}

		count ++
	}

	return count
}

// 1 2 4 5  limit=6
// 1:5 2:4 4:2 5:1 

// 1 3 2 3 2 limit=3
// 1:2 3:0 2:1 3:0 2:1 

// 1 1 1 1 
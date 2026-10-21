type TimeMap struct {
	values map[string][]pair
}

type pair struct {
	timestamp int
    value     string
}

func Constructor() TimeMap {
	return TimeMap{
		values: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	p := pair{
		timestamp: timestamp,
		value: value,
	}

	if _, ok := this.values[key]; !ok {
		this.values[key] = []pair{p}
		return
	}

	this.values[key] = append(this.values[key], p)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.values[key]

	l, r := 0, len(arr)-1
	ans := ""
	for l <= r {
		mid := l + (r - l) / 2

		if arr[mid].timestamp <= timestamp {
			ans = arr[mid].value
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}

// 1 2 3 4 5 6 7 8 


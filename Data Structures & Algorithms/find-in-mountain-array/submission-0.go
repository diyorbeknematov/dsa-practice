/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
    
	l, r := 0, n-1
	
	for l < r {
		mid := l + (r - l) / 2 
		if mid < n - 1 && mountainArr.get(mid) < mountainArr.get(mid+1) {
			l = mid + 1
		} else {
			r = mid 
		}
	}

	pIndx := l
	l, r = 0, pIndx
	for l <= r {
		mid := l + (r - l) / 2
		
		num := mountainArr.get(mid)
		if num == target {
			return mid
		}

		if num < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	l, r = pIndx+1, n-1
	for l <= r {
		mid := l + (r - l) / 2
		
		num := mountainArr.get(mid)
		if num == target {
			return mid
		}

		if num < target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

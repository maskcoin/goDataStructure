package midvaluesearch

func MidValueSearch(arr []int, data int) int {
	ret := -1
	left := 0
	right := len(arr) - 1
	for left <= right {
		leftV := float64(data - arr[left])
		allV := float64(arr[right] - arr[left])
		diff := float64(right - left)
		mid := int(float64(left) + diff*(leftV/allV))

		if mid < 0 || mid >= len(arr) {
			ret = -1
			return ret
		}

		if arr[mid] > data {
			right = mid - 1
		} else if arr[mid] < data {
			left = mid + 1
		} else {
			ret = mid
			break
		}
	}
	return ret
}

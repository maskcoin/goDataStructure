package thirdsearchsort

func ThirdSearchSort(arr []int, data int) int {
	left := 0
	right := len(arr) - 1

	ret := -1

	for left <= right  {
		if arr[left] > data {
			ret = -1
			break
		}

		if arr[right] < data {
			ret = -1
			break
		}

		mid1 := left + int((right-left)/3)
		mid2 := right - int((right-left)/3)
		mid1Val := arr[mid1]
		mid2Val := arr[mid2]
		if mid1Val == data {
			ret = mid1
			break
		} else if mid2Val == data {
			ret = mid2
			break
		}

		if mid1Val < data {
			left = mid1
		} else {
			right = mid1
		}

		if mid2Val > data {
			right = mid2
		} else {
			left = mid1
		}
	}

	return ret
}

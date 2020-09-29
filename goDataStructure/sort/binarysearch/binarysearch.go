package binarysearch

func BinSearch(arr []int, data int) int {
	ret := -1
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
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

//A找到第一个等于3的
func BinSearchA(arr []int, data int) int {
	ret := -1
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > data {
			right = mid - 1
		} else if arr[mid] < data {
			left = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != data {
				ret = mid
				break
			} else {
				right = mid - 1
			}
		}
	}
	return ret
}

//B找到最后一个等于3
func BinSearchB(arr []int, data int) int {
	ret := -1
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > data {
			right = mid - 1
		} else if arr[mid] < data {
			left = mid + 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != data {
				ret = mid
				break
			} else {
				left = mid + 1
			}
		}
	}
	return ret
}

// C找到第一个大于等于3
func BinSearchC(arr []int, data int) int {
	ret := -1
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] < data {
			left = mid + 1
		} else {
			if mid == 0 || arr[mid-1] < data {
				ret = mid
				break
			} else {
				right = mid - 1
			}
		}
	}
	return ret
}

// D找到最后一个小于等于6的数据
func BinSearchD(arr []int, data int) int {
	ret := -1
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > data {
			right = mid - 1
		} else {
			if mid == len(arr) - 1 || arr[mid+1] > data {
				ret = mid
				break
			} else {
				left = mid + 1
			}
		}
	}
	return ret
}


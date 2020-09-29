package combsort

func CombSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		gap := length
		for gap > 1 {
			gap = gap * 10 / 13
			for i := 0; i+gap < length; i++ {
				if arr[i] > arr[i+gap] {
					arr[i], arr[i+gap] = arr[i+gap], arr[i]
				}
			}
		}
	}

	return arr
}

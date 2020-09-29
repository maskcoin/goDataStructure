package insertsort

func InsertSort(arr []int) []int {
	length := len(arr)
	if length > 1 {
		for i := 1; i < length; i++ {
			backup := arr[i]
			j := i-1
			for j>=0 && backup<arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = backup
		}
	}

	return arr
}

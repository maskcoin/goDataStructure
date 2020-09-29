package selectsort

func SelectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length - 1; i++ {
		for j := i+1; j < length; j++ {
			if arr[j] > arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

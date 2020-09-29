package bubblesort

func BubbleSort(arr []int) []int {
	length := len(arr)

	for length > 1 {
		for i := 0; i < length-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
		length--
	}

	return arr
}

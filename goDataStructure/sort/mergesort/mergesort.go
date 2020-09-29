package mergesort

func Merge(leftArr []int, rightArr []int) []int {
	len1 := len(leftArr)
	len2 := len(rightArr)
	var retArr []int
	var i, j int
	for i < len1 && j < len2 {
		if leftArr[i] <= rightArr[j] {
			retArr = append(retArr, leftArr[i])
			i++
		} else if leftArr[i] > rightArr[j] {
			retArr = append(retArr, rightArr[j])
			j++
		}
	}
	for j < len2 {
		retArr = append(retArr, rightArr[j])
		j++
	}
	for i < len1 {
		retArr = append(retArr, leftArr[i])
		i++
	}

	return retArr
}

func MergeSort(arr []int) []int {
	length := len(arr)

	if length <= 1 {
		return arr
	} else {
		mid := length / 2
		leftArr := MergeSort(arr[:mid])
		rightArr := MergeSort(arr[mid:])
		return Merge(leftArr, rightArr)
	}
}


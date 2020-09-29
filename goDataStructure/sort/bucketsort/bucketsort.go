package bucketsort

import "datastructure/sort/heapsort"

func BitSort(arr []int, bit int) []int {
	length := len(arr)
	var bitCounts []int = make([]int, 10) // 统计长度
	for i := 0; i < length; i++ {
		num := (arr[i] / bit) % 10
		bitCounts[num]++ // 统计余数相等的个数
	}
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1] // 用于确定位置
	}
	tmp := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitCounts[num]-1] = arr[i]
		bitCounts[num]--
	}
	return tmp
}

func BucketSort(arr []int) []int {
	max := heapsort.HeapSortMax(arr, len(arr))[0]
	for bit := 1; max/bit > 1; bit *= 10 {
		// 按照数量级分段
		arr = BitSort(arr, bit) // 每次处理一个级别的排序
	}
	return arr
}

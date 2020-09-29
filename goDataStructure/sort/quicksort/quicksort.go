package quicksort

import (
	"datastructure/link/stack"
)

func GoroutineQuickSort(arr []int, lastChan chan int, level int, goroutineNum int)  {
	level = level*2 //每加深一个级别，协程数翻倍
	if len(arr) ==0 {
		close(lastChan)
		return
	} else if len(arr) == 1 {
		lastChan<-arr[0]
		close(lastChan)
		return
	} else {
		smallArr := make([]int, 0)
		bigArr := make([]int, 0)
		midArr := make([]int, 0)
		left := arr[0]
		for i := 0; i < len(arr); i++ {
			if arr[i] < left {
				smallArr = append(smallArr, arr[i])
			} else if arr[i] > left {
				bigArr = append(bigArr, arr[i])
			} else {
				midArr = append(midArr, arr[i])
			}
		}
		smallChan := make(chan int)
		bigChan := make(chan int)
		if level <= goroutineNum {
				go GoroutineQuickSort(smallArr, smallChan, level, goroutineNum)
				go GoroutineQuickSort(bigArr, bigChan, level, goroutineNum)
		} else {
			GoroutineQuickSort(smallArr, smallChan, level, goroutineNum)
			GoroutineQuickSort(bigArr, bigChan, level, goroutineNum)
		}
		for v := range smallChan {
			lastChan<-v
		}
		for _, v := range midArr {
			lastChan<-v
		}
		for v := range bigChan {
			lastChan<-v
		}
		close(lastChan)
	}
}

func position(arr []uint32, begin int, end int) int { //让begin对应的数据就位
	pv := arr[begin]
	for begin < end {
		for begin < end && arr[end] > pv {
			end--
		}
		arr[begin], arr[end] = arr[end], arr[begin]

		for begin < end && arr[begin] <= pv {
			begin++
		}
		arr[begin], arr[end] = arr[end], arr[begin]
	}

	arr[begin] = pv
	return begin
}
func QuickSort(arr []uint32) []uint32 {
	length := len(arr)
	if length > 1 {
		pov := position(arr, 0, length-1)
		arrLeft := arr[:pov]
		arrRight := arr[pov+1:]
		QuickSort(arrLeft)
		QuickSort(arrRight)
	}
	return arr
}

func QuickSortWithStack(arr []uint32, stack *stack.Node) []uint32 {
	stack.Push(arr)

	for !stack.IsEmpty() {
		popArr := stack.Pop()
		arr1 := popArr.([]uint32)
		length := len(arr1)
		if length > 1 {
			pov := position(arr1, 0, length-1)
			arrLeft := arr1[:pov]
			arrRight := arr1[pov+1:]
			stack.Push(arrLeft)
			stack.Push(arrRight)
		}
	}
	return arr
}

func QuickSortWithSpace(arr []int) []int {
	length := len(arr)

	if length <= 1 {
		return arr
	} else {
		pv := arr[0]
		var low []int
		var high []int
		var mid []int
		for i := 0; i < length; i++ {
			if arr[i] < pv {
				low = append(low, arr[i])
			} else if arr[i] > pv {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}

		low, high = QuickSortWithSpace(low), QuickSortWithSpace(high)
		return append(append(low, mid...), high...)
	}
}


package shellsort

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap { // 插入排序的变种
		backup := arr[i]
		j := i - gap
		for j >= 0 && backup < arr[j] {
			arr[i] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup
	}
}

func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ {
				ShellSortStep(arr, i, gap)
			}
			gap--
		}
	}

	return arr
}

func ShellSortGoRoutine(arr []int) []int {
	if len(arr) < 2 || arr == nil {
		return arr
	} else {
		//步长缩短
		for gap := len(arr) / 2; gap > 0; gap /= 2 {
			ch := make(chan int, gap)
			retChan := make(chan bool, gap)
			go func() {
				for k := 0; k < gap; k++ {
					ch <- k
				}
				close(ch)
			}()
			for k := 0; k < gap; k++ {
				go func() {
					for v := range ch {
						ShellSortStep(arr, v, gap)
					}
					retChan <- true
				}()
			}

			for i := 0; i < gap; i++ {
				<-retChan
			}
		}
		return arr
	}
}

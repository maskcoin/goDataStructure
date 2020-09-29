package truebucketsort

import "fmt"

func BucketSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		num := 4
		buckets := make([][]int, num)
		for i := 0; i < length; i++ {
			buckets[arr[i] - 1] = append(buckets[arr[i] - 1], arr[i])
		}

		var retArr []int
		for i :=0;i< num ; i++ {
			retArr = append(retArr, buckets[i]...)
		}
		return retArr
	}
}

func main1() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 2, 3, 1}
	fmt.Println(BucketSort(arr))
}

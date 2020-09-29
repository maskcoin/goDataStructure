package sleepsort

import (
	"fmt"
	"time"
)

var flag bool
var container chan bool
var count int

func toSleep(data int) {
	time.Sleep(time.Duration(data) * time.Microsecond * 1000)
	container <- true
}

func listen(size int) {
	for flag {
		select {
		case <-container:
			count++
			if count >= size {
				flag = false
			}
		}
	}
}

func main1() {
	arr := []int{16, 8, 1, 24, 30}
	flag = true
	container = make(chan bool, len(arr))

	for i := 0; i < len(arr); i++ {
		go toSleep(arr[i])
	}

	go listen(len(arr))

	for flag {
		time.Sleep(1 * time.Second)
	}

	fmt.Println(arr)
}

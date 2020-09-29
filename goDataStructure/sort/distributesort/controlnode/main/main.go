package main

import (
	"datastructure/sort/distributesort"
	"fmt"
	"net"
	"strconv"
)

func SendArray(arr []int, conn net.Conn) {
	start := distributesort.IntToBytes(0)
	start = append(start, distributesort.IntToBytes(0)...)
	conn.Write(start)

	for i := 0; i < len(arr); i++ {
		data := distributesort.IntToBytes(1)
		data = append(data, distributesort.IntToBytes(arr[i])...)
		conn.Write(data)
	}

	end := distributesort.IntToBytes(0)
	end = append(end, distributesort.IntToBytes(1)...)
	conn.Write(end)
}

func ServerMsgHandle(conn net.Conn, myChan chan<- int) {
	arr := []int{}
	for {
		buf := make([]byte, 16)
		n, err := conn.Read(buf)

		if n != 0 {
			if n == 16 {
				data1 := distributesort.BytesToInt(buf[:len(buf)/2])
				data2 := distributesort.BytesToInt(buf[len(buf)/2:])
				if data1 == 0 && data2 == 0 {
					//arr = make([]int, 0)
				}

				if data1 == 1 {
					arr = append(arr, data2)
				}

				if data1 == 0 && data2 == 1 {
					fmt.Println("数组接受完成", arr)
					for i := 0; i < len(arr); i++ {
						myChan <- arr[i] //
					}
					close(myChan)
				}
			}
		}

		if err != nil {
			fmt.Println("conn closed", conn)
			break
		}
	}
}

func main() {
	var chans  [2]chan int


	for i := 0; i < 2; i++ {
		conn, err := net.Dial("tcp", "0.0.0.0:700"+strconv.Itoa(i+1))
		if err != nil {
			panic(err)
		}

		arrs := [][]int{{1, 9, 2, 8, 7, 3, 5, 6, 10, 4, 23, 24}, {11, 19, 12, 18, 17, 13, 15, 16, 110, 14, 123, 124}}

		chans[i] = make(chan int)
		go SendArray(arrs[i], conn)

		go ServerMsgHandle(conn, chans[i])

	}

	lastChan := distributesort.Merge(chans[0], chans[1])
	for v := range lastChan {
		fmt.Println(v)
	}
}

package main

import (
	"datastructure/sort/distributesort"
	"fmt"
	"net"
)

func main()  {
	conn, err := net.Dial("tcp", "0.0.0.0:8848")
	if err != nil {
		panic(err)
	}

	arr := []int{1,9,2,8,7,3,5,6,10,4, 23, 24}
	length := len(arr)
	//0 0 开始传输
	//1 1
	//1 9
	//1 ...
	//1 4
	//0 1 传输结束
	start := distributesort.IntToBytes(0)
	start = append(start, distributesort.IntToBytes(0)...)
	conn.Write(start)

	for i := 0; i < length; i++ {
		data := distributesort.IntToBytes(1)
		data = append(data, distributesort.IntToBytes(arr[i])...)
		conn.Write(data)
	}
	
	end := distributesort.IntToBytes(0)
	end = append(end, distributesort.IntToBytes(1)...)
	conn.Write(end)

	arr = []int{}
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

				if data1==1 {
					arr = append(arr, data2)
				}

				if data1 == 0 && data2 == 1 {
					fmt.Println("数组接受完成", arr)
				}
			}
		}

		if err != nil {
			fmt.Println("conn closed", conn)
			break
		}
	}

	//for {
	//	var input string
	//	fmt.Scanf("%s\n", &input)
	//
	//	conn.Write([]byte(input))
	//	buf := make([]byte, 1024)
	//	n, err := conn.Read(buf)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(string(buf[:n]))
	//}
}

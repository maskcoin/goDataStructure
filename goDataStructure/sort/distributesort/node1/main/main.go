package main

import (
	"datastructure/sort/distributesort"
	"datastructure/sort/quicksort"
	"fmt"
	"net"
	"time"
)

//判断30秒内有没有产生通信
func HeartBeat(conn net.Conn, beatChan chan byte, timeOut int) {
	for {
		select {
		case <-beatChan:
			conn.SetDeadline(time.Now().Add(time.Duration(timeOut) * time.Second))
		case <-time.After(time.Duration(timeOut) * time.Second):
			fmt.Println("time out", conn.RemoteAddr())
			conn.Close()
		}
	}
}

//向心跳的channel中写入数据
func HeartChanHandle(msg []byte, beatChan chan byte) {
	for _, v := range msg {
		beatChan <- v
	}
}

func handleConnection(conn net.Conn) {
	var arr []int
	buf := make([]byte, 16)
	defer conn.Close()
	//beatChan := make(chan byte)
	//go HeartBeat(conn, beatChan, 3)
	for {
		n, err := conn.Read(buf)

		if n != 0 {
			//if string(buf[0:1]) == "0" {
			//	fmt.Println("host data:", string(buf[1:n]))
			//	conn.Write([]byte("收到数据：" + string(buf[1:n]) + "\n"))
			//} else {
			//	fmt.Println("host cmd:", string(buf[1:n]))
			//	cmd := exec.Command(string(buf[1:n]))
			//	err = cmd.Run()
			//	if err != nil {
			//		fmt.Println("err=", err)
			//	}
			//	conn.Write([]byte("收到命令：" + string(buf[1:n]) + "\n"))
			//}

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
					// 排序完后，再发回去
					arr = quicksort.QuickSort(arr)
					fmt.Println("排序后", arr)


					//arr = []int{}
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
			}

			//msg := buf[0:1]

			//go HeartChanHandle(msg, beatChan)
		}

		if err != nil {
			fmt.Println("conn closed", conn)
			break
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:7001")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

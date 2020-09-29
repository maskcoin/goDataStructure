package main

import (
	"bufio"
	"datastructure/link/stack"
	"datastructure/sort/AStar"
	"datastructure/sort/calc"
	"datastructure/sort/heapsort"
	"datastructure/sort/quicksort"
	"datastructure/sort/recursion"
	"datastructure/sort/shellsort"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	presetMap := []string{
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		"X . X X X X X X X X X X X X X X X X X X X X X X X X X",
		". . . . . . . . . . . . . . X . . . . . . . . . . . .",
		". . . . . . . . . . . . . . X . . . . . . . . . . . .",
		". . . . . . . . . . . . . . X . . . . . . . . . . . .",
		". . . . . . . . . . . . . . X . . . . . . . . . . . .",
		". . . . . . . . . . . . . . X . . . . . . . . . . . .",
		". X X X X X X X X X X X X X X . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		"X X X X X X X X X X X X X X X X X X X X X X X X . X X",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		"X . X X X X X X X X X X X X X X X X X X X X X X X X X",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
	}
	m := AStar.NewMap(presetMap)
	//m.PrintMap(nil)
	sr := AStar.NewSearchRoad(0,0, 4, 0, m)
	if sr.FindOutShortestPath() {
		fmt.Println("找到")
		m.PrintMap(sr)
	} else {
		fmt.Println("找不到")
	}
}

func main90() {
	count := 0
	fmt.Println(count)
}

func main89() {
	fmt.Println(calc.Calc("1+2*(1+(1+2*3))"))
}

func main88() {
	fmt.Println(recursion.Rec(0, recursion.W2))
	//fmt.Println(recursion.Solve())
	//for i := 0; i <= recursion.N2; i++ {
	//	for j := 0; j <= recursion.W2; j++ {
	//		fmt.Print(recursion.DP[i][j], " ")
	//	}
	//	fmt.Println()
	//}
}

func main87() {
	recursion.Init()
	recursion.StartGoQueen()
}

func main86() {
	recursion.Knapsack()
	fmt.Println(recursion.B[5][20])
}

func main85() {
	isOK := recursion.AIOut(recursion.AIData, 0, 0)
	if !isOK {
		fmt.Println("走不出来")
	}
}

func main84() {
	recursion.Show(recursion.Data)
	for {
		var inputStr string
		fmt.Scanf("%s\n", &inputStr)
		recursion.Run(inputStr)
	}
}

func main83() {
	arr := []uint32{1, 9, 2, 8, 3, 7, 6, 4, 5, 10}
	fmt.Println(arr)
	quicksort.QuickSortWithStack(arr, stack.NewStack())
	fmt.Println(arr)
}

func main82() {
	heap := heapsort.NewLeftHeap(3)
	heap = heap.Insert(2)
	heap = heap.Insert(1)
	heap = heap.Insert(4)

	heapsort.Print(heap)
	heap, v := heap.DeleteMin()
	fmt.Println(v)
	fmt.Println("删除后重新打印树")
	heapsort.Print(heap)
}

func main81() {
	arr := []int{1, 90, 2, 80, 13, 7, 6, 4, 5, 10}
	lastChan := make(chan int)
	go quicksort.GoroutineQuickSort(arr, lastChan, 0, 0)
	for v := range lastChan {
		fmt.Println(v)
	}
}

//func main14() {
//	p := distributesort.CreateNetworkPipe("/Users/xuchanghui/Desktop/big.in", 80, 2)
//	distributesort.WriteToFile(p, "/Users/xuchanghui/Desktop/big.out")
//	distributesort.ShowFile("/Users/xuchanghui/Desktop/big.out")
//}
//
//func main13() {
//	p := distributesort.CreatePipe("/Users/xuchanghui/Desktop/big.in", 8000000, 800)
//	distributesort.WriteToFile(p, "/Users/xuchanghui/Desktop/big.out")
//	distributesort.ShowFile("/Users/xuchanghui/Desktop/big.out")
//}
//
//func main12() {
//	pipe := distributesort.Merge(
//		distributesort.InMemorySort(distributesort.ArraySource(3, 9, 2, 1, 10)),
//		distributesort.InMemorySort(distributesort.ArraySource(13, 19, 12, 110)))
//	for v := range pipe {
//		fmt.Println(v)
//	}
//}
//
////生成随机数组
//func main11() {
//	var fileName = "/Users/xuchanghui/Desktop/data.in" //
//	count := 1000
//
//	file, _ := os.Create(fileName)
//	defer file.Close()
//
//	random := distributesort.RandomeSource(count)
//	writer := bufio.NewWriter(file)
//	distributesort.WriteSync(writer, random)
//	writer.Flush()
//	file1, _ := os.Open(fileName)
//	defer file1.Close()
//	reader := bufio.NewReader(file1)
//	myPipe := distributesort.ReadSource(reader, -1)
//	count = 0
//	for v := range myPipe {
//		count++
//		fmt.Println(v)
//		if count > 1000 {
//			break
//		}
//	}
//}

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
						shellsort.ShellSortStep(arr, v, gap)
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

func main9() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10, 0, 1000, 99, 34, 86, 73}
	fmt.Println(ShellSortGoRoutine(arr))
}

func main8() {
	t1 := time.Now()
	//const N = 6428632 // 需要开辟的内存
	const N = 6428458

	myStrs := make([]string, N)

	file, err := os.Open("/Users/xuchanghui/Desktop/CSDN-pass.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	path := "/Users/xuchanghui/Desktop/CSDN-pass-sort.txt"
	pwdFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer pwdFile.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(pwdFile)

	for i := 0; i < N; i++ {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//fmt.Println(line)
		//strs := strings.Split(line, " # ")
		//fmt.Println(strs[1])
		myStrs[i] = line

		//writer.Write([]byte(strs[1]))
		//fmt.Fprint(writer, strs[2])
	}

	//myStrs = quicksort.QuickSort(myStrs)

	used := time.Since(t1)
	fmt.Println(used)
	for i := 0; i < N; i++ {
		fmt.Fprint(writer, myStrs[i])
	}
	writer.Flush()
}

func main03() {
	//file, err := os.Open("/Users/xuchanghui/Desktop/CSDN-中文IT社区-600万.sql")
	file, err := os.Open("/Users/xuchanghui/Desktop/CSDN-pass.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//path := "/Users/xuchanghui/Desktop/CSDN-pass.txt"
	//pwdFile, err := os.Create(path)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer pwdFile.Close()

	reader := bufio.NewReader(file)
	//writer := bufio.NewWriter(pwdFile)
	i := 0
	for {
		//line, err:= reader.ReadString('\n')
		_, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//fmt.Println(line)
		//strs := strings.Split(line, " # ")
		//fmt.Println(strs[1])

		//writer.Write([]byte(strs[1]))
		//fmt.Fprint(writer, strs[2])
		i++
	}
	//writer.Flush()
	fmt.Println(i)
}

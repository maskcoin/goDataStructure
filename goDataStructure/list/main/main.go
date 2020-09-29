package main

import (
	"bufio"
	"fmt"
	"io"
	"list/doublelink"
	"list/singlelink"
	"os"
	"time"
)
//1048576
const N = 1048576

func main() {
	list := doublelink.NewDoubleList()
	node1 := doublelink.NewNode(1)
	node2 := doublelink.NewNode(2)
	node3 := doublelink.NewNode(3)
	node4 := doublelink.NewNode(4)
	node5 := doublelink.NewNode(5)
	node6 := doublelink.NewNode(6)
	node7 := doublelink.NewNode(7)
	list.InsertFront(node1)
	//fmt.Println(list)
	list.InsertFront(node2)
	//fmt.Println(list)
	list.InsertFront(node3)
	//fmt.Println(list)
	list.InsertFront(node4)
	//fmt.Println(list)
	list.InsertFront(node5)
	//fmt.Println(list)
	list.PrintFromHead()
	list.PrintFromTail()
	list.InsertNodeBack(node3, node6)
	list.InsertNodeFront(node1, node7)
	list.PrintFromHead()
	list.PrintFromTail()
	list.DeleteNode(node3)
	list.PrintFromHead()
	list.PrintFromTail()
	list.DeleteNodeFromIndex(2)
	list.PrintFromHead()
	list.PrintFromTail()
}

func main2() {
	list := singlelink.NewSingleList()
	path := "/Users/xuchanghui/Desktop/140W某信用卡购物网数据.txt"
	file, _ := os.Open(path)
	reader := bufio.NewReader(file)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		node := singlelink.NewNode(line)
		list.InsertListFront(node)
		i++
	}

	fmt.Println(i, "内存载入完成")
	fmt.Println("请输入要查询的数据")
	for {
		var data string
		fmt.Scanf("%s\n", &data)
		startTime := time.Now()
		destNode := list.FindNodeByValue(data)

		if destNode != nil {
			fmt.Println("destNode.value =", destNode.Value())
		}


		fmt.Println("本次查询使用了:", time.Since(startTime))
	}


}

func main1() {
	list := singlelink.NewSingleList()
	node1 := singlelink.NewNode(1)
	node2 := singlelink.NewNode(2)
	node3 := singlelink.NewNode(3)
	node4 := singlelink.NewNode(4)
	node5 := singlelink.NewNode(5)
	//list.InsertFront(node1)
	list.InsertListBack(node1)
	fmt.Println(list)
	//list.InsertFront(node2)
	list.InsertListBack(node2)
	fmt.Println(list)
	//list.InsertFront(node3)
	list.InsertListBack(node3)
	list.InsertNodeValueFront(2, node4)
	list.InsertNodeValueBack(2, node5)
	fmt.Println(list)
	fmt.Println(list.GetMid())
	list.Reverse()
	fmt.Println(list)

	//fmt.Println(list.GetNode(3))
	//list.DeleteNode(node5)
	//fmt.Println(list)
	//list.DeleteNodeFromIndex(2)
	//fmt.Println(list)
}

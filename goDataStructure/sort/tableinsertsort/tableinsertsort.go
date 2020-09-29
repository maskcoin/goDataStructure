package tableinsertsort

import "fmt"

const Int_Max = int(^uint(0) >> 1)

type Node struct {
	Value int
	Next  int // 下一个索引
}

type NodeList []Node

func NewNodeList(arr []int) NodeList {
	var list NodeList
	//插入第一个节点，确保第一个节点值最大，编号为1
	node := Node{
		Value: Int_Max,
		Next:  0,
	}
	list = append(list, node)
	for i := 0; i < len(arr); i++ {
		node = Node{
			Value: arr[i],
			Next:  0,
		}
		list = append(list, node)
	}
	fmt.Println(list)
	return list
}

//func (list NodeList) ListSort1() {
//	var i, low, high int
//	for i = 2; i < len(list); i++ {
//		low = 0
//		high = list[0].next
//		for list[high].value < list[i].value { // 寻找一个邻居的数据list[max] list[i]，插入list[min]
//			low = high
//			high = list[high].next
//		}
//		list[low].next = i
//		list[i].next = high // 插入数据到中间
//	}
//	fmt.Println(list)
//}

func (list NodeList) ListSort() {
	var i, front, current int
	//给环形链表的插入第一个元素
	front = 0
	current = 1
	list[front].Next = current
	for i = 2; i < len(list); i++ {
		front = 0
		current = 1
		for list[current].Value < list[i].Value {
			front = current
			current = list[current].Next
		}
		list[front].Next = i
		list[i].Next = current
	}
	fmt.Println(list)
}

func (list NodeList) Arrange() {
	//p := list[0].Next
	//for i := 1; i < len(list); i++ {
	//	for p<i {//i之前都是排序好的
	//		p = list[p].Next
	//	}
	//	q := list[p].Next
	//	if p !=i {
	//		list[p].Value, list[i].Value = list[i].Value, list[p].value
	//		list[p].next = list[i].next
	//		list[i].next = p
	//	}
	//	p = q
	//}
	//for i := 1; i < len(list); i++ {
	//	fmt.Println(list[i].value)
	//}
}

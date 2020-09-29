package main

import (
	"datastructure/day01/code/arraystack"
	"datastructure/day01/code/deque"
	"datastructure/day01/code/queue"
	"fmt"
	"math/rand"
	"time"
)

func Add(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + Add(num-1)
	}
}

func FAB(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return FAB(num-1) + FAB(num-2)
	}
}

// 判断数组是否从小到大
func IsOrder(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			return false
		}
	}
	return true
}

//洗牌算法
func ShuffleList(list []int)  {
	data := make([]int, len(list))
	copy(data, list)
	rand.Seed(time.Now().Unix())
	indexArr := rand.Perm(len(list))
	for i,k:=range indexArr {
		list[i] = data[k]
	}
}

func main() {
	list := []int{1, 9, 2, 8, 3, 7}
	fmt.Println(list)
	count := 0
	for {
		count++
		if IsOrder(list) {
			fmt.Println("排序完成")
			break
		} else {
			ShuffleList(list)
		}
	}
	fmt.Println(count)
	fmt.Println(list)
}

func main5() {
	dq := deque.NewDeque(10)
	dq.AddLeft(1)
	dq.AddLeft(2)
	dq.AddLeft(3)
	fmt.Println(dq)
}

func main4() {
	queue := queue.NewMyQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
}

func main3() {
	stack := arraystack.NewArrayStack()
	stack.Push(7)
	sum := 0

	for !stack.IsEmpty() {
		data := stack.Pop()
		if data == 1 || data == 2 {
			sum += 1
		} else {
			stack.Push(data.(int) - 1)
			stack.Push(data.(int) - 2)
		}
	}

	fmt.Println(sum)
}

func main2() {
	fmt.Println(FAB(6))
}

func main1() {
	//fmt.Println(Add(5))
	stack := arraystack.NewArrayStack()
	stack.Push(5)
	sum := 0

	for !stack.IsEmpty() {
		data := stack.Pop()
		if data == 0 {
			sum += 0
		} else {
			sum += data.(int)
			stack.Push(data.(int) - 1)
		}
	}

	fmt.Println(sum)
}

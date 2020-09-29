package main

import (
	"datastructure/link/hashmap"
	"datastructure/link/hashtable"
	"datastructure/link/queue"
	"datastructure/link/stack"
	"fmt"
	"sync"
)

func main()  {
	fileList := []string{"123","456","789"}
	hashMap := &hashmap.Ring{
		RMap:      map[uint32]string{},
		RIndexArr: nil,
		RWMutex:   sync.RWMutex{},
	}
	fmt.Println(fileList, hashMap)
	for _,v := range fileList {
		hashMap.AddNode(v)
	}
}

func main4()  {
	ht := hashtable.NewHashTable(1000)
	ht.Put("yincheng1", "123456")
	ht.Put("yincheng2", "1234567")
	ht.Put("yincheng3", "1234568")
	ht.Put("yincheng4", "1234569")
	ht.Put("yincheng5", "12345610")

	data, err := ht.Get("yincheng3")
	fmt.Println(data, err)
	err = ht.Del("yincheng3")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		data, err := ht.Get("yincheng3")
		fmt.Println(data, err)
	}
}

func main3()  {
	queue := queue.NewQueue()
	for i := 0; i < 100; i++ {
		queue.EnQueue(i)
	}

	for !queue.IsEmpty() {
		fmt.Println(queue.DeQueue())
	}
}

func main2()  {
	stack := stack.NewStack()
	for i := 0; i < 1000; i++  {
		stack.Push(i)
	}

	for  val := stack.Pop(); val != nil; val = stack.Pop() {
		fmt.Println(val)
	}
}

func main1()  {
	node1 := new(stack.Node)
	node2 := new(stack.Node)
	node3 := new(stack.Node)
	node4 := new(stack.Node)
	node1.Data = 1
	node1.PNext = node2
	node2.Data = 2
	node2.PNext = node3
	node3.Data = 3
	node3.PNext = node4
	node4.Data = 4
	fmt.Println(node1.Data)
}

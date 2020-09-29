package main

import (
	"fmt"
	"threadsafe/queue"
	"threadsafe/safemap"
	"time"
)

var smap *safemap.SyncMap
var done chan bool = make(chan bool, 2)

func writeF1() {
	keys := []string{"1", "2", "3"}
	for _, k := range keys {
		smap.Lock.Lock()
		smap.MyMap[k] = k
		smap.Lock.Unlock()
		time.Sleep(time.Second)
	}

	done <- true
}

func writeF2() {
	keys := []string{"a1", "b2", "c3"}
	for _, k := range keys {
		smap.Lock.Lock()
		smap.MyMap[k] = k
		smap.Lock.Unlock()
		time.Sleep(time.Second)
	}

	done <- true
}

func read()  {
	smap.Lock.RLock()
	fmt.Println("readLock")
	for k, v := range smap.MyMap {
		fmt.Println(k, v)
	}
	smap.Lock.RUnlock()
}


func main()  {
	smap = safemap.NewSyncMap()
	go writeF1()
	go writeF2()
	for {
		read()
		if len(done) == 2 {
			fmt.Println(smap.MyMap)
			for k, v := range smap.MyMap {
				fmt.Println(k, v)
			}
			break
		} else {
			time.Sleep(time.Second)
		}
	}
}

func main2()  {
	h:= queue.NewMinPriorityQueue()
	h.Insert(queue.NewPriorityItem(101, 11))
	h.Insert(queue.NewPriorityItem(102, 12))
	h.Insert(queue.NewPriorityItem(103, 15))
	h.Insert(queue.NewPriorityItem(104, 14))
	h.Insert(queue.NewPriorityItem(105, 13))
	h.Insert(queue.NewPriorityItem(106, 9))

	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
}

func main1() {
	h := queue.NewMinHeap()
	h.Insert(queue.Int(8))
	h.Insert(queue.Int(9))
	h.Insert(queue.Int(7))
	h.Insert(queue.Int(5))
	h.Insert(queue.Int(6))
	h.Insert(queue.Int(4))
	h.Insert(queue.Int(2))
	h.Insert(queue.Int(3))
	fmt.Println(h.Extract().(queue.Int))
	fmt.Println(h.Extract().(queue.Int))
	fmt.Println(h.Extract().(queue.Int))
	fmt.Println(h.Extract().(queue.Int))
	fmt.Println(h.Extract().(queue.Int))
}

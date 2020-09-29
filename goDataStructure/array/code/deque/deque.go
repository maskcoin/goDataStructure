package deque

import "sync"

//双端队列
type Deque struct {
	array []interface{}
	left int
	right int
	fixSize int
	lock *sync.RWMutex
}

func NewDeque(cap int) *Deque {
	if cap <= 0 {
		panic("队列容量必须大于0")
	}
	return &Deque{
		array:   make([]interface{}, cap),
		left:    0,
		right:   0,
		fixSize: cap,
		lock:    &sync.RWMutex{},
	}
}

func (dq *Deque) AddLeft(data interface{}) {
	if dq.left == dq.right && dq.left != 0 {
		panic("overflow")
	}
	dq.left--
	if dq.left == -1 {
		dq.left = dq.fixSize - 1 // 循环双端队列
	}
	dq.array[dq.left] = data
}

func (dq *Deque) AddRight(data interface{}) {
	if dq.left == dq.right {
		panic("overflow")
	}
	dq.array[dq.right] = data
	dq.right++
	if dq.right == dq.fixSize {
		dq.right = 0
	}
}

func (dq *Deque) DelLeft() interface{} {
	if dq.left == dq.right {
		panic("overflow")
	}
	data := dq.array[dq.left]
	dq.left++
	if dq.left == dq.fixSize {
		dq.left = 0
	}
	return data
}

func (dq *Deque) DelRight() interface{} {
	if dq.left == dq.right {
		panic("overflow")
	}
	dq.right--
	if dq.right == -1 {
		dq.right = dq.fixSize - 1
	}
	data := dq.array[dq.right]
	return data
}

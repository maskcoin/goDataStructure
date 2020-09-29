package queue

import "sync"

type Item interface {
	Less(obj Item ) bool
}

type Int int

func (i Int) Less(obj Item) bool  {
	return i < obj.(Int)
}

// 最小堆, 最大堆
type Heap struct {
	lock *sync.Mutex
	data []Item
	min bool
}

func NewHeap() *Heap {
	return &Heap{
		lock: new(sync.Mutex),
		data: nil,
		min:  true,
	}
}

func NewMinHeap() *Heap {
	return &Heap{
		lock: new(sync.Mutex),
		data: nil,
		min:  true,
	}
}

func NewMaxHeap() *Heap {
	return &Heap{
		lock: new(sync.Mutex),
		data: nil,
		min:  false,
	}
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Get(index int) Item {
	return h.data[index]
}

// 插入数据
func (h *Heap) Insert(it Item) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.data = append(h.data, it)
	h.Shift()
}

func (h *Heap) Less(a, b Item) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a)
	}
}

// 压缩，弹出一个
func (h *Heap) Extract() (el Item) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.Len() == 0 {
		return
	}
	el = h.data[0]
	last := h.data[h.Len()-1] // 最后一个
	if h.Len() == 1 {
		h.data = nil
		return
	}
	h.data = append([]Item{last}, h.data[1: h.Len()-1]...)
	h.Shift()
	return
}

func (h *Heap) Shift() {
	for i:= h.Len()/2 - 1; i >= 0; i-- {
		if h.Less(h.data[2*i + 1], h.data[i]) {
			h.data[i], h.data[2*i + 1] = h.data[2*i + 1], h.data[i]
		}
		if (2*i+2) < h.Len() && h.Less(h.data[2*i + 2], h.data[i]) {
			h.data[i], h.data[2*i + 2] = h.data[2*i + 2], h.data[i]
		}
	}
}
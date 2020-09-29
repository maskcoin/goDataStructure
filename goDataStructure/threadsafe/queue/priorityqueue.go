package queue

//队列中间的元素
type PriorityItem struct {
	value    interface{} // 数据
	priority int         // 优先级
}

func NewPriorityItem(value interface{}, priority int) *PriorityItem {
	return &PriorityItem{value, priority}
}

func (item *PriorityItem) Less(obj Item) bool {
	return item.priority < obj.(*PriorityItem).priority
}

// 优先队列，基于堆
type PriorityQueue struct {
	data *Heap
}

func NewMaxPriorityQueue() *PriorityQueue {
	return &PriorityQueue{data: NewMaxHeap()}
}

func NewMinPriorityQueue() *PriorityQueue {
	return &PriorityQueue{data: NewMinHeap()}
}

func (q *PriorityQueue) Len() int{
	return q.data.Len()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *PriorityQueue) Insert(el *PriorityItem) {
	q.data.Insert(el)
}

func (q *PriorityQueue) Extract() (el *PriorityItem) {
	return q.data.Extract().(*PriorityItem)
}

func (q *PriorityQueue) ChangePriority(val interface{}, priority int) bool {
	ret := true
	var storage Queue // 用一个队列来备份数据
	pop := q.Extract()
	for pop.value != val {
		if q.Len() == 0 {
			ret = false
			q.data.Insert(pop)

			for storage.Len() > 0 {
				q.data.Insert(storage.Shift().(*PriorityItem))
			}
			return ret
		}
		storage.Push(pop)
		pop = q.Extract()
	}

	pop.priority = priority
	q.data.Insert(pop)

	for storage.Len() > 0 {
		q.data.Insert(storage.Shift().(*PriorityItem))
	}

	return ret
}

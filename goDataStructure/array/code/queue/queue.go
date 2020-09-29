package queue

type Queue interface {
	Size() int
	Front() interface{}
	End() interface{}
	IsEmpty() bool
	EnQueue(data interface{})
	DeQueue() interface{}
	Clear()
}

type MyQueue struct {
	dataStore []interface{}
	theSize   int
}

func NewMyQueue() *MyQueue {
	queue := new(MyQueue)
	queue.dataStore = make([]interface{}, 0, 10)
	queue.theSize = 0
	return queue
}

func (queue *MyQueue) Size() int {
	return queue.theSize
}

func (queue *MyQueue) Front() interface{} {
	if !queue.IsEmpty() {
		return queue.dataStore[0]
	}
	return nil
}

func (queue *MyQueue) End() interface{} {
	if !queue.IsEmpty() {
		return queue.dataStore[queue.theSize-1]
	}
	return nil
}

func (queue *MyQueue) IsEmpty() bool {
	return queue.theSize == 0
}

func (queue *MyQueue) EnQueue(data interface{}) {
	queue.dataStore = append(queue.dataStore, data)
	queue.theSize++
}

func (queue *MyQueue) DeQueue() interface{} {
	if queue.IsEmpty() {
		return nil
	}

	data := queue.dataStore[0]
	if queue.theSize > 1 {
		queue.dataStore = queue.dataStore[1:queue.theSize]
	}
	queue.theSize--

	return data
}

func (queue *MyQueue) Clear() {
	queue.dataStore = make([]interface{}, 0, 10)
	queue.theSize = 0
}

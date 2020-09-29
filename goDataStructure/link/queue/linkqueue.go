package queue

type Queue interface {
	IsEmpty() bool
	EnQueue(data interface{})
	DeQueue() interface{}
	Length() int
}

type Node struct {
	data interface{}
	pNext *Node
}

type LinkQueue struct {
	front *Node
	rear *Node
}

func NewQueue() *LinkQueue {
	return &LinkQueue{}
}

func (queue *LinkQueue) IsEmpty() bool  {
	return queue.front == nil && queue.rear == nil
}

func (queue *LinkQueue) EnQueue(data interface{}) {
	node := &Node{
		data:  data,
		pNext: nil,
	}

	if queue.front == nil {
		queue.front = node
		queue.rear = node
	} else {
		queue.rear.pNext = node
		queue.rear = node
	}
}

func (queue *LinkQueue) DeQueue() interface{} {
	if queue.front == nil {
		return nil
	}

	val := queue.front.data
	if queue.front == queue.rear {
		queue.front = nil
		queue.rear = nil
	} else {
		queue.front = queue.front.pNext
	}
	return val
}

func (queue *LinkQueue) Length() int {
	p := queue.front
	length := 0
	if p != nil {
		length = 1
		for p.pNext != nil {
			p = p.pNext
			length++
		}
	}

	return length
}

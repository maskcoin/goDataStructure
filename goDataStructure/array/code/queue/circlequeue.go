package queue

const queueSize  = 100
type CircleQueue struct {
	data [queueSize]interface{}
	front int
	rear int
}

func initQueue(q *CircleQueue)  { 
	q.front = 0
	q.rear = 0
}

func (queue *CircleQueue) EnQueue(data interface{}) error {
	return nil
}

func (queue *CircleQueue) DeQueue() (interface{}, error) {
	return nil, nil
}

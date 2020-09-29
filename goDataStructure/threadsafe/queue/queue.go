package queue

import "sync"

type Queue struct {
	queue []interface{}
	len int
	lock *sync.Mutex // 锁
}

func NewQueue() *Queue {
	return &Queue{
		queue: nil,
		len:   0,
		lock:  new(sync.Mutex),
	}
}

// 解决了线程安全
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len
}

func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0
}

func (q *Queue) Shift() (el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	el, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return el
}

func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue[0]
}

func (q *Queue) Push(el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queue = append(q.queue, el)
	q.len++
}

package circle

import "fmt"

type Circle struct {
	head *Node
	tail *Node
}

type Node struct {
	num int
	next *Node
}

func NewCircle() *Circle {
	return &Circle{
		head: &Node{},
		tail: nil,
	}
}

//环链表
func (circle *Circle) AddNode(node *Node)  {
	if circle.tail == nil {
		circle.head.next = node
		node.next = circle.head.next
		circle.tail = node
	} else {
		circle.tail.next = node
		node.next = circle.head.next
		circle.tail = node
	}
}

func (circle *Circle) Print() {
	current := circle.head.next
	if current == nil {
		return
	} else {
		fmt.Println(current.num)
		for current.next != circle.head.next {
			current = current.next
			fmt.Println(current.num)
		}
	}
}



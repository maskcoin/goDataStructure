package singlelink

import (
	"fmt"
	"strings"
)

type List interface {
	//	增删查改
	GetFirstNode() *Node
	InsertListFront(node *Node)
	InsertListBack(node *Node)
	InsertNodeFront(dest *Node, node *Node) bool
	InsertNodeBack(dest *Node, node *Node) bool
	InsertNodeValueFront(dest interface{}, node *Node) bool
	InsertNodeValueBack(dest interface{}, node *Node) bool
	GetNode(index int) *Node
	DeleteNode(node *Node) bool
	DeleteNodeFromIndex(index int) bool
	FindNodeByValue(value interface{}) *Node
	String() string
	Clear()
}

type SingleList struct {
	head   *Node
	length int
}

func NewSingleList() *SingleList {
	return &SingleList{
		head:   NewNode(nil),
		length: 0,
	}
}

func (list *SingleList) GetFirstNode() *Node {
	return list.head.next
}

func (list *SingleList) InsertListFront(node *Node) {
	if node != nil {
		if list.head.next == nil {
			list.head.next = node
			node.next = nil
		} else {
			node.next = list.head.next
			list.head.next = node
		}

		list.length++
	}
}

func (list *SingleList) InsertListBack(node *Node) {
	if node != nil {
		if list.head.next == nil {
			list.head.next = node
			node.next = nil
		} else {
			tmp := list.head
			for tmp.next != nil {
				tmp = tmp.next
			}
			node.next = nil
			tmp.next = node
		}

		list.length++
	}
}

func (list *SingleList) InsertNodeFront(dest *Node, node *Node) bool {
	ret := true
	if list.head.next == nil || dest == nil || node == nil {
		ret = false
	} else {
		front := list.head
		current := list.head
		for current.next != nil && current != dest {
			front = current
			current = current.next
		}
		if current.next == nil {
			if current != dest {
				ret = false
			} else {
				front.next = node
				node.next = current
				list.length++
			}
		} else {
			front.next = node
			node.next = current
			list.length++
		}
	}

	return ret
}

func (list *SingleList) InsertNodeBack(dest *Node, node *Node) bool {
	ret := true
	if list.head.next == nil || dest == nil || node == nil {
		ret = false
	} else {
		current := list.head
		for current.next != nil && current != dest {
			current = current.next
		}
		if current.next == nil {
			if current != dest {
				ret = false
			} else {
				node.next = current.next
				current.next = node
				list.length++
			}
		} else {
			node.next = current.next
			current.next = node
			list.length++
		}
	}

	return ret
}

func (list *SingleList) InsertNodeValueFront(dest interface{}, node *Node) bool {
	ret := true
	if list.head.next == nil || dest == nil || node == nil {
		ret = false
	} else {
		front := list.head
		current := list.head
		for current.next != nil && current.value != dest {
			front = current
			current = current.next
		}
		if current.next == nil {
			if current != dest {
				ret = false
			} else {
				front.next = node
				node.next = current
				list.length++
			}
		} else {
			front.next = node
			node.next = current
			list.length++
		}
	}
	return ret
}

func (list *SingleList) InsertNodeValueBack(dest interface{}, node *Node) bool {
	ret := true
	if list.head.next == nil || dest == nil || node == nil {
		ret = false
	} else {
		current := list.head
		for current.next != nil && current.value != dest {
			current = current.next
		}
		if current.next == nil {
			if current != dest {
				ret = false
			} else {
				node.next = current.next
				current.next = node
				list.length++
			}
		} else {
			node.next = current.next
			current.next = node
			list.length++
		}
	}
	return ret
}

func (list *SingleList) GetNode(index int) *Node {
	var ret *Node

	if index <= list.length-1 {
		current := list.head

		i := 0
		for current.next != nil && i != index {
			current = current.next
			i++
		}
		ret = current.next
	}

	return ret
}

func (list *SingleList) DeleteNode(node *Node) bool {
	ret := true

	if list.head.next == nil || node == nil {
		ret = false
	} else {
		front := list.head
		current := list.head
		for current.next != nil && current != node {
			front = current
			current = current.next
		}
		if current.next == nil {
			if current != node {
				ret = false
			} else {
				front.next = nil
				list.length--
			}
		} else {
			front.next = node.next
			list.length--
		}
	}

	return ret
}

func (list *SingleList) DeleteNodeFromIndex(index int) bool {
	ret := true
	if index < 0 || index >= list.length {
		ret = false
	} else {
		if list.head.next == nil {
			ret = false
		} else {
			front := list.head
			current := list.head
			i := 0
			for current.next != nil && i != index {
				front = current
				current = current.next
				i++
			}
			if current.next == nil {
				ret = false
			} else {
				front = current
				current = current.next
				front.next = current.next
				list.length--
			}
		}
	}
	return ret
}

func (list *SingleList) String() string {
	var ret string
	tmp := list.head
	for tmp.next != nil {
		ret += fmt.Sprintf("%v-->", tmp.next.value)
		tmp = tmp.next
	}

	ret += fmt.Sprintf("%v", nil)
	return ret
}

func (list *SingleList) Clear() {
	list = NewSingleList()
}

func (list *SingleList) FindNodeByValue(value interface{}) *Node {
	var ret *Node
	if list.length > 0 {
		current := list.head
		for current.next != nil {
			strings.Contains(current.next.value.(string), value.(string))
			if current.next.value == value {
			//if strings.Contains(current.next.value.(string), value.(string)) {
				ret = current.next
				return ret
			} else {
				current = current.next
			}
		}
	}
	return ret
}

func (list *SingleList) GetMid() *Node {
	if list.head.next == nil {
		return nil
	} else {
		front := list.head
		current := list.head

		for current != nil && current.next != nil {
			front = front.next
			current = current.next.next
		}
		return front
	}
}

func (list *SingleList) Reverse() {
	if list.head.next == nil {
		return
	} else if list.head.next.next == nil {
		return
	} else {
		first := list.head.next
		second := first.next
		if second.next == nil {
			second.next = first
			first.next = nil
			list.head.next = second
		} else {
			third := second.next
			first.next = nil
			second.next = first
			for third.next != nil {
				first = second
				second = third
				third = third.next
				second.next = first
			}
			third.next = second
			list.head.next = third
		}
	}
}
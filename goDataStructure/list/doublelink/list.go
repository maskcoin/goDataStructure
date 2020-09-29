package doublelink

import "fmt"

type list interface {
	GetLength() int
	GetFirstNode() *Node
	InsertFront(node *Node)
	InsertBack(node *Node)
	String() string
	PrintFromHead()
	PrintFromTail()
	InsertNodeFront(dest *Node, node *Node) bool
	InsertNodeBack(dest *Node, node *Node) bool
	InsertNodeValueFront(dest interface{}, node *Node) bool
	InsertNodeValueBack(dest interface{}, node *Node) bool
	GetNode(index int) *Node
	DeleteNode(node *Node) bool
	DeleteNodeFromIndex(index int) bool
}

type DoubleList struct {
	head   *Node
	length int
}

func NewDoubleList() *DoubleList {
	return &DoubleList{
		head:   NewNode(nil),
		length: 0,
	}
}

func (list *DoubleList) GetLength() int {
	return list.length
}

func (list *DoubleList) GetFirstNode() *Node {
	return list.head.next
}

func (list *DoubleList) InsertFront(node *Node) {
	if node != nil {
		if list.head.next == nil {
			list.head.next = node
			node.next = nil
			node.prev = nil
		} else {
			node.next = list.head.next
			list.head.next.prev = node
			list.head.next = node
			node.prev = nil
		}

		list.length++
	}
}

func (list *DoubleList) InsertBack(node *Node) {
	if node != nil {
		if list.head.next == nil {
			list.head.next = node
			node.next = nil
			node.prev = nil
		} else {
			tmp := list.head
			for tmp.next != nil {
				tmp = tmp.next
			}
			node.next = nil
			tmp.next = node
			node.prev = tmp
		}

		list.length++
	}
}

func (list *DoubleList) String() string {
	var ret string
	tmp := list.head
	for tmp.next != nil {
		ret += fmt.Sprintf("%v-->", tmp.next.value)
		tmp = tmp.next
	}

	ret += fmt.Sprintf("%v", nil)
	return ret
}

func (list *DoubleList) PrintFromHead() {
	var ret string
	tmp := list.head
	for tmp.next != nil {
		ret += fmt.Sprintf("%v-->", tmp.next.value)
		tmp = tmp.next
	}

	ret += fmt.Sprintf("%v", nil)
	fmt.Println(ret)
}

func (list *DoubleList) PrintFromTail() {
	var ret string
	tmp := list.head
	for tmp.next != nil {
		//ret += fmt.Sprintf("%v-->", tmp.next.value)
		tmp = tmp.next
	}

	for tmp.prev != nil {
		ret += fmt.Sprintf("%v-->", tmp.value)
		tmp = tmp.prev
	}
	ret += fmt.Sprintf("%v-->", tmp.value)
	ret += fmt.Sprintf("%v", nil)
	fmt.Println(ret)
}

func (list *DoubleList) InsertNodeFront(dest *Node, node *Node) bool {
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
				if front != list.head {
					node.prev = front
				} else {
					node.prev = nil
				}
				node.next = current
				current.prev = node
				list.length++
			}
		} else {
			front.next = node
			if front != list.head {
				node.prev = front
			} else {
				node.prev = nil
			}
			node.next = current
			current.prev = node
			list.length++
		}
	}

	return ret
}

func (list *DoubleList) InsertNodeBack(dest *Node, node *Node) bool {
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
				node.prev = current
				list.length++
			}
		} else {
			node.next = current.next
			current.next.prev = node
			current.next = node
			node.prev = current
			list.length++
		}
	}

	return ret
}

func (list *DoubleList) InsertNodeValueFront(dest interface{}, node *Node) bool {
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
			if current.value != dest {
				ret = false
			} else {
				front.next = node
				if front != list.head {
					node.prev = front
				} else {
					node.prev = nil
				}
				node.next = current
				current.prev = node
				list.length++
			}
		} else {
			front.next = node
			if front != list.head {
				node.prev = front
			} else {
				node.prev = nil
			}
			node.next = current
			current.prev = node
			list.length++
		}
	}

	return ret
}

func (list *DoubleList) InsertNodeValueBack(dest interface{}, node *Node) bool {
	ret := true
	if list.head.next == nil || dest == nil || node == nil {
		ret = false
	} else {
		current := list.head
		for current.next != nil && current.value != dest {
			current = current.next
		}
		if current.next == nil {
			if current.value != dest {
				ret = false
			} else {
				node.next = current.next
				current.next = node
				node.prev = current
				list.length++
			}
		} else {
			node.next = current.next
			current.next.prev = node
			current.next = node
			node.prev = current
			list.length++
		}
	}

	return ret
}

func (list *DoubleList) GetNode(index int) *Node {
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

func (list *DoubleList) DeleteNode(node *Node) bool {
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
			if node.next != nil {
				if front != list.head {
					node.next.prev = front
				} else {
					node.next.prev = nil
				}
			}
			list.length--
		}
	}

	return ret
}

func (list *DoubleList) DeleteNodeFromIndex(index int) bool {
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
				if current.next != nil {
					if front != list.head {
						current.next.prev = front
					} else {
						current.next.prev = nil
					}
				}
				list.length--
			}
		}
	}
	return ret
}




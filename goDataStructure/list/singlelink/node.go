package singlelink

type Node struct {
	value interface{}
	next  *Node
}

func NewNode(data interface{}) *Node {
	return &Node{data, nil}
}

func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) Next() *Node {
	return node.next
}
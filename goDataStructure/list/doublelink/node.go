package doublelink

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		value: data,
		prev:  nil,
		next:  nil,
	}
}

func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) Prev() *Node {
	return node.prev
}

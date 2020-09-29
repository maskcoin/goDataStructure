package stack

type Stack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

type Node struct {
	Data  interface{}
	PNext *Node
}

func NewStack() *Node {
	return &Node{}
}

func (node *Node) IsEmpty() bool {
	return node.PNext == nil
}

func (node *Node) Push(data interface{}) {
	newNode := &Node{
		Data:  data,
	}

	newNode.PNext = node.PNext
	node.PNext = newNode
}

func (node *Node) Pop() interface{} {
	if node.IsEmpty() {
		return nil
	}
	val := node.PNext.Data
	node.PNext = node.PNext.PNext
	return val
}

func (node *Node) Length() int {
	p := node
	length := 0
	for p.PNext != nil {
		p = p.PNext
		length++
	}
	return length
}

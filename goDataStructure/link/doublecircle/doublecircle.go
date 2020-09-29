package doublecircle

type Node struct {
	Id   int
	Data interface{}
	Prev *Node
	Next *Node
}

type DoubleLinkList struct {
	Head *Node
	Tail *Node
}

func NewDoubleLinkList() *DoubleLinkList {
	return &DoubleLinkList{
		Head: &Node{},
		Tail: nil,
	}
}

func (list *DoubleLinkList) IsEmpty() bool {
	return list.Head.Next == nil && list.Tail == nil
}




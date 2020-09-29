package binarytree

import (
	"datastructure/link/queue"
	"datastructure/link/stack"
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
	Size int
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) GetSize() int {
	return t.Size
}

func (t *BinaryTree) IsEmpty() bool {
	return t.Size == 0
}

//返回值是一颗新树
func (node *Node) Add(data int) *Node {
	if node == nil {
		node = &Node{
			Data: data,
		}
	} else {
		if data < node.Data {
			node.Left = node.Left.Add(data)
		} else {
			node.Right = node.Right.Add(data)
		}
	}

	return node
}

func (t *BinaryTree) Add(data int) {
	t.Root = t.Root.Add(data)
	t.Size++
	return
}

func (node *Node) IsIn(data int) bool {
	ret := false
	if node != nil {
		if data == node.Data {
			ret = true
		} else if data < node.Data {
			ret = node.Left.IsIn(data)
		} else {
			ret = node.Right.IsIn(data)
		}
	}
	return ret
}

func (t *BinaryTree) IsIn(data int) bool {
	return t.Root.IsIn(data)
}

func (node *Node) FindMax() (ret int) {
	if node == nil {
		panic("二叉树为空")
	} else {
		if node.Right == nil {
			ret = node.Data
		} else {
			ret = node.Right.FindMax()
		}
	}
	return
}

func (node *Node) FindMin() (ret int) {
	if node == nil {
		panic("二叉树为空")
	} else {
		if node.Left == nil {
			ret = node.Data
		} else {
			ret = node.Left.FindMin()
		}
	}
	return ret
}

func (t *BinaryTree) FindMax() int {
	return t.Root.FindMax()
}

func (t *BinaryTree) FindMin() int {
	return t.Root.FindMin()
}

func (node *Node) PreOrder() {
	if node != nil {
		fmt.Print(node.Data, " ")
		node.Left.PreOrder()
		node.Right.PreOrder()
	}
}

func (t *BinaryTree) PreOrder() {
	t.Root.PreOrder()
}

func (node *Node) InOrder() {
	if node != nil {
		node.Left.InOrder()
		fmt.Print(node.Data, " ")
		node.Right.InOrder()
	}
}

func (t *BinaryTree) InOrder() {
	t.Root.InOrder()
}

func (node *Node) PostOrder() {
	if node != nil {
		node.Left.PostOrder()
		node.Right.PostOrder()
		fmt.Print(node.Data, " ")
	}
}

func (t *BinaryTree) PostOrder() {
	t.Root.PostOrder()
}

func (t *BinaryTree) RemoveMin() (ret int) {
	if t.Root == nil {
		panic("error")
	} else {
		if t.Root.Left == nil {
			ret = t.Root.Data
			t.Root = t.Root.Right
		} else {
			cur := t.Root
			var pre *Node
			for cur.Left != nil {
				pre = cur
				cur = cur.Left
			}
			ret = cur.Data
			pre.Left = nil
		}
	}
	return
}

func (t *BinaryTree) RemoveMax() (ret int) {
	if t.Root == nil {
		panic("error")
	} else {
		if t.Root.Right == nil {
			ret = t.Root.Data
			t.Root = t.Root.Left
		} else {
			cur := t.Root
			var pre *Node
			for cur.Right != nil {
				pre = cur
				cur = cur.Right
			}
			ret = cur.Data
			pre.Right = nil
		}
	}
	return
}

// 返回一颗树
func (node *Node) Remove(data int) *Node {
	if node != nil {
		if node.Data == data {
			//删除该节点
			if node.Right == nil {
				node = node.Left
			} else {
				right := node.Right
				cur := right
				for cur.Left != nil {
					cur = cur.Left
				}
				cur.Left = node.Left
				node = right
			}
		} else if data < node.Data {
			node.Left = node.Left.Remove(data)
		} else {
			node.Right = node.Right.Remove(data)
		}
	}

	return node
}

func (t *BinaryTree) Remove(data int) *Node {
	return t.Root.Remove(data)
}

func (node *Node) PreOrderWithStack(stack stack.Stack) {
	if node != nil {
		stack.Push(node)
		for !stack.IsEmpty() {
			n := stack.Pop()
			fmt.Print(n.(*Node).Data, " ")
			if n.(*Node).Right != nil {
				stack.Push(n.(*Node).Right)
			}
			if n.(*Node).Left != nil {
				stack.Push(n.(*Node).Left)
			}
		}
	}
}

func (t *BinaryTree) PreOrderWithStack(stack stack.Stack) {
	t.Root.PreOrderWithStack(stack)
}

func (node *Node) LevelShow(queue queue.Queue) {
	if node != nil {
		queue.EnQueue(node)
		for !queue.IsEmpty() {
			n := queue.DeQueue()
			fmt.Print(n.(*Node).Data, " ")
			if n.(*Node).Left != nil {
				queue.EnQueue(n.(*Node).Left)
			}
			if n.(*Node).Right != nil {
				queue.EnQueue(n.(*Node).Right)
			}
		}
	}
}

func (t *BinaryTree) LevelShow(queue queue.Queue) {
	t.Root.LevelShow(queue)
}

func (node *Node) FindLowerAncestor(n1, n2 *Node) (ret *Node) {
	if node != nil {
		if node == n1 || node == n2 {
			ret = node
		} else {
			if node.Left == n1 || node.Right == n2 {
				ret = node
			} else {
				ret = node.Left.FindLowerAncestor(n1, n2)
				if ret == nil {
					ret = node.Right.FindLowerAncestor(n1, n2)
				}
			}
		}
	}
	return
}

func (t *BinaryTree) FindLowerAncestor(n1, n2 *Node) (ret *Node) {
	ret = t.Root.FindLowerAncestor(n1, n2)
	return
}

func (node *Node) Depth() (depth int) {
	if node == nil {
		depth = 0
	} else {
		depth++
		depth1 := node.Left.Depth()
		depth2 := node.Right.Depth()
		if depth1 > depth2 {
			depth += depth1
		} else {
			depth += depth2
		}
	}
	return
}

func (t *BinaryTree) Depth() int {
	return t.Root.Depth()
}

func (node *Node) NumNode() (ret int) {
	if node != nil {
		ret++
		ret += node.Left.NumNode()
		ret += node.Right.NumNode()
	}
	return
}

func (t *BinaryTree) NumNode() (ret int) {
	ret = t.Root.NumNode()
	return
}

package avltree

import (
	"errors"
	"fmt"
)

//AVL树适用于没有删除的情况
//红黑树的增删查改最优
type Node struct {
	Data   interface{}
	Left   *Node
	Right  *Node
	Height int
}

//Comparator 函数指针类型
type Comparator func(a, b interface{}) int

//compare 函数指针
var compare Comparator

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data:   data,
		Height: 1,
	}
}

func NewAVLTree(data interface{}, myF Comparator) (node *Node, err error) {
	if data == nil && myF == nil {
		err = errors.New("参数不能为空")
		return
	}
	node = NewNode(data)
	compare = myF
	return
}

func (node *Node) GetData() (ret interface{}) {
	if node != nil {
		ret = node.Data
	}
	return
}

func (node *Node) SetData(data interface{}) {
	if node != nil {
		node.Data = data
	}
}

func (node *Node) GetLeft() (ret *Node) {
	if node != nil {
		ret = node.Left
	}
	return
}

func (node *Node) GetHeight() (ret int) {
	if node != nil {
		ret = node.Height
	} else {
		ret = 0
	}
	return
}

func (node *Node) GetRight() (ret *Node) {
	if node != nil {
		ret = node.Right
	}
	return
}

func (node *Node) FindMin() (ret *Node) {
	if node.Left == nil {
		ret = node
	} else {
		ret = node.Left.FindMin()
	}
	return
}

func (node *Node) FindMax() (ret *Node) {
	if node.Right == nil {
		ret = node
	} else {
		ret = node.Right.FindMax()
	}
	return
}

func (node *Node) Find(data interface{}) (ret *Node) {
	if node != nil {
		switch compare(data, node.Data) {
		case -1:
			if node.Left != nil {
				ret = node.Left.Find(data)
			}
		case 1:
			if node.Right != nil {
				ret = node.Right.Find(data)
			}
		case 0:
			ret = node
		}
	}

	return
}

func AddValues(values []interface{}, node *Node) []interface{} {
	if node != nil {
		values = AddValues(values, node.Left)
		values = append(values, node.Data)
		fmt.Println(node.Data, node.Height)
		values = AddValues(values, node.Right)
	}

	return values
}

func (node *Node) GetAll() []interface{} {
	var values []interface{}
	return AddValues(values, node)
}

//左旋，逆时针
func (node *Node) LeftRotate() (ret *Node) {
	ret = node.Right
	node.Right = ret.Left
	ret.Left = node
	node.Height = Max(node.Left.GetHeight(), node.Right.GetHeight()) + 1
	ret.Height = Max(ret.Left.GetHeight(), ret.Right.GetHeight()) + 1
	return
}

//右旋，顺时针
func (node *Node) RightRotate() (ret *Node) {
	ret = node.Left
	node.Left = ret.Right
	ret.Right = node
	node.Height = Max(node.Left.GetHeight(), node.Right.GetHeight()) + 1
	ret.Height = Max(ret.Left.GetHeight(), ret.Right.GetHeight()) + 1
	return
}

//两次左旋

//两次右旋

//先左旋，再右旋
func (node *Node) LeftThenRightRotate() (ret *Node) {
	left := node.Left.LeftRotate()
	node.Left = left
	ret = node.RightRotate()
	return
}

//先右旋，再左旋
func (node *Node) RightThenLeftRotate() (ret *Node) {
	right := node.Right.RightRotate()
	node.Right = right
	ret = node.LeftRotate()
	return
}

//自动处理不平衡，差距为1平衡，差距为2不平衡
func (node *Node) Adjust() *Node {
	if node.Right.GetHeight()-node.Left.GetHeight() == 2 {
		right := node.Right
		if right.Right.GetHeight() > right.Left.GetHeight() {
			node = node.LeftRotate()
		} else {
			node = node.RightThenLeftRotate() //极限思维
		}
	} else if node.Left.GetHeight()-node.Right.GetHeight() == 2 {
		left := node.Left
		if left.Left.GetHeight() > left.Right.GetHeight() {
			node = node.RightRotate()
		} else {
			node = node.LeftThenRightRotate()
		}
	}

	return node
}

func (node *Node) Insert(data interface{}) *Node {
	if node == nil {
		node = &Node{
			Data:   data,
			Height: 1,
		}
		return node
	} else {
		switch compare(data, node.Data) {
		case -1:
			node.Left = node.Left.Insert(data)
			node = node.Adjust()

		case 1:
			node.Right = node.Right.Insert(data)
			node = node.Adjust()

		case 0:
			fmt.Println("数据已经存在")
		}
	}

	node.Height = Max(node.Left.GetHeight(), node.Right.GetHeight()) + 1
	return node
}

func (node *Node) Delete(data interface{}) (ret *Node) {
	if node != nil {
		switch compare(data, node.Data) {
		case -1:
			if node.Left != nil {
				node.Left = node.Left.Delete(data)
				ret = node.Adjust()
			}
		case 1:
			if node.Right != nil {
				node.Right = node.Right.Delete(data)
				ret = node.Adjust()
			}
		case 0:
			if node.Right != nil {
				ret = node.Right
				min := ret.FindMin()
				min.Left = node.Left
			} else {
				ret = node.Left
			}
			ret = ret.Adjust()
		}
	}
	return
}

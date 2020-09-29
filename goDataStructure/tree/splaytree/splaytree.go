package splaytree

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	Key    interface{}
	Value  interface{}
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewNNode(key, value interface{}) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

type SplayTree interface {
	SetRoot(node *Node)
	GetRoot() *Node
	Order(key1, key2 interface{}) int //排序
}

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	return &Tree{Root: nil}
}

func (t *Tree) SetRoot(node *Node) {
	t.Root = node
}

func (t *Tree) GetRoot() *Node {
	return t.Root
}

func (t *Tree) Order(key1, key2 interface{}) (ret int) {
	if key1.(int) < key2.(int) {
		ret = -1
	} else if key1.(int) == key2.(int) {
		ret = 0
	} else {
		ret = 1
	}

	return
}

func (node *Node) Order(key1, key2 interface{}) (ret int) {
	if key1.(int) < key2.(int) {
		ret = -1
	} else if key1.(int) == key2.(int) {
		ret = 0
	} else {
		ret = 1
	}

	return
}

func (node *Node) SearchNode(key interface{}) (ret *Node) {
	if node != nil {
		if node.Order(node.Key, key) == 0 {
			ret = node

			return
		} else if node.Order(node.Key, key) < 0 {
			ret = node.Right.SearchNode(key)
		} else {
			ret = node.Left.SearchNode(key)
		}
	}

	return
}

func (t *Tree) SearchNode(key interface{}) (ret *Node) {
	ret = t.Root.SearchNode(key)
	t.Splay(ret)
	return
}

func (t *Tree) Insert(key, value interface{}) (*Node, error) {
	cur := t.Root
	pre := cur
	for cur != nil {
		if cur.Order(cur.Key, key) < 0 {
			pre = cur
			cur = cur.Right
		} else if cur.Order(cur.Key, key) == 0 {
			return nil, errors.New("要插入的数据已经存在")
		} else {
			pre = cur
			cur = cur.Left
		}
	}

	newNode := NewNNode(key, value)

	if pre != nil {
		if pre.Order(pre.Key, key) < 0 {
			pre.Right = newNode
			newNode.Parent = pre
		} else {
			pre.Left = newNode
			newNode.Parent = pre
		}
	} else {
		t.Root = newNode
	}

	if newNode.Parent != nil {
		t.Splay(newNode)
	}

	return newNode, nil
}

func Swap(n1, n2 *Node) {
	n1.Key, n2.Key = n2.Key, n1.Key
	n1.Value, n2.Value = n2.Value, n1.Value
}

func (t *Tree) Delete(key interface{}) error {
	cur := t.Root
	pre := cur
	for cur != nil {
		if cur.Order(cur.Key, key) < 0 {
			pre = cur
			cur = cur.Right
		} else if cur.Order(cur.Key, key) == 0 {
			if cur.Right == nil && cur.Left == nil {
				if cur.Parent == nil {
					t.Root = nil
				} else {
					if cur.Parent.Left == cur {
						cur.Parent.Left = nil
					} else {
						cur.Parent.Right = nil
					}
				}
			} else if cur.Right != nil {
				successor := t.Successor(cur)
				Swap(cur, successor)

				if successor.Parent.Left == successor {
					successor.Parent.Left = nil
				} else {
					successor.Parent.Right = nil
				}

				t.Splay(cur)
			} else {
				predecessor := t.Predecessor(cur)
				Swap(cur, predecessor)

				if predecessor.Parent.Left == predecessor {
					predecessor.Parent.Left = nil
				} else {
					predecessor.Parent.Right = nil
				}

				t.Splay(cur)
			}

			return nil
		} else {
			pre = cur
			cur = cur.Left
		}
	}

	if pre != nil {
		return errors.New("你要删除的key不存在")
	} else {
		return errors.New("你要删除的是一颗空树")
	}
}

func (t *Tree) Predecessor(node *Node) (ret *Node) {
	if node != nil {
		if node.Left != nil {
			cur := node.Left
			pre := cur
			cur = cur.Right
			for cur != nil {
				pre = cur
				cur = cur.Right
			}
			ret = pre
		} else {
			cur := node
			parent := cur.Parent
			for parent != nil && cur != parent.Right {
				cur = parent
				parent = cur.Parent
			}

			if parent != nil {
				ret = parent
			}
		}
	}
	return
}

func (t *Tree) Successor(node *Node) (ret *Node) {
	if node != nil {
		if node.Right != nil {
			cur := node.Right
			pre := cur
			cur = cur.Left
			for cur != nil {
				pre = cur
				cur = cur.Left
			}
			ret = pre
		} else {
			cur := node
			parent := cur.Parent
			for parent != nil && cur != parent.Left {
				cur = parent
				parent = cur.Parent
			}

			if parent != nil {
				ret = parent
			}
		}
	}
	return
}

func (t *Tree) PrintTree() {
	t.Root.Print(0)
}

func (node *Node) Print(level int) {
	if node != nil {
		fmt.Println(strings.Repeat("-", 2*level), node.Key, node.Value)
		node.Left.Print(level + 1)
		node.Right.Print(level + 1)
	}
}

func (t *Tree) Splay(n *Node) {
	for n != t.Root {
		if n.Parent == t.Root {
			if n.Parent.Left == n {
				t.Zig(n)
			} else {
				t.Zag(n)
			}
		} else {
			if n.Parent.Left == n && n.Parent.Parent.Left == n.Parent {
				t.ZigZig(n)
			} else if n.Parent.Right == n && n.Parent.Parent.Right == n.Parent {
				t.ZagZag(n)
			} else if n.Parent.Left == n && n.Parent.Parent.Right == n.Parent {
				t.ZigZag(n)
			} else if n.Parent.Right == n && n.Parent.Parent.Left == n.Parent {
				t.ZagZig(n)
			}
		}
	}
}

//Zag表示左旋，Zig表示右旋
func (t *Tree) Zag(n *Node) {
	if n != t.Root {
		p := n.Parent

		if p == t.Root {
			p.Right = n.Left
			if n.Left != nil {
				n.Left.Parent = p
			}
			n.Left = p
			p.Parent = n
			n.Parent = nil
			t.Root = n
		} else {
			g := p.Parent

			p.Right = n.Left
			if n.Left != nil {
				n.Left.Parent = p
			}

			n.Left = p
			if g.Left == p {
				g.Left = n
			} else {
				g.Right = n
			}
			n.Parent = g
			p.Parent = n
		}
	}
}

func (t *Tree) Zig(n *Node) {
	if n != t.Root {
		p := n.Parent

		if p == t.Root {
			p.Left = n.Right
			if n.Right != nil {
				n.Right.Parent = p
			}
			n.Right = p
			p.Parent = n
			n.Parent = nil
			t.Root = n
		} else {
			g := p.Parent

			p.Left = n.Right
			if n.Right != nil {
				n.Right.Parent = p
			}

			n.Right = p
			if g.Right == p {
				g.Right = n
			} else {
				g.Left = n
			}
			n.Parent = g
			p.Parent = n
		}
	}
}

func (t *Tree) ZagZag(n *Node) {
	t.Zag(n.Parent)
	t.Zag(n)
}

func (t *Tree) ZigZig(n *Node) {
	t.Zig(n.Parent)
	t.Zig(n)
}

func (t *Tree) ZagZig(n *Node) {
	t.Zag(n)
	t.Zig(n)
}

func (t *Tree) ZigZag(n *Node) {
	t.Zig(n)
	t.Zag(n)
}

package rbtree

const (
	RED   = false
	BLACK = true
)

type Item interface {
	Less(b Item) bool
}

type Node struct {
	Item
	Color  bool
	Left   *Node
	Right  *Node
	Parent *Node
}

type RBTree struct {
	Root  *Node
	NIL   *Node
	Count uint
}

func (t *RBTree) NewNode(data Item) *Node {
	return &Node{
		Item:   data,
		Color:  RED,
		Left:   t.NIL,
		Right:  t.NIL,
		Parent: nil,
	}
}

func NewRBTree() *RBTree {
	node := &Node{
		Color: BLACK,
	}
	return &RBTree{
		Root:  node,
		NIL:   node,
		Count: 0,
	}
}

func (t *RBTree) Len() uint {
	return t.Count
}

func (t *RBTree) FindMax(node *Node) (ret *Node) {
	if node == t.NIL {
		ret = t.NIL
	} else {
		for node.Right != t.NIL {
			node = node.Right
		}
		ret = node
	}

	return
}

func (t *RBTree) FindMin(node *Node) (ret *Node) {
	if node == t.NIL {
		ret = t.NIL
	} else {
		for node.Left != t.NIL {
			node = node.Left
		}
		ret = node
	}

	return
}

func (t *RBTree) Find(node *Node) (ret *Node) {
	if node == t.NIL {
		ret = t.NIL
	} else {
		cur := t.Root
		for cur != t.NIL {
			if cur.Item == node.Item {
				break
			} else if node.Item.Less(cur.Item) {
				cur = cur.Left
			} else {
				cur = cur.Right
			}
		}
		ret = cur
	}

	return
}

func (t *RBTree) FindByValue(value Item) (ret *Node) {
	cur := t.Root

	for cur != t.NIL {
		if cur.Item == value {
			break
		} else if value.Less(cur.Item) {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	ret = cur

	return
}

func (t *RBTree) LeftRotate(node *Node) {
	if node != nil && node != t.NIL {
		if node.Right != t.NIL {
			right := node.Right
			node.Right = right.Left
			if right.Left != t.NIL {
				right.Left.Parent = node
			}

			right.Parent = node.Parent
			if node.Parent == nil {
				//根节点
				t.Root = right
			} else if node == node.Parent.Left {
				node.Parent.Left = right
			} else {
				node.Parent.Right = right
			}

			right.Left = node
			node.Parent = right
		}
	}
}

func (t *RBTree) RightRotate(node *Node) {
	if node != t.NIL {
		if node.Left != t.NIL {
			left := node.Left // 把node的左孩子备份
			node.Left = left.Right //改变node的左孩子
			if left.Right != t.NIL {
				left.Right.Parent = node
			}

			left.Parent = node.Parent //备份node的parent
			if node.Parent == nil {
				t.Root = left
			} else if node.Parent.Left == node {
				node.Parent.Left = left
			} else {
				node.Parent.Right = left
			}

			left.Right = node
			node.Parent = left
		}
	}
}

func (t *RBTree) Insert(node *Node) {
	if node != nil {
		cur := t.Root
		pre := cur
		//寻找插入位置
		for cur != t.NIL {
			pre = cur
			if node.Item.Less(cur.Item) {
				cur = cur.Left
			} else if cur.Item.Less(node.Item) {
				cur = cur.Right
			} else { //相等
				return
			}
		}

		if pre == t.NIL {
			node.Color = BLACK
			t.Root = node
			t.Count++
		} else {
			node.Parent = pre
			if node.Item.Less(pre.Item) {
				pre.Left = node
			} else {
				pre.Right = node
			}
			t.Count++
			t.FixAfterInsert2(node)
		}
	}
}

//插入之后，旋转和变色的调整
//1、2-3-4树：新增元素+2节点合并（节点中只有1个元素）=3节点（节点中有2个元素）
//	红黑树：新增一个红色节点+黑色父亲节点=上黑下红-------------不要调整
//2、2-3-4树：新增元素+3节点合并（节点中有2个元素）=4节点（节点中有3个元素）
//	这里有4种小情况（左3，右3，还有2个左中右不需要调整）-----左3，右3需要通过旋转调整，其余2个不需要调整
//	红黑树：新增红色节点+上黑下红=调整后最终结果应该是中间节点是黑色，两边节点都是红色（3节点）
//3、2-3-4树：新增一个元素+4节点合并=原来的4节点分裂，中间元素升级为父节点，新增元素与剩下的其中一个合并
//	红黑树：新增红色节点+爷爷节点黑色，父节点和叔叔节点都是红色=调整后，爷爷节点变红，父亲和叔叔变黑，如果爷爷是根节点，则再变黑

//开始玩魔方
//本质上父节点是黑色就不需要调整
func (t *RBTree) FixAfterInsert(node *Node) { //默认插入节点的颜色是红色
	for node != nil && node != t.Root && node.Parent.Color == RED {
		if node.Parent == node.Parent.Parent.Left { //左3
			//叔叔节点
			pr := node.Parent.Parent.Right
			if pr.Color == RED { //说明新增元素+4节点合并，对应上面注释的第3种情况，需要变色
				node.Parent.Parent.Color = RED
				node.Parent.Color = BLACK
				pr.Color = BLACK
				node = node.Parent.Parent
			} else {
				//说明新增元素+3节点合并，对应上面注释的第2种情况
				if node.Parent.Item.Less(node.Item) {
					p := node.Parent
					t.LeftRotate(node.Parent)
					node = p
				}
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				t.RightRotate(node.Parent.Parent)
			}
		} else { //右3
			//叔叔节点
			pr := node.Parent.Parent.Left
			if pr.Color == RED { //说明新增元素+4节点合并，对应上面注释的第3种情况，需要变色
				node.Parent.Parent.Color = RED
				node.Parent.Color = BLACK
				pr.Color = BLACK
				node = node.Parent.Parent
			} else {
				//说明新增元素+3节点合并，对应上面注释的第2种情况
				if node.Item.Less(node.Parent.Item) {
					p := node.Parent
					t.RightRotate(node.Parent)
					node = p
				}
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				t.LeftRotate(node.Parent.Parent)
			}
		}
	}
	t.Root.Color = BLACK
}

func (t *RBTree) FixAfterInsert2(node *Node) { //默认插入节点的颜色是红色
	if node != nil && node != t.Root && node.Parent.Color == RED {
		if node.Parent == node.Parent.Parent.Left { //左3
			//叔叔节点
			pr := node.Parent.Parent.Right
			if pr.Color == RED { //说明新增元素+4节点合并，对应上面注释的第3种情况，需要变色
				node.Parent.Parent.Color = RED
				node.Parent.Color = BLACK
				pr.Color = BLACK
				node = node.Parent.Parent
				t.FixAfterInsert2(node)
			} else {
				//说明新增元素+3节点合并，对应上面注释的第2种情况
				if node.Parent.Item.Less(node.Item) {
					p := node.Parent
					t.LeftRotate(node.Parent)
					node = p
				}
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				t.RightRotate(node.Parent.Parent)
			}
		} else { //右3
			//叔叔节点
			pr := node.Parent.Parent.Left
			if pr.Color == RED { //说明新增元素+4节点合并，对应上面注释的第3种情况，需要变色
				node.Parent.Parent.Color = RED
				node.Parent.Color = BLACK
				pr.Color = BLACK
				node = node.Parent.Parent
				t.FixAfterInsert2(node)
			} else {
				//说明新增元素+3节点合并，对应上面注释的第2种情况
				if node.Item.Less(node.Parent.Item) {
					p := node.Parent
					t.RightRotate(node.Parent)
					node = p
				}
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				t.LeftRotate(node.Parent.Parent)
			}
		}
	}
	t.Root.Color = BLACK
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

func (t *RBTree) Depth(node *Node) (depth int) {
	if t.Root != t.NIL && node != nil && node != t.NIL {
		depth1 := t.Depth(node.Left) + 1
		depth2 := t.Depth(node.Right) + 1
		if depth1 < depth2 {
			depth = depth2
		} else {
			depth = depth1
		}
	}
	return
}

//近似查找，找一个<=node.Item的节点
func (t *RBTree) Search(data Item) (ret *Node) {
	cur := t.Root

	for cur != t.NIL {
		if data.Less(cur.Item) {
			cur = cur.Left
		} else if cur.Item.Less(data) {
			cur = cur.Right
		} else {
			ret = cur
			return
		}
	}

	return
}

func (t *RBTree) Predecessor(node *Node) (ret *Node) {
	if node != nil {
		if node.Left != t.NIL {
			cur := node.Left
			pre := cur
			cur = cur.Right
			for cur != t.NIL {
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

func (t *RBTree) Successor(node *Node) (ret *Node) {
	if node != nil {
		if node.Right != t.NIL {
			cur := node.Right
			pre := cur
			cur = cur.Left
			for cur != t.NIL {
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

//删除操作：
//删除节点，可以转化为删除2-3-4树的叶子节点
func (t *RBTree) Delete(data Item) {
	n := t.Search(data)

	if n != nil {
		//3、n节点有2个孩子，最终转换为1，2情况
		if n.Left != t.NIL && n.Right != t.NIL {
			successor := t.Successor(n)
			n.Item = successor.Item
			n = successor
		}

		//2、删除的节点，只有一个子节点，那么用子节点来替代
		if n.Right != t.NIL {
			replacement := n.Right
			replacement.Parent = n.Parent
			if n.Parent != nil {
				if n == n.Parent.Left {
					n.Parent.Left = replacement
				} else {
					n.Parent.Right = replacement
				}
			} else {
				t.Root = replacement
				t.Root.Color = BLACK
				return
			}
			////删除完之后需要调整平衡
			if n.Color == BLACK {
				t.FixAfterDelete2(replacement)
			}
		} else if n.Left != t.NIL {
			replacement := n.Left
			replacement.Parent = n.Parent
			if n.Parent != nil {
				if n == n.Parent.Right {
					n.Parent.Right = replacement
				} else {
					n.Parent.Left = replacement
				}
			} else {
				t.Root = replacement
				t.Root.Color = BLACK
				return
			}
			//删除完之后需要调整平衡
			if n.Color == BLACK {
				t.FixAfterDelete2(replacement)
			}
		} else {
			//1、删除叶子节点，直接删除
			//删除节点是根节点
			if n.Parent == nil {
				t.Root = t.NIL
			} else {
				if n.Color == BLACK {
					t.FixAfterDelete2(n) //调整后再删除
				}
				p := n.Parent
				if n == p.Left {
					p.Left = t.NIL
				} else {
					p.Right = t.NIL
				}
			}
		}
	}
}

func (t *RBTree) FixAfterDelete(node *Node) {
	//此时的解决方案是要么给要删除的节点所在的子树增加一个黑节点，要么让兄弟子树减少一个黑节点，此时需要整体思考
	for node != t.Root && node.Color == BLACK {
		if node == node.Parent.Left {
			brother := node.Parent.Right
			//判断此时兄弟节点是否是真正的兄弟节点
			if brother.Color == RED {
				brother.Color = BLACK
				brother.Parent.Color = RED
				t.LeftRotate(brother.Parent)
				//找到红黑树中的对应2-3-4树中的兄弟节点
				brother = node.Parent.Right
			}
			//情况3：找兄弟借，兄弟没的借
			if brother.Left.Color == BLACK && brother.Right.Color == BLACK { //兄弟是2节点
				//if brother.Left == t.NIL && brother.Right == t.NIL { //因为有循环的过程，所以这样写是错的
				//情况复杂，暂时不写
				brother.Color = RED
				node = node.Parent
			} else {
				//情况2：找兄弟借，兄弟有的借
				//分两种小情况：兄弟节点是3节点或者是4节点
				if brother.Right == t.NIL {
					brother.Color = RED
					brother.Left.Color = BLACK
					t.RightRotate(brother)
					brother = node.Parent.Right
				}
				brother.Color = node.Parent.Color
				node.Parent.Color = BLACK
				brother.Right.Color = BLACK
				t.LeftRotate(brother.Parent)
				node = t.Root
			}
		} else {
			brother := node.Parent.Left
			//判断此时兄弟节点是否是真正的兄弟节点
			if brother.Color == RED {
				brother.Color = BLACK
				brother.Parent.Color = RED
				t.LeftRotate(brother.Parent)
				//找到红黑树中的对应2-3-4树中的兄弟节点
				brother = node.Parent.Left
			}
			//情况3：找兄弟借，兄弟没的借
			//表示brother的左右都是t.Nil节点
			if brother.Right.Color == BLACK && brother.Left.Color == BLACK {
				//情况复杂，暂时不写
				brother.Color = RED
				node = node.Parent
			} else {
				//情况2：找兄弟借，兄弟有的借
				//分两种小情况：兄弟节点是3节点或者是4节点
				if brother.Left == t.NIL {
					brother.Color = RED
					brother.Right.Color = BLACK
					t.LeftRotate(brother)
					brother = node.Parent.Left
				}
				brother.Color = node.Parent.Color
				node.Parent.Color = BLACK
				brother.Left.Color = BLACK
				t.RightRotate(brother.Parent)
				node = t.Root
			}
		}
	}

	if node.Color == RED {
		node.Color = BLACK
	}
}

//此时的解决方案是要么给要删除的节点所在的子树增加一个黑节点，要么让兄弟子树减少一个黑节点，此时需要整体思考
func (t *RBTree) FixAfterDelete2(node *Node) {
	if node != t.Root && node.Color == BLACK {
		if node == node.Parent.Left {
			brother := node.Parent.Right
			//判断此时兄弟节点是否是真正的兄弟节点
			if brother.Color == RED {
				brother.Color = BLACK
				brother.Parent.Color = RED
				t.LeftRotate(brother.Parent)
				//找到红黑树中的对应2-3-4树中的兄弟节点
				brother = node.Parent.Right
			}
			//情况3：找兄弟借，兄弟没的借
			if brother.Left.Color == BLACK && brother.Right.Color == BLACK { //兄弟是2节点
				//if brother.Left == t.NIL && brother.Right == t.NIL { //因为有循环的过程，所以这样写是错的
				//情况复杂，暂时不写
				//让兄弟子树减少一个黑节点
				brother.Color = RED
				node = node.Parent
				t.FixAfterDelete(node)
			} else {
				//情况2：找兄弟借，兄弟有的借
				//分两种小情况：兄弟节点是3节点或者是4节点
				if brother.Right == t.NIL {
					brother.Color = RED
					brother.Left.Color = BLACK
					t.RightRotate(brother)
					brother = node.Parent.Right
				}
				brother.Color = node.Parent.Color
				node.Parent.Color = BLACK
				brother.Right.Color = BLACK
				t.LeftRotate(brother.Parent)
				node = t.Root
			}
		} else {
			brother := node.Parent.Left
			//判断此时兄弟节点是否是真正的兄弟节点
			if brother.Color == RED {
				brother.Color = BLACK
				brother.Parent.Color = RED
				t.LeftRotate(brother.Parent)
				//找到红黑树中的对应2-3-4树中的兄弟节点
				brother = node.Parent.Left
			}
			//情况3：找兄弟借，兄弟没的借
			//表示brother的左右都是t.Nil节点
			if brother.Right.Color == BLACK && brother.Left.Color == BLACK {
				//情况复杂，暂时不写
				brother.Color = RED
				node = node.Parent
			} else {
				//情况2：找兄弟借，兄弟有的借
				//分两种小情况：兄弟节点是3节点或者是4节点
				if brother.Left == t.NIL {
					brother.Color = RED
					brother.Right.Color = BLACK
					t.LeftRotate(brother)
					brother = node.Parent.Left
				}
				brother.Color = node.Parent.Color
				node.Parent.Color = BLACK
				brother.Left.Color = BLACK
				t.RightRotate(brother.Parent)
				node = t.Root
			}
		}
	}

	if node.Color == RED {
		node.Color = BLACK
	}
}

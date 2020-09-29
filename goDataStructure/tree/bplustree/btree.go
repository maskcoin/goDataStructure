package bplustree

import (
	"fmt"
	"strconv"
)

type Node struct {
	Leaf     bool //是否是叶子节点
	Parent   *Node
	N        int   //当前节点中Key的数量
	Keys     []int //存储数据
	Values   []interface{}
	Children []*Node
}

//M = 2*degree-1，也就是说当插入节点数=M时，此时需要分裂，分裂为2个key的数量分别为degree-1的节点。
//M=2*degree-1 => degree=(M+1)/2 => degree-1=(M-1)/2 => degree-1=ceil(M/2)-1
//此处的degree-1是除根节点外的码值的最小值，对应的最小分叉是degree=ceil(M/2)，所以这颗B树也叫(ceil(M/2), M)树
func NewNode(degree int, leaf bool) (ret *Node) {
	ret = &Node{
		Leaf:     leaf,
		Keys:     make([]int, 2*degree-1),
		Values:   make([]interface{}, 2*degree-1),
		Children: make([]*Node, 2*degree),
	}

	return
}

type BTree struct {
	Root   *Node
	Degree int
}

func NewBTree(degree int) *BTree {
	return &BTree{
		Degree: degree,
	}
}

//默认规定插入的数据不能相同
func (node *Node) FindInsertLeafNode(key int) (ret *Node) {
	if node != nil {
		if node.Leaf == true {
			ret = node
		} else {
			i := 0
			for i < node.N && node.Keys[i] < key {
				i++
			}

			ret = node.Children[i].FindInsertLeafNode(key)
		}
	}

	return
}

func (t *BTree) InsertToTargetNode(target *Node, keyNode *Node) {
	index := 0
	for index < target.N && target.Keys[index] < keyNode.Keys[0] {
		index++
	}

	for i := target.N - 1; i >= index; i-- {
		target.Keys[i+1] = target.Keys[i]
		target.Children[i+2] = target.Children[i+1]
	}

	target.Keys[index] = keyNode.Keys[0]

	if !target.Leaf {
		//这一步很重要，错一位移动children的指针，画图来理解
		for i := target.N; i >= index+1; i-- {
			target.Children[i+1] = target.Children[i]
		}

		target.Children[index] = keyNode.Children[0]
		keyNode.Children[0].Parent = target
		target.Children[index+1] = keyNode.Children[1] //至此插入成功
		keyNode.Children[1].Parent = target
	}

	target.N++

	if target.Leaf {
		target.Children[target.N] = target.Children[target.N-1]
	}
}

func (t *BTree) SplitInternal(target *Node) {
	degree := t.Degree

	newNode := NewNode(degree, target.Leaf)
	newNode.N = degree - 1

	//拷贝右边degree-1个码值到新的节点
	for i := 0; i < degree-1; i++ {
		newNode.Keys[i] = target.Keys[i+degree]
		target.Keys[i+degree] = 0
	}
	//拷贝右边degree个孩子指针到新的节点
	if !target.Leaf {
		for j := 0; j < degree; j++ {
			newNode.Children[j] = target.Children[j+degree]
			//清理target中被拷贝走的孩子
			target.Children[j+degree] = nil
			newNode.Children[j].Parent = newNode
		}
	}

	//修改target节点中的码值的数量
	target.N = degree - 1

	pNode := NewNode(t.Degree, false)

	pNode.Keys[0] = target.Keys[degree-1]
	//清理中间的那个码值
	target.Keys[degree-1] = 0
	pNode.N++
	pNode.Children[0] = target
	pNode.Children[1] = newNode

	parent := target.Parent

	target.Parent = pNode
	newNode.Parent = pNode

	if parent == nil {
		t.Root = pNode
	} else {
		//把pNode节点插入到parent节点中
		t.InsertToTargetNode(parent, pNode)
		if parent.N == 2*t.Degree-1 {
			t.SplitInternal(parent) //递归调整
		}
	}
}

func (t *BTree) SplitLeaf(target *Node) {
	degree := t.Degree

	newNode := NewNode(degree, target.Leaf)
	newNode.N = degree

	//拷贝右边degree个码值到新的节点
	for i := 0; i < degree; i++ {
		newNode.Keys[i] = target.Keys[i+degree-1]
		newNode.Values[i] = target.Values[i+degree-1]
		target.Keys[i+degree-1] = 0
		target.Values[i+degree-1] = nil
	}

	//让target的孩子指针指向newNode形成链表
	target.Children[degree-1] = newNode

	//修改target节点中的码值的数量
	target.N = degree - 1

	pNode := NewNode(t.Degree, false)

	//给pNode.Keys[0]的码值赋值newNode.Keys[0]，也就是原来target.Keys[degree-1]的值
	pNode.Keys[0] = newNode.Keys[0]
	pNode.Values = nil

	pNode.N++
	pNode.Children[0] = target
	pNode.Children[1] = newNode

	parent := target.Parent

	target.Parent = pNode
	newNode.Parent = pNode

	if parent == nil {
		t.Root = pNode
	} else {
		//把pNode节点插入到parent节点中
		t.InsertToTargetNode(parent, pNode)
		if parent.N == 2*t.Degree-1 {
			t.SplitInternal(parent) //递归调整
		}
	}
}

func (t *BTree) Insert(key int) {
	if t.Root == nil {
		keyNode := NewNode(t.Degree, true)
		keyNode.Keys[0] = key
		keyNode.N++
		t.Root = keyNode
	} else {
		target := t.Root.FindInsertLeafNode(key)

		keyNode := NewNode(t.Degree, true)
		keyNode.Keys[0] = key
		//直接插入，因为target是叶子节点，且target.N一定<2*degree-1，否则就会在之前插入后调整而分裂掉了
		t.InsertToTargetNode(target, keyNode)

		//插入之后判断target节点是否满了
		if target.N == 2*t.Degree-1 {
			//此时该节点需要被分裂
			t.SplitLeaf(target)
		}
	}
}

func (node *Node) String() (ret string) {
	//return fmt.Sprintf("{n=%d, keys=%v, children=%v}\n", node.N, node.Keys, node.Children)
	if node != nil && !node.Leaf {
		for i := 0; i < node.N; i++ {
			ret += strconv.Itoa(node.Keys[i])
		}

		for i := 0; i <= node.N; i++ {
			if node.Children[i] != nil {
				ret += node.Children[i].String()
			}
		}
	}

	return
}

func (t *BTree) String() (ret string) {
	ret = t.Root.String()
	ret += "\n*******************************************************************\n"
	node, _ := t.FindLeftChildMinNodeAndIndex(t.Root, 0)
	for i := 0; i < node.N; i++ {
		ret += strconv.Itoa(node.Keys[i])
	}

	for node.Children[node.N] != nil {
		node = node.Children[node.N]
		for i := 0; i < node.N; i++ {
			ret += strconv.Itoa(node.Keys[i])
		}
	}

	return
}

func (t *BTree) PrintTree(node *Node) {
	if node != nil {
		for i := 0; i < node.N; i++ {
			fmt.Print(node.Keys[i])
		}
		fmt.Println()
		i := 0
		for node.Children[i] != nil && i < 2*t.Degree-1 {
			t.PrintTree(node.Children[i])
			i++
		}
	}
}

func (t *BTree) FindLeftChildMaxNodeAndIndex(node *Node, index int) (ret *Node, idx int) {
	cur := node
	if !cur.Leaf {
		//当前码值的左子树
		cur = cur.Children[index]
		for !cur.Leaf {
			cur = cur.Children[cur.N]
		}

		ret = cur
		idx = cur.N - 1
	}

	return
}

func (t *BTree) FindLeftChildMinNodeAndIndex(node *Node, index int) (ret *Node, idx int) {
	cur := node
	if !cur.Leaf {
		//当前码值的左子树
		cur = cur.Children[index]
		for !cur.Leaf {
			for i := 0; i <= cur.N; i++ {
				if cur.Children[i] != nil {
					idx = i;
					break
				}
			}
			cur = cur.Children[idx]
		}

		ret = cur
		idx = 0
	}

	return
}


func (t *BTree) FindRightChildMinNodeAndIndex(node *Node, index int) (ret *Node, idx int) {
	cur := node
	if !cur.Leaf {
		//当前码值的左子树
		cur = cur.Children[index+1]
		for !cur.Leaf {
			cur = cur.Children[0]
		}
		ret = cur
		index = 0
	}

	return
}

func (t *BTree) Successor(key int) (ret *Node, index int) {
	//寻找到对应的节点
	target, idx := t.Root.Search(key)
	if !target.Leaf {
		ret, index = t.FindRightChildMinNodeAndIndex(target, idx)
	}

	return
}

func (t *BTree) Predecessor(key int) (ret *Node, index int) {
	//寻找到对应的节点
	target, idx := t.Root.Search(key)
	if !target.Leaf {
		ret, index = t.FindLeftChildMaxNodeAndIndex(target, idx)
	}
	return
}

//用后面的数据把自己给覆盖掉了
func (t *BTree) DeleteKeyAndNode(node *Node, index int) {
	j := index
	for j = index; j <= node.N-1; j++ {
		node.Keys[j] = node.Keys[j+1] //用后面的数据把自己给覆盖掉了
		node.Children[j] = node.Children[j+1]
	}

	node.Children[j] = nil
	node.N--
}

func (t *BTree) Merge(parent *Node, index int) (ret *Node) {
	leftNode := parent.Children[index]
	rightNode := parent.Children[index+1]

	ret = NewNode(t.Degree, rightNode.Leaf)

	for i := 0; i < t.Degree-1; i++ {
		ret.Keys[i] = leftNode.Keys[i]
		ret.Children[i] = leftNode.Children[i]
	}

	for j := 0; j < t.Degree-1; j++ {
		ret.Keys[t.Degree-1+j] = rightNode.Keys[j]
		ret.Children[t.Degree-1+j] = rightNode.Children[j]
	}
	ret.Children[2*t.Degree-2] = rightNode.Children[t.Degree-1]

	ret.N = 2*t.Degree - 2

	return
}

//B+树的删除不用写，直接把叶子节点的value改为nil就可以了
//删除叶子节点node.Keys[index]但是node.N==t.Degree-1，此时不能删，需要调整
//把父节点对应的码借借过来，父节点再从右兄弟或者左兄弟那里借一个码值过来(前提是兄弟有得借)
//如果兄弟没得借，此时强行把父节点所对应的码值，借过来，然后与兄弟进行合并，把合同后新的节点赋值给parent.Keys[idx+1]
//也就是让父节点码值的右孩子指向merge后产生的新的节点，此时父节点的左孩子parent.Keys[idx]，和parent.Children[idx]就没有用了
//把当前节点设置为父节点，index设置为父节点对应码值左孩子所在的索引，再次递归调整。
func (t *BTree) FixDelete(node *Node, index int) {
	if node.Parent == nil { //当前节点是根节点，两种情况统一成了一种解决方案
		t.DeleteKeyAndNode(node, index)
		if node.N == 0 {
			t.Root = node.Children[0]
		}
	} else if node.N > t.Degree-1 {
		parent := node.Parent

		//寻找到当前节点是父亲节点的第idx个孩子
		idx := 0
		for idx = 0; idx <= parent.N; idx++ {
			if parent.Children[idx] == node {
				break
			}
		}

		//对最右边的叶子节点做了特殊处理
		if index== node.N-1 && node.Children[node.N] == nil {
			t.DeleteKeyAndNode(node, index)
			if node.N == 0 {
				parent.Children[idx] = nil
				pre := parent.Children[idx-1]
				pre.Children[pre.N] = nil
			}

			return
		}

		//对最左边的叶子节点做了特殊处理
		minNode, ix := t.FindLeftChildMinNodeAndIndex(t.Root, 0)
		if node == minNode && index == ix {
			t.DeleteKeyAndNode(node, 0)
			if node.N == 0 {
				parent.Children[0] = nil
			}

			return
		}

		t.DeleteKeyAndNode(node, index)
	} else {
		parent := node.Parent

		//寻找到当前节点是父亲节点的第idx个孩子
		idx := 0
		for idx = 0; idx <= parent.N; idx++ {
			if parent.Children[idx] == node {
				break
			}
		}

		//对最右边的叶子节点做了特殊处理
		if index== node.N-1 && node.Children[node.N] == nil {
			t.DeleteKeyAndNode(node, index)
			if node.N == 0 {
				parent.Children[idx] = nil
				pre := parent.Children[idx-1]
				pre.Children[pre.N] = nil
			}

			return
		}

		//对最左边的叶子节点做了特殊处理
		minNode, ix := t.FindLeftChildMinNodeAndIndex(t.Root, 0)
		if node == minNode && index == ix {
			t.DeleteKeyAndNode(node, 0)
			if node.N == 0 {
				parent.Children[0] = nil
			}

			return
		}


		//因为需要把parent.Keys[idx]的值拿下来跟node.Keys[index]进行值替代
		//所有需要用parent.Keys[idx]的左孩子的最大值，或者parent.Keys[idx]的右孩子的最小值进行替代parent.Keys[idx]
		if idx == parent.N { //idx是父节点最右边的孩子
			idx-- //调整idx
			leftBrother := parent.Children[idx]
			leftBrotherKey := leftBrother.Keys[leftBrother.N-1]

			if leftBrother.N > t.Degree-1 {
				//此时leftBrother有得借
				//把parent.Keys[idx]的值拷贝下来
				for i := index; i > 0; i-- {
					node.Keys[i] = node.Keys[i-1]
					node.Children[i] = node.Children[i-1]
				}

				node.Keys[0] = parent.Keys[idx]

				//parent.Keys[idx]再把leftBrotherKey借过来
				parent.Keys[idx] = leftBrotherKey
				node.Children[0] = leftBrother.Children[leftBrother.N]
				leftBrother.Children[leftBrother.N] = nil
				//删除被借走的leftBrotherKey
				//t.DeleteKeyAndNode(leftBrother, leftBrother.N-1)
				leftBrother.Keys[leftBrother.N-1] = 0
				leftBrother.N--
			} else {
				//兄弟没得借，只能强行跟老父亲借，然后+兄弟节点合并
				//把parent.Keys[idx]借下来
				for i := index; i > 0; i-- {
					node.Keys[i] = node.Keys[i-1]
				}
				node.Keys[0] = parent.Keys[idx]

				//然后+parent.Keys[idx]+兄弟节点合并
				mergedNode := t.Merge(parent, idx)
				parent.Children[idx+1] = mergedNode

				t.FixDelete(parent, idx)
			}
		} else if idx == 0 { //idx是父节点最左边的孩子
			rightBrother := parent.Children[idx+1]
			rightBrotherKey := rightBrother.Keys[0]

			if rightBrother.N > t.Degree-1 {
				//此时rightBrother有得借
				//把parent.Keys[idx]的值拷贝下来
				for i := index; i <= node.N-2; i++ {
					node.Keys[i] = node.Keys[i+1]
				}

				for i := index; i <= node.N-1; i++ {
					node.Children[i] = node.Children[i+1]
				}

				node.Keys[node.N-1] = parent.Keys[idx]

				//parent.Keys[idx]再把successorNode.Keys[successorKey]借过来
				parent.Keys[idx] = rightBrotherKey
				node.Children[node.N] = rightBrother.Children[0]
				//删除被借走的rightBrotherKey
				t.DeleteKeyAndNode(rightBrother, 0)
			} else {
				//兄弟没得借，只能强行跟老父亲借，然后+兄弟节点合并
				//把parent.Keys[idx]借下来
				for i := index; i <= node.N-2; i++ {
					node.Keys[i] = node.Keys[i+1]
				}

				for i := index; i <= node.N-1; i++ {
					node.Children[i] = node.Children[i+1]
				}

				node.Keys[node.N-1] = parent.Keys[idx]

				//然后+兄弟节点合并
				mergedNode := t.Merge(parent, idx)
				parent.Children[idx+1] = mergedNode

				t.FixDelete(parent, idx)
			}
		} else { //idx同时有右兄弟和左兄弟
			//先尝试跟右兄弟借
			rightBrother := parent.Children[idx+1]
			rightBrotherKey := rightBrother.Keys[0]

			if rightBrother.N > t.Degree-1 {
				//此时rightBrother有得借
				//把parent.Keys[idx]的值拷贝下来
				for i := index; i <= node.N-2; i++ {
					node.Keys[i] = node.Keys[i+1]
				}

				for i := index; i <= node.N-1; i++ {
					node.Children[i] = node.Children[i+1]
				}

				node.Keys[node.N-1] = parent.Keys[idx]

				//parent.Keys[idx]再把successorNode.Keys[successorKey]借过来
				parent.Keys[idx] = rightBrotherKey
				node.Children[node.N] = rightBrother.Children[0]
				//删除被借走的rightBrotherKey
				t.DeleteKeyAndNode(rightBrother, 0)
			} else {
				//右兄弟没得借，再尝试跟左兄弟借
				leftBrother := parent.Children[idx]
				leftBrotherKey := leftBrother.Keys[leftBrother.N-1]

				if leftBrother.N > t.Degree-1 {
					//此时leftBrother有得借
					//把parent.Keys[idx]的值拷贝下来
					for i := index; i > 0; i-- {
						node.Keys[i] = node.Keys[i-1]
						node.Children[i] = node.Children[i-1]
					}

					node.Keys[0] = parent.Keys[idx]

					//parent.Keys[idx]再把leftBrotherKey借过来
					parent.Keys[idx] = leftBrotherKey
					node.Children[0] = leftBrother.Children[leftBrother.N]
					leftBrother.Children[leftBrother.N] = nil
					//删除被借走的leftBrotherKey
					//t.DeleteKeyAndNode(leftBrother, leftBrother.N-1)
					leftBrother.Keys[leftBrother.N-1] = 0
					leftBrother.N--
				} else {
					//兄弟没得借，只能强行跟老父亲借，然后+右兄弟节点合并
					//把parent.Keys[idx]借下来
					for i := index; i <= node.N-2; i++ {
						node.Keys[i] = node.Keys[i+1]
					}

					for i := index; i <= node.N-1; i++ {
						node.Children[i] = node.Children[i+1]
					}

					node.Keys[node.N-1] = parent.Keys[idx]

					//然后+兄弟节点合并
					mergedNode := t.Merge(parent, idx)
					parent.Children[idx+1] = mergedNode

					t.FixDelete(parent, idx)
				}
			}
		}
	}
	//}
}

func (t *BTree) Delete(key int) {
	node, idx := t.Root.Search(key) //找到了要删除的码值所在的节点和下标
	if node.Leaf == false {
		////寻找对应码值右子树的最小值
		successorNode, index := t.Successor(key)
		if successorNode.N == 1 {
			node.Keys[idx] = successorNode.Children[1].Keys[0]
		} else {
			node.Keys[idx] = successorNode.Keys[index+1]
		}


		node, idx = successorNode, index
	}

	t.FixDelete(node, idx)
}

func (node *Node) Search(key int) (ret *Node, index int) {
	if node != nil {
		i := 0
		for i < node.N && node.Keys[i] < key {
			i++
		}

		if i < node.N && node.Keys[i] == key {
			ret, index = node, i //找到
		} else if node.Leaf == false { //要么i=node.N，要么node节点中没有key这个值，此时到i所对应的指向的叶子节点中去寻找
			//进入孩子节点继续递归搜索
			ret, index = node.Children[i].Search(key)
		}
	}

	return
}

package segmenttree

type Merger interface {
	Merge(right Merger) Merger
}

type Int int

func (i Int) Merge(right Merger) Merger {
	return i + right.(Int)
}

type SegmentTree struct {
	Elements []Merger
	Tree     []Merger
}

func NewSegmentTree(arr []Merger) (ret *SegmentTree) {
	ret = &SegmentTree{
		Elements: make([]Merger, len(arr)),
		Tree:     make([]Merger, 4*len(arr)),
	}

	for i, v := range arr {
		ret.Elements[i] = v
	}

	return
}

//从Tree []Merge数组的treeIndex位置，创建表示Elements []Merge数组区间[l..r]的线段树
//所谓构建一颗树就是把t.Tree[treeIndex]的这个元素中填入有意义的值
func (t *SegmentTree) BuildSegmentTree(treeIndex int, l int, r int) {
	if l == r {
		t.Tree[treeIndex] = t.Elements[l]
		return
	} else if l < r {
		leftTreeIndex := t.LeftChild(treeIndex)
		rightTreeIndex := t.RightChild(treeIndex)
		mid := l + (r-l)/2
		t.BuildSegmentTree(leftTreeIndex, l, mid)
		t.BuildSegmentTree(rightTreeIndex, mid+1, r)

		t.Tree[treeIndex] = t.Tree[leftTreeIndex].Merge(t.Tree[rightTreeIndex])
	}
}

//在以t.Tree []Merge数组的treeIndex为根的线段树中的t.Tree.[l..r]的范围内，搜索区间t.Elements[qr..qr]的值
//t.Tree[treeIndex]是线段树的树根节点，l,r是这颗线段树中t.Elements[l..r]的这颗区间所代表的逻辑节点的左右下标，
//ql,qr是t.Elements[ql..qr]这个数组区间的左右下标
func (t *SegmentTree) query(treeIndex, l, r, ql, qr int) (ret Merger) {
	if l == ql && r == qr {
		return t.Tree[treeIndex]
	}

	mid := l + (r-l)/2

	if mid < ql {
		ret = t.query(t.RightChild(treeIndex), mid+1, r, ql, qr)
	} else if qr <= mid {
		ret = t.query(t.LeftChild(treeIndex), l, mid, ql, qr)
	} else {
		left := t.query(t.LeftChild(treeIndex), l, mid, ql, mid)
		right := t.query(t.RightChild(treeIndex), mid+1, r, mid+1, qr)
		ret = left.Merge(right)
	}

	return
}

//返回区间[l,q]的值
func (t *SegmentTree) Query(ql, qr int) (ret Merger) {
	if (ql < 0 || ql >= len(t.Elements)) || (qr < 0 || qr >= len(t.Elements)) || (ql > qr) {
		panic("index wrong")
	}

	return t.query(0, 0, len(t.Elements)-1, ql, qr)
}

//将index位置的值，更新为elem
func (t *SegmentTree) Set(index int, elem Merger) {
	if index < 0 || index >= len(t.Elements) {
		panic("index err")
	}

	t.Elements[index] = elem
	//t.BuildSegmentTree(0, 0, len(t.Elements)-1)
	t.setElem(0, 0, len(t.Elements)-1, index, elem)
}

func (t *SegmentTree) setElem(treeIndex, l, r, index int, elem Merger) {
	if l == r {
		t.Tree[treeIndex] = t.Elements[l]
		return
	}

	mid := l + (r-l)/2

	if index <= mid {
		t.setElem(t.LeftChild(treeIndex), l, mid, index, elem)
	} else {
		t.setElem(t.RightChild(treeIndex), mid+1, r, index, elem)
	}

	t.Tree[t.LeftChild(treeIndex)].Merge(t.Tree[t.RightChild(treeIndex)])
}

func (t *SegmentTree) Get(index int) Merger {
	if index < 0 || index >= len(t.Elements) {
		panic("index error")
	}

	return t.Elements[index]
}

func (t *SegmentTree) GetSize() int {
	return len(t.Elements)
}

//返回一个索引所表示的元素的左孩子节点的索引
func (t *SegmentTree) LeftChild(index int) int {
	return 2*index + 1
}

//返回一个索引所表示的元素的右孩子节点的索引
func (t *SegmentTree) RightChild(index int) int {
	return 2*index + 2
}

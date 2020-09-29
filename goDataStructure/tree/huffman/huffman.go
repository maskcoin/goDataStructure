package huffman

import (
	"container/heap"
	"fmt"
)

type HuffmanTree interface {
	Freq() int
}

type HuffmanLeaf struct {
	freq  int
	value rune //int32
}

type HuffmanNode struct {
	freq        int
	left, right HuffmanTree
}

func (leaf *HuffmanLeaf) Freq() int {
	return leaf.freq
}

func (node *HuffmanNode) Freq() int {
	return node.freq
}

type TreeHeap []HuffmanTree

func (heap TreeHeap) Len() int {
	return len(heap)
}

func (heap TreeHeap) Less(i, j int) bool {
	return heap[i].Freq() < heap[j].Freq()
}

//压入
func (heap TreeHeap) Push(elem interface{}) {
	heap = append(heap, elem.(HuffmanTree))
}

//弹出
func (heap TreeHeap) Pop() (ret interface{}) {
	ret = heap[len(heap)-1]
	heap = heap[:len(heap)-1]
	return
}

func (heap TreeHeap) Swap(i,j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func BuildTree(symbFreqs map[rune]int) HuffmanTree {
	var trees TreeHeap
	for c, f := range symbFreqs {
		trees = append(trees, &HuffmanLeaf{
			freq:  f,
			value: c,
		})
	}

	heap.Init(trees) //给函数参数一个原材料，容器对象，这个容器对象需要实现一定的接口，初始化后这个容器就是按堆排序排序好了的

	for trees.Len() > 1 {
		a := heap.Pop(trees).(HuffmanTree)
		b := heap.Pop(trees).(HuffmanTree)
		heap.Push(trees, &HuffmanNode{
			freq:  a.Freq() + b.Freq(),
			left:  a,
			right: b,
		})
	}

	return heap.Pop(trees).(HuffmanTree)
}

func ShowFreq(tree HuffmanTree, prefix []byte)  {
	switch i := tree.(type) {
	case *HuffmanLeaf:
		fmt.Printf("%c\t%d\n", i.value, i.freq)
	case *HuffmanNode:
	}
}
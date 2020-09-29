package heapsort

import "fmt"

func HeapSortMax(arr []int, length int) []int {
	if length > 1 {
		n := length/2 - 1 // n代表a[n]的下标，即有0-n编号是个有叶节点的树根的下标编号
		for i := n; i >= 0; i-- {
			topMax := i
			left := 2*i + 1
			right := 2*i + 2

			if arr[left] > arr[topMax] {
				topMax = left
			}

			if right <= length-1 && arr[right] > arr[topMax] {
				topMax = right
			}

			if topMax != i {
				arr[i], arr[topMax] = arr[topMax], arr[i]
			}
		}
	}
	return arr
}

func HeapSort(arr []int) []int {
	length := len(arr)

	for i := 0; i < length; i++ {
		lastLen := length - i
		HeapSortMax(arr, lastLen)
		arr[0], arr[lastLen-1] = arr[lastLen-1], arr[0]
	}
	return arr
}

type TreePriorityQueueNode struct {
	element interface{}
	left    *TreePriorityQueueNode
	right   *TreePriorityQueueNode
	priority int
}


//开辟一个节点
func NewLeftHeap(element interface{}) *TreePriorityQueueNode {
	head := new(TreePriorityQueueNode)
	head.element = element
	head.left = nil
	head.right = nil
	head.priority = 0
	return head
}

//heap1.element.(int) < heap2.element.(int)
func MergeSort(heap1, heap2 *TreePriorityQueueNode) *TreePriorityQueueNode {
	if heap1.left == nil {
		heap1.left = heap2
	} else {
		heap1.right = Merge(heap1.right, heap2)
		if heap1.left.priority < heap1.right.priority {
			heap1.left, heap1.right = heap1.right, heap1.left
		}
		heap1.priority = heap1.right.priority+1
	}
	return heap1
}

//确保有序
func Merge(heap1, heap2 *TreePriorityQueueNode) *TreePriorityQueueNode {
	if heap1 == nil {
		return heap2
	}

	if heap2 == nil {
		return heap1
	}
	//递归
	if heap1.element.(int) < heap2.element.(int) {
		return MergeSort(heap1, heap2)
	} else {
		return MergeSort(heap2, heap1)
	}
}

func (queue *TreePriorityQueueNode) Insert(data interface{}) *TreePriorityQueueNode{
	node := new(TreePriorityQueueNode)
	node.element = data
	node.left = nil
	node.right = nil
	node.priority = 0
	//插入用归并实现
	ret := Merge(node, queue)
	return ret
}

func (queue *TreePriorityQueueNode) DeleteMin() (*TreePriorityQueueNode, interface{}) {
	if queue == nil {
		return nil, nil
	} else {
		leftHeap := queue.left
		rightHeap := queue.right
		value := queue.element
		queue = nil
		ret := Merge(leftHeap, rightHeap)
		return ret, value
	}
}

func Print(queue *TreePriorityQueueNode) {
	if queue != nil {
		fmt.Println(queue.element)
		Print(queue.left)
		Print(queue.right)
	}
}
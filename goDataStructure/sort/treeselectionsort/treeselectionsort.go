package treeselectionsort

import "math"

type node struct {
	value int // 叶子的数据
	isOK  bool
	rank  int // 叶子的排序
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func TreeSelectionSort(arr []int) []int {
	// 树的层数
	var level int                         // 树的层数
	var retArr = make([]int, 0, len(arr)) // 保存最终结果
	for pow(2, level) < len(arr) {
		level++ // 求出可以覆盖所有元素的层数
	}
	//fmt.Println("level=", level)
	leaf := pow(2, level)
	//fmt.Println("leaf=", leaf)
	tree := make([]node, leaf*2-1) // 构造逻辑树的数组，即树的节点数量
	// 填充叶子
	for i := 0; i < len(arr); i++ {
		tree[leaf+i-1] = node{
			value: arr[i],
			isOK:  true,
			rank:  i,
		}
	}
	for i := 0; i < len(arr); i++ {
		// 进行对比
		levelCopy := level
		for i := levelCopy; i >= 1; i-- {
			nodeCount := pow(2, i)
			for j := 0; j < nodeCount/2; j++ { // j相当于叶子节点对应的根节点的逻辑序号
				leftNode := nodeCount - 1 + j*2
				rightNode := leftNode + 1
				mid := (leftNode - 1) / 2
				// 中间节点存储最小值
				if tree[leftNode].isOK && tree[rightNode].isOK && tree[leftNode].value < tree[rightNode].value {
					tree[mid] = tree[leftNode]
				} else if tree[leftNode].isOK && tree[rightNode].isOK && tree[leftNode].value >= tree[rightNode].value {
					tree[mid] = tree[rightNode]
				} else if tree[leftNode].isOK && !tree[rightNode].isOK{
					tree[mid] = tree[leftNode]
				} else if !tree[leftNode].isOK && tree[rightNode].isOK {
					tree[mid] = tree[rightNode]
				} else if !tree[leftNode].isOK && !tree[rightNode].isOK {
					tree[mid] = tree[leftNode]
				}
			}
		}
		retArr = append(retArr, tree[0].value) // 保存最顶端的最小数
		tree[leaf-1+tree[0].rank] = node{
		}
	}
	return retArr
}

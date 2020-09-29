package AStar

//探测的栈
type OpenList []*AStarPoint

func (list OpenList) Len() int {
	return len(list)
}

func (list OpenList) Less(i, j int) bool {
	return list[i].FVal < list[j].FVal
}

func (list OpenList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// 节点加入到栈中
func (list OpenList) Push(data interface{}) {
	list = append(list, data.(*AStarPoint))
}

func (list OpenList) Pop() interface{} {
	old := list
	n := len(old)
	x := old[n-1] // 最后一个
	list = list[:n-1] // 切片的截取，去掉了最后一个
	return x
}
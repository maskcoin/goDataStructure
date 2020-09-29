package main

import (
	"fmt"
	"os"
)

var Parent [6]int //定义前驱节点
var Rank [6]int //定义前驱节点

func Initialize(parent []int) {
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}
}

func FindRoot(x int) int { //查询根节点
	for Parent[x] != -1 {
		x = Parent[x]
	}

	return x
}

//返回1，成功合并x和y节点
//返回0，当两个点在同一个集合中时，合并失败
func Union(x int, y int) (ret int) {
	x_root := FindRoot(x)
	y_root := FindRoot(y)

	if x_root != y_root {
		if Rank[x_root] > Rank[y_root] {
			Parent[y_root] = x_root
		} else if Rank[x_root] < Rank[y_root] {
			Parent[x_root] = y_root
		} else {
			Parent[x_root] = y_root
			Rank[y_root]++
		}

		ret = 1
	}

	return
}

func main() {
	Initialize(Parent[:])
	edge := [6][2]int{
		{0, 1},
		{1, 2},
		{1, 3},
		{3, 4},
		{2, 5},
		{5, 4},
	}
	var i int
	for i = 0; i < len(edge); i++ {
		x := edge[i][0]
		y := edge[i][1]
		if Union(x, y) == 0 {
			fmt.Println("circle detected!")
			os.Exit(0)
		}
	}

	fmt.Println("no circles found")
}

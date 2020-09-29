package recursion

import "fmt"

//解决递归走出来
func AIOut(aIData [L][C]int, i int, j int) bool {
	aIData[i][j] = 3 //避免回头路
	if i == L-1 && j == C-1 {
		CanGoOut = true
		fmt.Println("迷宫可以走出来")
		Show(aIData)
	} else {
		if j+1 <= C-1 && aIData[i][j+1] < 2 && CanGoOut != true {
			AIOut(aIData, i, j+1)
		}
		if i+1 <= L-1 && aIData[i+1][j] < 2 && CanGoOut != true {
			AIOut(aIData, i+1, j)
		}
		if j-1 >= 0 && aIData[i][j-1] < 2 && CanGoOut != true {
			AIOut(aIData, i, j-1)
		}
		if i-1 >= 0 && aIData[i-1][j] < 2 && CanGoOut != true {
			AIOut(aIData, i-1, j)
		}
		if CanGoOut != true {
			aIData[i][j] = 0
		}
	}
	return CanGoOut
}

func AIMoveOut()  {
	AIData[0][0] = 1
	for IPos != L-1 && JPos != C -1 {
		if IPos-1>=0 && AIData[IPos-1][JPos] == 3{
			AIData[IPos-1][JPos] = 0
			Run("w")
		}
		if IPos+1<=L-1 && AIData[IPos+1][JPos] == 3{
			AIData[IPos+1][JPos] = 0
			Run("s")
		}
		if JPos+1<=C-1 && AIData[IPos][JPos+1] == 3{
			AIData[IPos][JPos+1] = 0
			Run("d")
		}
		if JPos-1>=0 && AIData[IPos][JPos-1] == 3{
			AIData[IPos][JPos-1] = 0
			Run("a")
		}
	}
}

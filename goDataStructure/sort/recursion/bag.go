package recursion

const N2 = 4
const W2 = 5

var Weight2 = []int{3, 2, 1, 2}
var Value2 = []int{4, 3, 2, 1}
var DP [N2 + 1][W2 + 1]int

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//从最小值开始，填写一个二维数组，二维数组中对应的每个值，是只有第i个商品，背包容量为j时，小偷能偷出的价值的最大值
func Solve() int {
	length := len(Value2)
	for i := 0; i < length; i++ {
		for j := Weight2[i]; j <= W2; j++ {
			DP[i+1][j] = Max(DP[i][j], DP[i][j-Weight2[i]]+Value2[i])
		}
	}
	return DP[N2][W2]
}

//挑选第i个物品，j是背包剩余的空间
func Rec(i,j int) int {
	ret := 0
	// N2是数组的下标
	if i == N2 || j <= 0 {
		// 已经没有剩余的物品了
		return ret
	}

	if j < Weight2[i] {
		// 无法挑选这个物品
		ret = Rec(i+1, j)
	} else {
		// 挑选和不挑选两种情况都尝试一下
		ret = Max(Rec(i+1, j), Rec(i+1, j-Weight2[i]) + Value2[i])
	}

	return ret
}

//挑选第i个物品，j是背包剩余的空间
func Rec2(i, j int) int {
	//备份Rec1(i,j)递归出来的结果，下次运行时，不需要再重复计算了
	if DP[i][j] >= 0 {
		// 已经计算过的话，直接使用之前的结果
		return DP[i][j]
	}

	ret := 0
	if i == N2 || j <= 0 {
		// 已经没有剩余的物品了
		return ret
	}
	if j < Weight2[i] {
		// 无法挑选这个物品
		ret = Rec(i+1, j)
	} else {
		// 挑选和不挑选两种情况都尝试一下
		ret = Max(Rec(i+1, j), Rec(i+1, j-Weight2[i]) + Value2[i])
	}

	DP[i][j] = ret
	return ret
}
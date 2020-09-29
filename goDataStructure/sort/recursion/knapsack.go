package recursion

const N = 6
const W = 21

var B [N][W]int //B[5][20]
var Weight = []int{0, 2, 3, 4, 5, 9}
var Value = []int{0, 3, 4, 5, 8, 10}

func Knapsack()  {
	var k, w int
	for k = 1; k < N; k++ {
		for w = 1; w < W; w++ {
			if Weight[k] > w{
				B[k][w] = B[k-1][w]
			} else {
				value1 := B[k-1][w-Weight[k]] + Value[k]
				//value2 := B[k-1][w]
				//if value1 > value2 {
				//	B[k][w] = value1
				//} else {
				//	B[k][w] = value2
				//}
				B[k][w] = value1
			}
		}
	}
}


//var Weight = []int{3, 2, 1, 2}
//var Val = []int{4, 3, 2, 2}
//var DP [N+1][W+1]int
//var Record  [N][W+1]int //保存中间结果

//func Init() {
//	for i := 0; i < N; i++ {
//		for j := 0; j <= W; j++ {
//			//Record[i][j] = -1
//		}
//	}
//}

package recursion

var Position [9]int
var SubNum = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var ZhiNum = []int{1, 2, 3, 5, 7, 11, 13, 17, 19}

func IsZhi(n int) bool  {
	for i := 0; i < 9; i++ {
		if n == ZhiNum[i] {
			return true
		}
	}
	return false
}

func CheckB(i, n int) bool {
	//纵
	if i-3 >= 0 { // 跳过0，1，2
		if IsZhi(Position[i]+Position[i-3]) == false {
			return false
		}
	}
	//横
	if i%3 != 0 { // 跳过0，3，6
		if IsZhi(Position[i]+Position[i-1]) == false {
			return false
		}
	}
	return true
}

func FillBox(i,n,r int, count *int)  {
	
}

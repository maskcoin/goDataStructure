package recursion

import (
	"datastructure/link/stack"
	"fmt"
)

const Size = 8
const Num = Size + 2

type Pos struct {
	X int
	Y int
}

var ChessBoard [Num][Num]int

var Direction = [3]*Pos {
	&Pos{
		X: -1,
		Y: -1,
	},
	&Pos{
		X: 0,
		Y: -1,
	},
	&Pos{
		X: 1,
		Y: -1,
	},
}

var Solution = stack.NewStack()

var Count int

func Init()  {
	for i := 0; i < Num; i+=(Num-1) {
		for j := 0; j < Num; j++ {
			ChessBoard[i][j] = 2
			ChessBoard[j][i] = 2
		}
	}
}

func Print()  {
	for !Solution.IsEmpty() {
		pos:=Solution.Pop().(*Pos)
		fmt.Print("( ", pos.X, pos.Y, ") ")
	}
	fmt.Println()
	for i := 0; i < Num; i++ {
		for j := 0; j < Num; j++ {
			switch ChessBoard[i][j] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func Check(x,y,d int) bool {
	flag := true
	for flag {
		x += Direction[d].X
		y += Direction[d].Y
		flag = ChessBoard[x][y] == 0
	}
	return ChessBoard[x][y] == 2
}

func RunQueen(i int)  {
	if i <= Size {
		for j := 1; j <= Size; j++ {
			if Check(j, i, 0) && Check(j, i, 1) && Check(j, i, 2) {
				ChessBoard[j][i] = 1
				Solution.Push(&Pos{
					X: j,
					Y: i,
				})
				RunQueen(i+1)
				ChessBoard[j][i] = 0
				Solution.Pop()
			}
		}
	} else {
		Count++
		Print()
	}
}

func StartGoQueen()  {
	RunQueen(1)
	fmt.Println("Total:", Count)
}




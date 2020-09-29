package recursion

import "fmt"

func Show(arr [L][C]int)  {
	fmt.Println("-------------------------------------")
	for i := 0; i < L; i++ {
		for j := 0; j < C; j++ {
			fmt.Printf("%3d", arr[i][j])
		}
		fmt.Println()
	}
}

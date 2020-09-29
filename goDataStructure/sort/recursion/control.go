package recursion

//"wsad" w上，s下,a左，d右
func Run(direct string) {
	if direct == "w" {
		if IPos-1 >= 0 {
			Data[IPos][JPos], Data[IPos-1][JPos] = Data[IPos-1][JPos], Data[IPos][JPos]
			IPos -= 1
		}
	} else if direct == "s" {
		if IPos+1 <= L-1 {
			Data[IPos][JPos], Data[IPos+1][JPos] = Data[IPos+1][JPos], Data[IPos][JPos]
			IPos += 1
		}
	} else if direct == "a" {
		if JPos-1 >= 0 {
			Data[IPos][JPos], Data[IPos][JPos-1] = Data[IPos][JPos-1], Data[IPos][JPos]
			JPos -= 1
		}
	} else if direct == "d" {
		if JPos+1 <= C-1 {
			Data[IPos][JPos], Data[IPos][JPos+1] = Data[IPos][JPos+1], Data[IPos][JPos]
			JPos += 1
		}
	} else {

	}

	Show(Data)
}

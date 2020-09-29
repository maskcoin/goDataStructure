package AStar

type Point struct {
	X    int    // X坐标
	Y    int    // Y坐标
	View string // *,-,.
}

func NewPoint(x, y int, view string) *Point {
	return &Point{
		X:    x,
		Y:    y,
		View: view,
	}
}

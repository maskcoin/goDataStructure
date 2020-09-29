package AStar

import "math"

//AStar算法地图上的点结构
// AStarPoint可以理解为Point的元数据
type AStarPoint struct {
	Point  *Point
	Father *AStarPoint
	GVal   int //g(n)表示从初始结点到任意结点n的代价
	HVal   int //h(n)表示从结点n到目标点的启发式评估代价
	FVal   int // 当从初始点向目标点移动时，A*权衡这两者。每次进行主循环时，它检查f(n)最小的结点n，其中f(n) = g(n) + h(n)。
}

func NewAStarPoint(p *Point, father *AStarPoint, end *AStarPoint) (ap *AStarPoint) {
	ap = &AStarPoint{
		Point:  p,
		Father: father,
	}
	if end != nil {
		ap.CalcFVal(end) // 创建的时候就计算节点的评估
	}
	return ap
}

func (p *AStarPoint) CalcGVal() int {
	if p.Father != nil {
		// 从父节点走来需要的代价
		deltaX := math.Abs(float64(p.Father.Point.X - p.Point.X))
		deltaY := math.Abs(float64(p.Father.Point.Y - p.Point.Y))
		if deltaX == 1 && deltaY == 0 {
			//移动一步
			p.GVal = p.Father.GVal + 10 //评估代价，姑且加10,可以理解为走过来需要的体力或者能量
		} else if deltaX == 0 && deltaY == 1 {
			p.GVal = p.Father.GVal + 10
		} else if deltaX == 1 && deltaY == 1 {
			p.GVal = p.Father.GVal + 14
		} else {
			panic("error")
		}
	}
	return p.GVal
}

// 计算当前节点与目标节点的差距
func (p *AStarPoint) CalcHVal(end *AStarPoint) int {
	p.HVal = int(math.Abs(float64(end.Point.X-p.Point.X)) + math.Abs(float64(end.Point.Y-p.Point.Y)))
	return p.HVal
}

func (p *AStarPoint) CalcFVal(end *AStarPoint) int {
	p.FVal = p.CalcGVal() + p.CalcHVal(end)
	return p.FVal
}

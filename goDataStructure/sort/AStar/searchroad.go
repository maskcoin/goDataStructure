package AStar

import "container/heap"

type SearchRoad struct {
	M       *Map
	Start   *AStarPoint
	End     *AStarPoint
	CloseLi map[string]*AStarPoint // 关闭， 不通的节点集合
	OpenLi  OpenList               // 通路
	OpenSet map[string]*AStarPoint // 去掉重复
	Road    []*AStarPoint          // 通道
}

func NewSearchRoad(startX, startY, endX, endY int, m *Map) *SearchRoad {
	sr := &SearchRoad{
		M:       m,
		Start:   NewAStarPoint(NewPoint(startX, startY, "S"), nil, nil),
		End:     NewAStarPoint(NewPoint(endX, endY, "E"), nil, nil),
		CloseLi: make(map[string]*AStarPoint),
		OpenSet: make(map[string]*AStarPoint),
	}
	heap.Init(sr.OpenLi)
	heap.Push(sr.OpenLi, sr.Start)                                        // 最小堆压入开始节点
	sr.OpenSet[PointAsKey(sr.Start.Point.X, sr.Start.Point.Y)] = sr.Start // 开放集合压入开始节点

	//所有的障碍加入block
	for k, v := range m.Blocks {
		sr.CloseLi[k] = NewAStarPoint(v, nil, nil)
	}
	return sr
}

//A*算法的核心
func (sr *SearchRoad) FindOutShortestPath() bool {
	// 如果开放节点大于0，永远循环下去
	for len(sr.OpenLi) > 0 {
		//如果找不到路从开放节点中取出放入关闭节点
		x := heap.Pop(sr.OpenLi)                                              //取出一个节点
		curPoint := x.(*AStarPoint)                                           // 取得当前节点
		delete(sr.OpenSet, PointAsKey(curPoint.Point.X, curPoint.Point.Y))    //走过的路
		sr.CloseLi[PointAsKey(curPoint.Point.X, curPoint.Point.Y)] = curPoint //障碍及走过的路

		adjacs := sr.M.GetAdjacentPoints(curPoint.Point)
		for _, p := range adjacs {
			asp := NewAStarPoint(p, curPoint, sr.End) // 创建A*节点
			// 我们找到了结束的节点
			if PointAsKey(asp.Point.X, asp.Point.Y) == PointAsKey(sr.End.Point.X, sr.End.Point.Y){
				for asp.Father != nil {
					sr.Road = append(sr.Road, asp)
					asp.Point.View = "*"
					asp = asp.Father
				}
				return true
			}

			_, ok := sr.CloseLi[PointAsKey(p.X, p.Y)]
			if ok {
				continue
			}

			existAP, ok := sr.OpenSet[PointAsKey(p.X, p.Y)] //取出开放的节点
			if !ok {
				heap.Push(sr.OpenLi, asp) // 节点不存在就压入
				sr.OpenSet[PointAsKey(asp.Point.X, asp.Point.Y)] = asp
			} else {
				//如果节点存在，经过对比，取得最短路径
				oldGVal, oldfather := existAP.GVal, existAP.Father
				existAP.Father = curPoint
				existAP.CalcGVal()//计算最短的值

				//新的节点距离比老节点还要短
				if existAP.GVal > oldGVal {
					//保存最短的
					existAP.Father = oldfather
					existAP.GVal = oldGVal
				}
			}
		}
	}
	return false
}

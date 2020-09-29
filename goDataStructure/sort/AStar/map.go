package AStar

import (
	"fmt"
	"strconv"
	"strings"
)

func PointAsKey(x, y int) (key string) {
	key = strconv.Itoa(x) + "," + strconv.Itoa(y) //坐标转化为字符串
	return
}

type Map struct {
	Points [][]*Point        // 地图，保存矩阵
	Blocks map[string]*Point // 字符串对应每一个节点，标记障碍
	MaxX   int               // 最大的X坐标，X表示第几行
	MaxY   int               // 最大的Y坐标，Y表示第几列
}

func NewMap(charMap []string) *Map {
	m := &Map{}
	m.Points = make([][]*Point, len(charMap))
	m.Blocks = make(map[string]*Point, len(charMap)*2) // 两倍边长
	for x, row := range charMap {
		cols := strings.Split(row, " ") // 基于空格，进行切割
		m.Points[x] = make([]*Point, len(cols))
		for y, view := range cols {
			m.Points[x][y] = &Point{
				X:    x,
				Y:    y,
				View: view,
			}
			if view == "X" {
				//标记障碍
				m.Blocks[PointAsKey(x, y)] = m.Points[x][y]
			}
		}
	}
	m.MaxX = len(m.Points)
	m.MaxY = len(m.Points[0])
	return m
}

// 抓取相邻节点，返回一个集合
//	   x+1, y-1   x+1,y   x+1,y+1
//		 x, y-1  	x,y，	x,y+1
//		 x-1, y-1   x-1,y   x-1,y+1
//
func (m *Map) GetAdjacentPoints(curPoint *Point) (adjacentPoints []*Point) {
	if x, y := curPoint.X, curPoint.Y-1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacentPoints = append(adjacentPoints, m.Points[x][y])
	}
	if x, y := curPoint.X, curPoint.Y+1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacentPoints = append(adjacentPoints, m.Points[x][y])
	}
	if x, y := curPoint.X+1, curPoint.Y; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacentPoints = append(adjacentPoints, m.Points[x][y])
	}
	if x, y := curPoint.X-1, curPoint.Y; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacentPoints = append(adjacentPoints, m.Points[x][y])
	}
	if x, y := curPoint.X+1, curPoint.Y-1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacentPoints = append(adjacentPoints, m.Points[x][y])
	}
	if x, y := curPoint.X-1, curPoint.Y+1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacentPoints = append(adjacentPoints, m.Points[x][y])
	}
	if x, y := curPoint.X+1, curPoint.Y+1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacentPoints = append(adjacentPoints, m.Points[x][y])
	}
	if x, y := curPoint.X-1, curPoint.Y-1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacentPoints = append(adjacentPoints, m.Points[x][y])
	}
	return
}

//打印地图，按照我们寻找好的通路
func (m *Map) PrintMap(road *SearchRoad) {
	fmt.Println("地图的边界", m.MaxX, m.MaxY)
	for x := 0; x < m.MaxX; x++ {
		for y := 0; y < m.MaxY; y++ {
			if road != nil {
				if x==road.Start.Point.X && y== road.Start.Point.Y {
					fmt.Printf("%2s", "S") // S代表开始
					goto Next
				}

				if x==road.End.Point.X && y== road.End.Point.Y {
					fmt.Printf("%2s", "E") // E代表结束
					goto Next
				}

				for i := 0; i < len(road.Road); i++ { // 循环找路
					if road.Road[i].Point.X == x && road.Road[i].Point.Y == y {
						fmt.Printf("%2s", "*") // 代表走过
						goto Next
					}
				}
			}
			fmt.Printf("%2s", m.Points[x][y].View)
			Next:
		}
		fmt.Println()
	}
}

package hashtree

type Node struct {
	Key      int
	Num    int
	Children map[int]*Node
}

func (node *Node) InsertValues(keys []int) {
	if len(keys) > 0 {
		childNode, ok := node.Children[keys[0]]

		if !ok {
			childNode = &Node{
				Key:      keys[0],
			}
			if node.Children == nil {
				node.Children = make(map[int]*Node)
			}

			node.Children[keys[0]] = childNode

			node.Num += 1 //记录层次
		}

		if len(keys) > 1 {
			childNode.InsertValues(keys[1:])
		}
	}
}



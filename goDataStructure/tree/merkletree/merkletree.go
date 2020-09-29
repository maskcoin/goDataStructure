package merkletree

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Hash  []byte
}

type Tree struct {
	Root *Node
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func NewNode(left, right *Node, hash []byte) (ret *Node) {
	ret = &Node{
		Left:  left,
		Right: right,
	}

	if left == nil && right == nil {
		ret.Hash = hash
	} else {
		var tempHash []byte
		tempHash = append(tempHash, left.Hash...)
		tempHash = append(tempHash, right.Hash...)
		hashSum256Once := sha256.Sum256(tempHash)
		hashSum256Twice := sha256.Sum256(hashSum256Once[:])
		ret.Hash = hashSum256Twice[:]
	}

	return
}

func NewTree(hashes [][]byte) (ret *Tree) {
	ret = &Tree{}

	var nodes []*Node

	for _, hash := range hashes {
		leaf := NewNode(nil, nil, hash)
		nodes = append(nodes, leaf)
	}

	j := 0 // 每一层的第一个元素
	//每次折半处理
	for length := len(hashes); length > 1; length = (length + 1) / 2 {
		for i := 0; i < length; i += 2 {
			right := min(i+1, length-1)
			node := NewNode(nodes[j+i], nodes[j+right], nil)
			nodes = append(nodes, node)
		}
		j += length
	}

	ret.Root = nodes[j]

	return
}

func (node *Node) PrintSelf() {
	if node != nil {
		fmt.Printf("%x\n",node.Hash)
		node.Left.PrintSelf()
		node.Right.PrintSelf()
	}
}
